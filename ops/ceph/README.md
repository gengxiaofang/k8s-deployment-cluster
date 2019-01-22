### 一、安装 ceph
**1.1安装ansible**
```
yum -y install http://dist.yongche.com/centos/7/epel/x86_64/Packages/a/ansible-2.6.5-1.el7.noarch.rpm
```
提示：官方对 ansible 版本有明确的要求，目前支持 2.4 以及 2.6 不能太高也不能太新。

**1.2下载ceph-ansible**  
到 https://github.com/ceph/ceph-ansible/releases 下载最新稳定版本，建议下载对应自己想安装的版本。如何知道自己下载的 ceph-ansible 支持自己想安装的 ceph 版本？只能查看与之对应的 Changelog。(当前下载 v3.2.0)
```
git clone https://github.com/XiaoMuYi/k8s-deployment-cluster.git
```
**1.3查看相关配置**
```
$ cat hosts
$ egrep -v "^#|^$" group_vars/all.yml
$ egrep -v "^#|^$" group_vars/mgrs.yml
$ egrep -v "^#|^$" group_vars/osds.yml
$ egrep -v "^#|^$" site.yml
```
提示：这里只需要注释掉其他内容即可，我这里显示的是注释后的内容。
**1.4执行安装操作**
```
$ ansible-playbook -i hosts site.yml
```
重启服务：
```
systemctl restart ceph-mds.target
systemctl restart ceph-mgr.target
systemctl restart ceph-mon.target
systemctl restart ceph-osd.target
```
### 二、Ceph 文件系统
ceph 文件系统，需要部署 mds（元数据服务器）。基本依赖解决之后，就可以为 cephfs 创建 pool，并且至少需要两个rados池，一个用于数据，一个用于元数据。我们这里是ansible部署，所以很多过程已经实现。手动操作如下：
**2.1创建pool**
```
$ ceph osd pool create cephfs_data 128
pool 'cephfs_data' created
$ ceph osd pool create cephfs_metadata 128
pool 'cephfs_metadata' created
```
**2.2创建文件系统**
```
$ ceph fs new cephfs cephfs_metadata cephfs_data
$ ceph fs ls
```
**2.3启用dashboard**
```
$ ceph mgr module enable dashboard	# 启用dashboard模块
```
默认情况下，仪表板的所有HTTP连接均使用SSL/TLS进行保护。要快速启动并运行仪表板，可以使用以下内置命令生成并安装自签名证书:
```
$ ceph dashboard create-self-signed-cert
```
创建具有管理员角色的用户
```
$ ceph dashboard set-login-credentials admin admin
```
默认下，仪表板的守护程序(即当前活动的管理器)将绑定到TCP端口8443或8080
```
$ ceph mgr services
{
    "dashboard": "https://k8store01.ops.bj2.yongche.com:8443/"
}

$ ceph config-key set mgr/dashboard/server_addr 0.0.0.0
$ ceph config-key set mgr/dashboard/server_port 9000
$ systemctl restart ceph-mgr\@k8store01.service
```
**2.4设置pg数**
查看当前pg数
```
$ ceph osd pool get cephfs_data pg_num
pg_num: 8
$ ceph osd pool get cephfs_data pgp_num
pgp_num: 8

$ ceph osd pool get cephfs_metadata pg_num
pg_num: 8
$ ceph osd pool get cephfs_metadata pgp_num
pgp_num: 8
```
设置 cephfs_metadata pg_num 数
```
$ ceph osd pool set cephfs_metadata pg_num 32
$ ceph osd pool set cephfs_metadata pgp_num 32

$ ceph osd pool set cephfs_metadata pg_num 64
$ ceph osd pool set cephfs_metadata pgp_num 64

$ ceph osd pool set cephfs_metadata pg_num 128
$ ceph osd pool set cephfs_metadata pgp_num 128
```
设置 cephfs_metadata pg_num 数
```
$ ceph osd pool set cephfs_data pg_num 128
$ ceph osd pool set cephfs_data pgp_num 128

$ ceph osd pool set cephfs_data pg_num 512
$ ceph osd pool set cephfs_adata pgp_num 512

$ ceph osd pool set cephfs_data pg_num 1024
$ ceph osd pool set cephfs_data pgp_num 1024
```
### 三、kubernetes 集成 cephfs 
**3.1 创建ceph-secret这个k8s secret对象**
在ceph集群主机执行
```
$ ceph auth get-key client.admin
AQCCVSBcLK5nLhAAD3sehi8lweCwT+FJbvGSIA==
```
在 kubernetes master 主机执行
```
$ echo "AQCCVSBcLK5nLhAAD3sehi8lweCwT+FJbvGSIA==" > /tmp/secret
$ kubectl create ns cephfs
$ kubectl create secret generic ceph-secret-admin --from-file=/tmp/secret --namespace=cephfs
```
部署 CephFS provisioner Install with RBAC roles
```
$ git clone https://github.com/kubernetes-incubator/external-storage.git
$ cd external-storage/ceph/cephfs/deploy
$ NAMESPACE=cephfs
$ sed -r -i "s/namespace: [^ ]+/namespace: $NAMESPACE/g" ./rbac/*.yaml
$ kubectl -n $NAMESPACE apply -f ./rbac
```
参考链接：https://github.com/kubernetes-incubator/external-storage/tree/master/ceph/cephfs/deploy
**3.2创建动态PV/PVC**
创建一个storageclass
```
$ cd example
$ kubectl create -f local-class.yaml
$ kubectl get storageclass
NAME     PROVISIONER       AGE
cephfs   ceph.com/cephfs   10s
```
创建PVC使用cephfs storageClass动态分配PV
```
$ kubectl create -f local-claim.yaml
$ kubectl get pvc claim-local -n cephfs
NAME          STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
claim-local   Bound    pvc-127dd47b-08ee-11e9-9e4f-faf206331800   1Gi        RWX            cephfs         8s
```
创建Pod，并检查是否挂载cephfs卷成功
```
$ kubectl create -f test-pod.yaml -n cephfs
$ kubectl get pod cephfs-pv-pod1 -n cephfs
```
提示：secret和provisioner不在同一个namespace中的话，获取secret权限不够。

