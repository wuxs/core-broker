# tkeel core-broker

这是一个封装了对 Core 基础功能，从而实现一些需要复杂操作从而满足用户需求的一个代理服务。目前提供了：Websocket 服务、 Subscribe 服务。供平台侧使用。

## 依赖
- 集群模式下的 tKeel 平台
- 一个 MySQL 服务
- tKeel Core 服务
- tKeel Device 服务
- dapr 边车模式开启 core-broker 服务

## 环境配置
以下是该服务用到的环境变量：
```bash
// 该变量用于指定数据订阅生成的 amqp 服务地址指向
export AMQP_SERVER=amqp://tkeel.io:5672

// 用于定义该服务连接的 MySQL 配置 DSN
export DSN=user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
```
## Build 
```bash
make build
```