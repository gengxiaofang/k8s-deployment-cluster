# 手动安装 kubernetes 1.13.0
主要目的是方便学习 kubernetes 的朋友能够实现快速安装，不在因为从安装入门到放弃 kubernetes。本文的基于 Centos 7.4 作为基础系统环境，其他相信操作参考其他文档

本文参考链接：https://github.com/opsnull/follow-me-install-kubernetes-cluster 


![ssl-success](images/dashboard.png)

* [环境说明]()
  * [1.基础环境准备](./01.基础环境准备.md)
* [etcd 部署]()
  * [1.安装etcd集群](./02.创建etcd集群.md)
* [master 节点部署]()
  * [1.部署api-server高可用组件](./03.keepalived+haproxy负载.md)
  * [2.二进制部署api-server组件](./04.kube-apiserver.md)
  * [3.二进制部署kube-controller-manager组件](./05.kube-controller-manager.md)
  * [4.二进制部署kube-scheduler组件](./06.kube-scheduler.md)

* [node 节点部署]()
  * [1.二进制部署kubelet组件](./07.kubelet-node部署.md)
  * [2.二进制部署kube-proxy组件](./08.kube-proxy部署.md)
  
* [附加组件部署]( )
  * [网络插件部署]( )
     * [1.calico](./09-1.calico网络设置.md)
     * [2.flannel](./09-2.flannel网络设置.md)
