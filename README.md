# go-frameworks

项目启动命令
-------
项目依赖下载：go mod tidy

本地启动：make run

测试：make test

分层设计
-------
项目使用后端主要业务代码放在internal里，其主要结构如下
>internal
>>pkg
>>>db

>>>redis

>>app
>>>app1

>>>...

其中：

pkg放入公用的模块；

app里放置**业务代码**，采用DDD模式，分成四层架构：

>app1
>>interface

>>application

>>domain

>>infra

* interface定义api接收参数并校验规整,调用application
* application用接口调domain，主要用来业务聚合
* domain实现上层接口，完成单一领域逻辑，并用接口调用infra
* infa实现上层接口，使用基础设施

**接口与实现之间需要用wire绑定**


测试情况
-------

分层测试，测试分成
- api测试
- service测试

api测试包括哪些层

service测试包括

wire依赖注入
-------

wire总体包含两个包，当你要新增一个包，需要提供provide，需要在上层注册进去

其他信息
-------

项目使用go 1.6版本
