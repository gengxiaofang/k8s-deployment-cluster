# cephfs与kubernetes集成
---

**ansible安装**
```
$ git clone https://github.com/XiaoMuYi/k8s-deployment-cluster.git
$ yum -y install http://dist.yongche.com/centos/7/epel/x86_64/Packages/a/ansible-2.6.5-1.el7.noarch.rpm
```
提示：到`https://github.com/ceph/ceph-ansible/releases`下载最新稳定版本，并且官方对ansible目前只支持2.4以及2.6版本。

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

$ ceph osd pool set cephfs_data pg_num 512
$ ceph osd pool set cephfs_adata pgp_num 512

$ ceph osd pool set cephfs_data pg_num 1024
$ ceph osd pool set cephfs_data pgp_num 1024
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

