# Sbyter Man

#### 介绍
本项目设计技术包括：Golang，Gin，Gorm，MySQL，Redis，Kafka，Cron。实现了”简约版“抖音的基本功能，包括的扩展功能有：登录注册的邮箱验证、
kakfa的异步任务等。

#### 项目文件夹说明
| 文件夹名称 | 功能                |
| ---------- | ------------------- |
| configs    | 配置文件            |
| DDL    | 数据库定义语句            |
| docs       | swager api文档/其他 |
| global     | 全局变量            |
| internal   | 内部功能实现        |
| pkg        | 公共模块实现        |
| scripts    | 脚本实现            |
| storage    | 缓存/日志文件存储   |

#### 软件架构说明
![img.png](img.png)
#### 演示视频
https://sbyterman.oss-cn-hangzhou.aliyuncs.com/video/file_v2_09cbcc8a-8ef9-486d-900f-7514bca3b53g_2033260184.mp4
#### 安装教程

```go
//安装项目所需模块 
go mod tidy
```

#### 使用说明

```go
go run main.go
```

