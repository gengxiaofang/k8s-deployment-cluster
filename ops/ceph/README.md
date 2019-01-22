# cephfs与kubernetes集成

**ansible安装**
```
$ git clone https://github.com/XiaoMuYi/k8s-deployment-cluster.git
$ yum -y install http://dist.yongche.com/centos/7/epel/x86_64/Packages/a/ansible-2.6.5-1.el7.noarch.rpm
```
提示：https://github.com/ceph/ceph-ansible/releases 下载最新稳定版本，并且官方对ansible目前只支持2.4以及2.6版本。

**执行安装**
```
$ cd k8s-deployment-cluster/ops/ceph/ceph-ansible-3.2.0
$ ansible-playbook -i hosts site.yml
```
**设置pg数**  
设置 cephfs_data pg_num 数
```
$ ceph osd pool set cephfs_data pg_num 128
$ ceph osd pool set cephfs_data pgp_num 128

$ ceph osd pool set cephfs_data pg_num 1024
$ ceph osd pool set cephfs_data pgp_num 1024
```
设置 cephfs_metadata pg_num 数
```
$ ceph osd pool set cephfs_metadata pg_num 32
$ ceph osd pool set cephfs_metadata pgp_num 32

$ ceph osd pool set cephfs_metadata pg_num 128
$ ceph osd pool set cephfs_metadata pgp_num 128
```

**创建ceph-secret这个k8s secret对象**
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
**创建动态PV/PVC**
创建一个storageclass
```
$ cd ./k8s-deployment-cluster/ops/ceph/example/
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
**创建测试Pod并检查是否挂载cephfs卷成功**
```
$ kubectl create -f test-pod.yaml -n cephfs
$ kubectl get pod cephfs-pv-pod1 -n cephfs
```
提示：`secret` 和 `provisioner` 不在同一个 `namespace` 中的话，获取 `secret` 权限不够。
