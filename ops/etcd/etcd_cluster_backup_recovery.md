### etcd 集群高可用说明
etcd的设计成能够承受集群中机器故障。  
etcd集群能够自动的从临时故障（例如，机器重新引导）中恢复，并且容忍N个成员集群中最多（N-1）/2 个永久故障。当一个集群成员永久性地失败时，无论是由于硬件故障还是磁盘损坏，它都会失去对集群的访问。如果集群永久性地丢失了超过（N-1）/2 个成员，那么它就会灾难性地失败，无法挽回地丢失了法定人数。一旦丢失了仲裁，集群就无法达成共识，因此无法继续接受更新。为了从灾难性故障中恢复，etcd v3提供了快照和恢复工具，以便集群故障的情况下重新创建集群。

### 快照备份
```
$ mkdir /data/etcd_backup -p
$ export ETCDCTL_API=3;
$ etcdctl --cacert=/etc/kubernetes/ssl/ca.pem \
     --cert=/etc/etcd/ssl/etcd.pem \
     --key=/etc/etcd/ssl/etcd-key.pem \
     --endpoints https://192.168.133.128:2379,https://192.168.133.129:2379,https://192.168.133.130:2379 \
     snapshot save /data/etcd_backup/$(date +%Y%m%d_%H%M%S)_snapshot.db
Snapshot saved at /data/etcd_backup/20181215_145639_snapshot.db
$ ls /data/etcd_backup/
20181215_145639_snapshot.db
```
### 回复快照备份
模拟 etcd 数据丢失
```
$ rm -rf /var/lib/etcd
```
停止kube-apiserver
```
[root@k8s-m01 ~]# systemctl stop kube-apiserver
[root@k8s-m01 ~]# systemctl status kube-apiserver

[root@k8s-m02 ~]# systemctl stop kube-apiserver
[root@k8s-m02 ~]# systemctl status kube-apiserver

[root@k8s-m02 ~]# systemctl stop kube-apiserver
[root@k8s-m02 ~]# systemctl status kube-apiserver
```
**提示：** 停止 apiserver 对整个现有运行pod无任何影响;
停止etcd服务
```
[root@k8s-m01 ~]# systemctl stop etcd
[root@k8s-m01 ~]# systemctl status etcd

[root@k8s-m02 ~]# systemctl stop etcd
[root@k8s-m02 ~]# systemctl status etcd

[root@k8s-m02 ~]# systemctl stop etcd
[root@k8s-m02 ~]# systemctl status etcd

```
恢复etcd数据
```
# 拷贝 k8s-m01 备份数据到其他2台etcd主机
$ scp /data/etcd_backup/20181215_145639_snapshot.db 192.168.133.129:/tmp/                                                               
$ scp /data/etcd_backup/20181215_145639_snapshot.db 192.168.133.130:/tmp/

# 开始恢复
export ETCDCTL_API=3
[root@k8s-m01 ~]# etcdctl --name=k8s-m01 \
   --endpoints=https://192.168.133.128:2379 \
   --cert=/etc/etcd/ssl/etcd.pem \
   --key=/etc/etcd/ssl/etcd-key.pem \
   --cacert=/etc/kubernetes/ssl/ca.pem \
   --initial-advertise-peer-urls=https://192.168.133.128:2380 \
   --initial-cluster-token=etcd-cluster \
   --initial-cluster=k8s-m01=https://192.168.133.128:2380,k8s-m02=https://192.168.133.129:2380,k8s-m03=https://192.168.133.130:2380 \
   --data-dir=/var/lib/etcd snapshot restore /data/etcd_backup/20181215_145639_snapshot.db
[root@k8s-m01 ~]# chown etcd:etcd -R /var/lib/etcd

[root@k8s-m02 ~]# etcdctl --name=k8s-m02 \
   --endpoints=https://192.168.133.129:2379 \
   --cert=/etc/etcd/ssl/etcd.pem \
   --key=/etc/etcd/ssl/etcd-key.pem \
   --cacert=/etc/kubernetes/ssl/ca.pem \
   --initial-advertise-peer-urls=https://192.168.133.129:2380 \
   --initial-cluster-token=etcd-cluster \
   --initial-cluster=k8s-m01=https://192.168.133.128:2380,k8s-m02=https://192.168.133.129:2380,k8s-m03=https://192.168.133.130:2380 \
   --data-dir=/var/lib/etcd snapshot restore /tmp/20181215_145639_snapshot.db
[root@k8s-m03 ~]# chown etcd:etcd -R /var/lib/etcd

[root@k8s-m03 ~]# etcdctl --name=k8s-m03 \
   --endpoints=https://192.168.133.130:2379 \
   --cert=/etc/etcd/ssl/etcd.pem \
   --key=/etc/etcd/ssl/etcd-key.pem \
   --cacert=/etc/kubernetes/ssl/ca.pem \
   --initial-advertise-peer-urls=https://192.168.133.130:2380 \
   --initial-cluster-token=etcd-cluster \
   --initial-cluster=k8s-m01=https://192.168.133.128:2380,k8s-m02=https://192.168.133.129:2380,k8s-m03=https://192.168.133.130:2380 \
   --data-dir=/var/lib/etcd snapshot restore /tmp/20181215_145639_snapshot.db
[root@k8s-m03 ~]# chown etcd:etcd -R /var/lib/etcd
```
启动etcd
```
systemctl start etcd && systemctl status etcd
```

启动kube-apiserver
```
systemctl restart kube-apiserver && systemctl status kube-apiserver
```
### 查看恢复的数据
```
$ export ETCDCTL_API=3
$ etcdctl \
     --cacert=/etc/kubernetes/ssl/ca.pem \
     --cert=/etc/etcd/ssl/etcd.pem \
     --key=/etc/etcd/ssl/etcd-key.pem \
     --endpoints=https://192.168.133.128:2379,https://192.168.133.128:2379,https://192.168.133.128:2379 \
     get /registry/ --prefix --keys-only
```