### 问题总结
**问题1. 每个 OSD 上的PG数量小于最小数目30个**
```
$ ceph health
HEALTH_WARN too few PGs per OSD (2 < min 30)

$ ceph -s
  cluster:
    id:     7e2de501-7b34-4121-a431-0776ee1cb004
    health: HEALTH_WARN
            too few PGs per OSD (2 < min 30)

  services:
    mon: 3 daemons, quorum k8store01,k8store02,k8store03
    mgr: k8store02(active), standbys: k8store01, k8store03
    mds: cephfs-1/1/1 up  {0=k8store03=up:active}, 2 up:standby
    osd: 18 osds: 18 up, 18 in

  data:
    pools:   2 pools, 16 pgs
    objects: 22  objects, 2.2 KiB
    usage:   18 GiB used, 33 TiB / 33 TiB avail
    pgs:     16 active+clean
```
从上面可以看到，提示说每个osd上的pg数量小于最小的数目30个。pgs为 16，因为是3副本的配置，所以当有18个osd的时候，每个osd上均分了16/18 * 3=2个pgs,也就是出现了如上的错误 小于最小配置30个。  

集群这种如果进行数据的存储和操作，会引发集群卡死，无法响应io，同事会导致大面积的 osd down。  

cephfs 需要用到两个pool ： fs_data 和fs_metadata。 在初次使用ceph 就能之前需要首先规划 集群一共承载多少存储业务，创建多少个 pool，最后得到每个存储应该分配多少个pg。  
参考链接：http://docs.ceph.com/docs/mimic/rados/operations/placement-groups/  

