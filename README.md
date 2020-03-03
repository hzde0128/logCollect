# logAgent 日志收集服务

通过在运维平台上配置日志收集项，logAgent从etcd中获取要收集的日志信息从业务服务器读取日志信息，发往kafka，logTransfer负责从kafka读取日志，写入到Elasticsearch或者Prometheus，通过Kibana进行日志检索和Grafana进行可视化展示。

# v0.1.0版本实现的功能

- 读取日志文件

- 写入到kafka中

- 可以自行配置要收集的日志文件
