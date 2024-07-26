官网https://prometheus.io/download/ 下载prometheus-2.53.1.darwin-amd64.tar.gz 解压到本地运行

运行：./prometheus --config.file=prometheus.yml

http://localhost:9090



安装pushgateway

从官网https://prometheus.io/download/ 下载pushgateway-1.9.0.darwin-amd64.tar.gz 解压到本地运行
./pushgateway

http://localhost:9091

echo 'test_metric_1m{a="a", b="b"} 3.15' | curl --data-binary @- http://localhost:9091/metrics/job/testJob
echo 'test_metric_1m{a="a", b="c"} 3.16' | curl --data-binary @- http://localhost:9091/metrics/job/testJob


分布式链路追踪
https://www.jaegertracing.io/download/
./jaeger-all-in-one  all-in-one 数据存在内存中，用于学习


链路追踪生产者（ Tracer Provider ）

资源（ Resource ）

导出器（ Exporter ）


深入理解 OpenTelemetry 中的 Tracer Name 与 Service Name 及其相互关系

引言
在分布式系统的观测中，OpenTelemetry 提供了强大的工具和框架来收集和分析遥测数据。理解 OpenTelemetry 中的关键概念，如 Tracer Name 和 Service Name，对于有效利用这个框架非常重要。本文将详细介绍这两个概念以及它们之间的关系。


1. Tracer Name 的概念
   Tracer Name 在 OpenTelemetry 中用于标识生成遥测数据的 Tracer 实例。它通常与特定的库或应用程序模块相关联。Tracer Name 的主要作用是：

模块标识：帮助识别遥测数据是由应用程序的哪个部分或库生成的。
便于管理和调试：在复杂的应用程序中，不同模块可能需要独立的追踪策略，Tracer Name 在此中起到了关键作用。
2. Service Name 的概念
   Service Name 是 OpenTelemetry 中用来标识整个服务或应用程序的名称。在微服务架构中，每个服务一般都有自己的 Service Name。Service Name 的主要用途包括：

服务识别：在微服务架构中，Service Name 用于区分不同的服务。
数据聚合：在监控和分析工具中，Service Name 用于聚合和对比来自不同服务的遥测数据。
3. Tracer Name 与 Service Name 的关系
   虽然 Tracer Name 和 Service Name 都用于标识和区分遥测数据的来源，但它们在应用中扮演不同的角色：

层级差异：Service Name 通常表示更高层级的应用程序或服务，而 Tracer Name 则更多关联于应用程序内的具体模块或库。
关联性：Tracer Name 通常用于细粒度的追踪，如库或模块级别，而 Service Name 用于表示整个服务或应用程序的性能和状态。
4. 实际应用中的例子
   在一个电商平台的微服务架构中，Service Name 可能是 payment-service，代表处理支付的服务。而在这个服务中，可能有一个用于信用卡处理的模块，其 Tracer Name 可以是 credit-card-processing。在这个例子中，Service Name 用于标识整个支付服务，而 Tracer Name 用于标识服务中具体的模块。

结论
在 OpenTelemetry 中，Tracer Name 和 Service Name 是理解和操作遥测数据的关键概念。正确地理解和使用这两个概念，对于构建高效、可维护且易于监控的分布式系统至关重要。清晰的命名策略有助于提高系统的可观测性，从而优化性能和快速定位问题。