# 开发文档

## 相关文档

[抖音项目介绍](https://bytedance.feishu.cn/docx/doxcnbgkMy2J0Y3E6ihqrvtHXPg)

[抖音极简版开发文档-Sbyter Man](https://qjz04nkgx0.feishu.cn/wiki/wikcnKiaUte4QMEop1cL3wQwAge#)

[抖音极简版API](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18875092)

[Kitex官方文档](https://www.cloudwego.io/zh/docs/kitex/)

1. 生成代码:

```shell
cd service/user && kitex -service service-user -type protobuf -module douyin_service -I ../../idl user.proto
cd service/video && kitex -service service-video -type protobuf -module douyin_service -I ../../idl video.proto
```

## 启动

1. 启动etcd服务

```shell
/tmp/etcd-download-test/etcd
```

2. 运行单个服务

```shell
cd service/user && go run main.go handler.go
cd service/video && go run main.go handler.go
```