必须选择pg_num的值，因为它无法自动计算。以下是常用的一些值：
 * 少于 5 OSDs pg_num 设置为 128
 * 在 5 到 10 OSDs 之间 pg_num 设置为 512
 * 在 10 到 50 OSDs 之间 pg_num 设置为 1024

如果您有超过50个OSD，则需要了解折衷以及如何自己计算pg_num值要自己计算pg_num值，请使用pgcalc工具。参考链接：http://docs.ceph.com/docs/mimic/rados/operations/placement-groups/

关于cephfs_metadata pug_num正确配置参考：https://ceph.com/planet/cephfs-ideal-pg-ratio-between-metadata-and-data-pools/
**问题2. 设置 pg_num 提示错误**
```
$ ceph osd pool set cephfs_data pg_num 1024
Error E2BIG: specified pg_num 1024 is too large (creating 1016 new PGs on ~8 OSDs exceeds per-OSD max with mon_osd_max_split_count of 32)
```
结果出现这个错误，参考“http://www.selinuxplus.com/?p=782”，原来是一次增加的数量有限制。最后选择用暴力的方法解决问题：
```
$ ceph osd pool set cephfs_metadata pg_num 32
$ ceph osd pool set cephfs_metadata pgp_num 32

$ ceph osd pool set cephfs_metadata pg_num 64
$ ceph osd pool set cephfs_metadata pgp_num 64
```
**问题3. mgr 模块无法监听地址**
```
Jan 22 10:48:21 k8store01 ceph-mgr: ChannelFailures: error('No socket could be created',)
Jan 22 10:48:21 k8store01 ceph-mgr: [22/Jan/2019:10:48:21] ENGINE Bus STOPPING
Jan 22 10:48:21 k8store01 ceph-mgr: [22/Jan/2019:10:48:21] ENGINE HTTP Server cherrypy._cpwsgi_server.CPWSGIServer(('::', 9283)) already shut down
Jan 22 10:48:21 k8store01 ceph-mgr: [22/Jan/2019:10:48:21] ENGINE No thread running for None.
Jan 22 10:48:21 k8store01 ceph-mgr: [22/Jan/2019:10:48:21] ENGINE Bus STOPPED
Jan 22 10:48:21 k8store01 ceph-mgr: [22/Jan/2019:10:48:21] ENGINE Bus EXITING
Jan 22 10:48:21 k8store01 ceph-mgr: [22/Jan/2019:10:48:21] ENGINE Bus EXITED
Jan 22 10:48:21 k8store01 ceph-mgr: 2019-01-22 10:48:21.395684 7fad9120d700 -1 log_channel(cluster) log [ERR] : Unhandled exception from module 'prometheus' while running on mgr.k8store01: error('No socket could be created',)
Jan 22 10:48:21 k8store01 ceph-mgr: 2019-01-22 10:48:21.395699 7fad9120d700 -1 prometheus.serve:
Jan 22 10:48:21 k8store01 ceph-mgr: 2019-01-22 10:48:21.395701 7fad9120d700 -1 Traceback (most recent call last):
Jan 22 10:48:21 k8store01 ceph-mgr: File "/usr/lib64/ceph/mgr/prometheus/module.py", line 718, in serve
Jan 22 10:48:21 k8store01 ceph-mgr: cherrypy.engine.start()
Jan 22 10:48:21 k8store01 ceph-mgr: File "/usr/lib/python2.7/site-packages/cherrypy/process/wspbus.py", line 250, in start
Jan 22 10:48:21 k8store01 ceph-mgr: raise e_info
Jan 22 10:48:21 k8store01 ceph-mgr: ChannelFailures: error('No socket could be created',)
```
问题分析：
因为我把centos7的ipv6关闭了所以报错了，mgr默认是同时开启ipv4和ipv6，解决方案是指定使用ipv4启动mgr。
解决方案：
```
ceph config-key set mgr/dashboard/server_addr 0.0.0.0
ceph mgr module disable dashboard
ceph mgr module enable dashboard
```
