# cephfs与kubernetes集成

**开启mgr prometheus端口**
```
$ ceph config-key set mgr/prometheus/server_addr 0.0.0.0
$ ceph mgr module disable prometheus
$ ceph mgr module enable prometheus
```
**部署prometheus 自动发现target**
```
$ cd k8s-deployment-cluster/manifests/prometheus/ceph
$ kuberctl create -f ./
```
