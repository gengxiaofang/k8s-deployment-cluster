### 一、简介
目前 ceph 提供对象存储（RADOSGW）、块存储RDB以及 CephFS 文件系统这 3 种功能。对于这3种功能介绍，分别如下：  
* 1.对象存储，也就是通常意义的键值存储，其接口就是简单的GET、PUT、DEL 和其他扩展，代表主要有 Swift 、S3 以及 gluster 等；
* 2.块存储，这种接口通常以 QEMU Driver 或者 Kernel Module 的方式存在，这种接口需要实现 Linux 的 Block Device 的接口或者由 QEMU 提供的 Block Driver 接口，如 Sheepdog，AWS 的 EBS，青云的云硬盘和阿里云的盘古系统，还有 Ceph 的 RBD（RBD是Ceph面向块存储的接口）。在常见的存储中 DAS、SAN 提供的也是块存储；
* 3.文件存储，通常意义是支持 POSIX 接口，它跟传统的文件系统如 Ext4 是一个类型的，但区别在于分布式存储提供了并行化的能力，如 Ceph 的 CephFS (CephFS是Ceph面向文件存储的接口)，但是有时候又会把 GlusterFS ，HDFS 这种非POSIX接口的类文件存储接口归入此类。当然 NFS、NAS也是属于文件系统存储。

**专业术语**
 * rados，全称 Reliable Autonomic Distributed Object Store，即可靠分布式对象存储，其作为在整个 Ceph 集群核心基础设施，向外部提供基本的数据操作。
 * mon，负载整个集群的运行状况监控，信息由集群成员的守护程序来提供，Ceph monitor map 主要包括 osd map、pg map、mds map、crush等。
 * osd，由物理磁盘驱动器在其Linux文件系统以及 ceph osd 服务组成。osd 将数据以对象的形式存储到集群中的每个节点的物理磁盘上，完成存储数据的工作绝大多数是由 osd daemon 进程实现。
 * mds，ceph 元数据，ceph 块设备和rdb不需要mds，mds 只为 cephfs 服务。
 * ados 块设备，它能够自动精简配置并可能调整大小，而且将数据分散存储在多个osd上。
 * librados，librados 库，为应用程序提供访问接口，同事也为块存储、对象存储、文件系统原生的接口。
 * radosgw，网关接口，提供对象存储服务。它使用librgw 和 librados 来实现允许应用程序与 ceph 对象存储建立连接。并且提供s3 和 swift 兼容的RESTful API接口。
 * crush，全称 Controlled Replication Under Scalable Hashing，它表示数据存储的分布式选择算法，ceph 的高性能、高可用就是采用这种算法实现。crush 算法取代了在元数据表中为每个客户端请求进行查找，它通过计算机系统中数据应该被写入或读出的位置。CRUSH能够感知基础架构，能够理解基础设施各个部件之间的关系。并且CRUSH保存数据的多个副本，这样即使一个故障域的几个组件都出现故障，数据依然可用。CRUSH 算是使得 ceph 实现了自我管理和自我修复。  
 
Ceph文件系统至少需要两个RADOS池，一个用于数据，一个用于元数据。

### 二、安装 ceph
#### 1. 安装ansible
```
yum -y install http://dist.yongche.com/centos/7/epel/x86_64/Packages/a/ansible-2.6.5-1.el7.noarch.rpm
```
提示：官方对 ansible 版本有明确的要求，目前支持 2.4 以及 2.6 不能太高也不能太新。

