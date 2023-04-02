# vizzy

vizzy是一个数据可视化项目，用于数据大屏展示。本README将介绍如何使用和配置项目。

## Feature

* 打包体积轻量，仅20MB。
* 使用无头浏览器lorca，可自定义Chrome和JavaScript之间的交互。
* 支持交叉编译到Windows和Mac系统。
* 离线状态下也可以正常运行。
* 可以运行本地服务，减轻服务器压力。
* 编译速度快，运行性能优秀。

## 依赖项

该项目的依赖项如下：

* Go 1.20+
* node 14.8+

## 安装

1. git clone

2. cd frontend && yarn

3. 安装完成后，打开终端并运行以下命令：

```bash
./start.sh
```

## 发布

```bash
./build.sh
```


## 目录结构

```text
.
├── README.md
├── bootstrap 启动程序
├── build.sh 生产构建脚本
├── common 公共模块
├── dao 
├── dto
├── entity
├── frontend 前端
├── go.mod
├── go.sum
├── job 定时任务
├── lorca 
├── main.db
├── main.go
├── resources 前端生产包
├── service 接口
├── start.sh 开发环境构建脚本
└── start.sh 开发环境构建脚本
```

## why go？

![img.png](wonderful-go.png)

## Go组件库
| 特性     | 地址           |
|--------|--------------|
| gin    | web服务        |
| gorm   | ORM（对象关系映射）库 |
| sqlite | 数据库          |
| cron   | 定时器          |
| lorca  | 无头浏览器        |
| gzip   | gzip         |
| logrus | 日志           |

## JS组件库
| 特性                     | 地址    |
|------------------------|-------|
| vite                   | 打包工具  |
| react                  | spa   |
| echarts                | 图表    |
| rc-table               | 表格    |
| react-use              | hooks |
| @tanstack/react-query" | 数据管理库 |
| axios                  | 请求    |

## 日志

开发环境输出在控制台

生产环境输出在app.log