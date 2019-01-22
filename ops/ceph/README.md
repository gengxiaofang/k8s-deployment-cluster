# cephfs与kubernetes集成
---

**ansible安装**
```
$ git clone https://github.com/XiaoMuYi/k8s-deployment-cluster.git
$ yum -y install http://dist.yongche.com/centos/7/epel/x86_64/Packages/a/ansible-2.6.5-1.el7.noarch.rpm
```
**提示**：到`https://github.com/ceph/ceph-ansible/releases`下载最新稳定版本，并且官方对ansible版本目前只支持2.4以及2.6。
**执行安装**
```
$ cd k8s-deployment-cluster/ops/ceph/ceph-ansible-3.2.0
$ ansible-playbook -i hosts site.yml
```