#### 2. 下载 ceph-ansible
到 https://github.com/ceph/ceph-ansible/releases 下载最新稳定版本，建议下载对应自己想安装的版本。如何知道自己下载的 ceph-ansible 支持自己想安装的 ceph 版本？只能查看与之对应的 Changelog。(当前下载 v3.2.0)
```
cd /home/yangsheng/ceph-ansible-3.2.0
```
#### 3. 添加hosts文件
```
$ cat ./hosts
[mons]
172.17.3.32
172.17.3.33
172.17.3.34

[osds]
172.17.3.32
172.17.3.33
172.17.3.34

[mgrs]
172.17.3.32
172.17.3.33
172.17.3.34

[mdss]
172.17.3.32
172.17.3.33
172.17.3.34

[clients]
172.17.80.29
172.17.80.30
172.17.80.31

$ cp group_vars/all.yml.sample group_vars/all.yml
$ cp group_vars/osds.yml.sample group_vars/osds.yml
$ 
$ cp site.yml.sample site.yml
```
#### 4. 配置全局变量
```
$ egrep -v "^#|^$" group_vars/all.yml
---
dummy:
ceph_origin: repository
ceph_repository: community
ceph_mirror: https://mirrors.aliyun.com/ceph
ceph_stable_key: https://mirrors.aliyun.com/ceph/keys/release.asc
ceph_stable_release: mimic
ceph_stable_repo: "{{ ceph_mirror }}/rpm-{{ ceph_stable_release }}"
fsid: 17ffc828-5d8c-4937-a5bb-f6adb2384d20
generate_fsid: true
ceph_conf_key_directory: /etc/ceph
cephx: true
monitor_interface: bond0
public_network: 172.17.0.0/16
cluster_network: 172.17.0.0/16
ceph_conf_overrides:
  global:
    rbd_default_features: 7
    auth cluster required: cephx
    auth service required: cephx
    auth client required: cephx
    osd journal size: 2048
    osd pool default size: 3
    osd pool default min size: 1
    mon_pg_warn_max_per_osd: 1024
    osd pool default pg num: 1024
    osd pool default pgp num: 1024
    max open files: 131072
    osd_deep_scrub_randomize_ratio: 0.01
  mon:
    mon_allow_pool_delete: true

  client:
    rbd_cache: true
    rbd_cache_size: 335544320
    rbd_cache_max_dirty: 134217728
    rbd_cache_max_dirty_age: 10

  osd:
    osd mkfs type: xfs
    ms_bind_port_max: 7100
    osd_client_message_size_cap: 2147483648
    osd_crush_update_on_start: true
    osd_deep_scrub_stride: 131072
    osd_disk_threads: 4
    osd_map_cache_bl_size: 128
    osd_max_object_name_len: 256
    osd_max_object_namespace_len: 64
    osd_max_write_size: 1024
    osd_op_threads: 8
    osd_recovery_op_priority: 1
    osd_recovery_max_active: 1
    osd_recovery_max_single_start: 1
    osd_recovery_max_chunk: 1048576
    osd_recovery_threads: 1
    osd_max_backfills: 4
    osd_scrub_begin_hour: 23
    osd_scrub_end_hour: 7

$ egrep -v "^#|^$" group_vars/mgrs.yml
---
dummy:
ceph_mgr_modules: [status,dashboard]
```

#### 5. osds.yml 内容如下
```
$ egrep -v "^#|^$" group_vars/osds.yml
---
dummy:
devices:
  - /dev/sdb
  - /dev/sdc
  - /dev/sdd
  - /dev/sde
  - /dev/sdf
  - /dev/sdg
osd_scenario: collocated
osd_objectstore: bluestore
```
#### 6. site.yml 内容如下
```
$ egrep -v "^#|^$" site.yml
---
- hosts:
  - mons
  - osds
  - mdss
  - clients
  - mgrs
```
提示：这里只需要注释掉其他内容即可，我这里显示的是注释后的内容。

#### 7. 执行安装操作
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
### 三、Ceph 文件系统
ceph 文件系统，需要部署 mds（元数据服务器）。基本依赖解决之后，就可以为 cephfs 创建 pool，并且至少需要两个rados池，一个用于数据，一个用于元数据。我们这里是ansible部署，所以很多过程已经实现。手动操作如下：
**创建pool**
```
$ ceph osd pool create cephfs_data 128
pool 'cephfs_data' created
$ ceph osd pool create cephfs_metadata 128
pool 'cephfs_metadata' created
```
**创建文件系统**
```
$ ceph fs new cephfs cephfs_metadata cephfs_data
```
**查看创建好的Ceph FS**
```
$ ceph fs ls
```
#### 1. 启用 dashboard
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
**查看当前pg数**
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
**设置 cephfs_metadata pg_num 数**
```
$ ceph osd pool set cephfs_data pg_num 128
$ ceph osd pool set cephfs_data pgp_num 128

$ ceph osd pool set cephfs_data pg_num 512
$ ceph osd pool set cephfs_adata pgp_num 512

$ ceph osd pool set cephfs_data pg_num 1024
$ ceph osd pool set cephfs_data pgp_num 1024
```

### 四、kubernetes 集成 cephfs 

#### 4.1 创建ceph-secret这个k8s secret对象
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

创建一个storageclass
```
$ cd example
$ kubectl create -f local-class.yaml
$ kubectl get storageclass
NAME     PROVISIONER       AGE
cephfs   ceph.com/cephfs   10s

创建PVC使用cephfs storageClass动态分配PV
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

**问题总结**
* 问题1. 每个 OSD 上的PG数量小于最小数目30个
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
* 问题2. 设置 pg_num 提示错误
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
* 问题3. mgr 模块无法监听地址
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
ceph config-key set mgr/prometheus/server_addr 0.0.0.0
ceph mgr module disable dashboard
ceph mgr module enable dashboard
ceph mgr module disable prometheus
ceph mgr module enable prometheus
```
