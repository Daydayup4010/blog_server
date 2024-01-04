# blog_server
### 一.项目目录
api: 存放各个接口的handler func

config: 存放mysql, logger, server, cloud的配置信息结构体, 映射settings.yaml

core: gorm, logger等init方法

global: 全局变量

log: 存放gin 接口请求日志

middleware: 云服务, 鉴权token, logger等gin中间件

models: 各模块的sql处理

routers: 接口路由

utils: 
    errmsg: 接口返回code对应的message
    validator: 请求参数数据验证

