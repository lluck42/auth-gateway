# gq_volunteer_platform

#### 介绍
auth gateway
使用反向代理方式进行登录、权限认证

admin.sql 是数据库文件
config.yml 是配置文件，填写所需代理的后端地址

1、管理员模块
2、登录模块
3、token验证模块
4、图形验证码

形如 http://0.0.0.0:3000/admin/xxx 不可用（网关管理模块）
其他地址可用

没有被注册的路由地址会走反向代理
header 中有 adminID 可用于业务

#### 软件架构
gin + gorm + mysql


#### 安装教程
git clone https://gitee.com/lluck42/gq_volunteer_platform.git
