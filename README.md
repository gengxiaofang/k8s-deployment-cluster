# 手动安装 kubernetes v1.13.0 高可用集群


   主要目的是方便学习 kubernetes 的朋友能够实现快速安装，不在因为从安装入门到放弃 kubernetes。本文的基于 Centos 7.4 作为基础系统环境，其他详细操作请参考其他文档。

### 目录

* [实验环境]()
  * [1.安装前准备](./01.基础环境准备.md)
  * [2.etcd集群准备](./02.创建etcd集群.md)
  * [3.通过keepalived+haproxy实现apiserver高可用](./03.keepalived+haproxy负载.md)
  * [4.etcd备份](./ops/etcd/etcd_cluster_backup_recovery.md)
* [master 节点部署]()
  * [1.二进制部署api-server组件](./04.kube-apiserver.md)
  * [2.二进制部署kube-controller-manager组件](./05.kube-controller-manager.md)
  * [3.二进制部署kube-scheduler组件](./06.kube-scheduler.md)
* [node 节点部署]()
  * [1.二进制部署kubelet组件](./07.kubelet-node部署.md)
  * [2.二进制部署kube-proxy组件](./08.kube-proxy部署.md)
* [附加组件部署]( )
  * [网络插件部署](二选一)
     * [1.calico](./09-1.calico网络设置.md)
     * [2.flannel](./09-2.flannel网络设置.md)
  * [CoreDNS插件部署](./10.coredns.md)
  * [Dashborad插件部署](./11.dashboard.md)
  * [Heapster插件部署](./12.heapster.md)
  * [metrics-server插件部署](./13.metrics-server.md)
* [其他组件部署]( )
  * [helm](./15.helm安装部署.md)
  * [HPA自动伸缩](./14.hpa.md)
* [prometheus-operator]( )
  * [prometheus-operator部署](./16.通过helm部署prometheusoperator监控.md)
