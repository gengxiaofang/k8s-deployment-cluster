# 手动安装 kubernetes v1.14.x 高可用集群
随着 kubernetes 的火热程度，目前已经成为开发、运维必要的技术之一。但是了解到在初学的过程中，大部分的部署方式是通过 kubeadm 以及 ansible 等工具实现的快速部署。往往无法了解 kubernetes 各个组件的工作原理，这样对于想深入了解它的同学们来说是一种障碍。因此我希望初学者通过二进制部署方式来进行手动部署，这样可以熟悉 kubernetes 集群之间各组件的关系以及依赖。

### 安装步骤
 * [01.基础环境准备](./01.基础环境准备.md)
 * [02.安装etcd集群准备](./02.创建etcd集群.md)
 * [03.安装keepalived以及haproxy](./03.keepalived+haproxy负载.md)
 * [04.部署api-server组件](./04.kube-apiserver.md)
 * [05.部署kube-controller组件](./05.kube-controller-manager.md)
 * [06.部署kube-scheduler组件](./06.kube-scheduler.md)
 * [07.部署kubelet组件](./07.kubelet-node部署.md)
 * [08.部署kube-proxy组件](./08.kube-proxy部署.md)
 * [09.网络flannel部署](./09-2.flannel网络设置.md)
 * [10.CoreDNS插件部署](./10.coredns.md)
 * [11.Dashborad部署](./11.dashboard.md)
 * [12.Heapster部署](./12.heapster.md)
 * [13.metrics-server部署](./13.metrics-server.md)
 * [14.helm安全安装](./15.helm安装部署.md)
 * [15.通过prometheus-operator部署k8s监控](./16.通过helm部署prometheus-operator监控.md)


### 其他说明
 * [HPA自动伸缩](./14.HorizontalPodAutoscaling.md)
 * [关于etcd备份](./ops/etcd/etcd_cluster_backup_recovery.md)
 * [网络calico部署](./09-1.calico网络设置.md)