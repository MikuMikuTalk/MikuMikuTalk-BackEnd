
<a name="v0.1.0"></a>
## v0.1.0 (2024-10-25)

### Feat

* 配置日志格式
* 更新模板，更新日志,添加swagger脚本
* 使用自己的response进行封装,更新auth_api.api
* 更新main.go,测试im_auth模块是否可以正常连接数据库并运行
* 更新im_auth中的配置文件
* 更新目录结构，更新连接数据库的方法，在启动im_auth的服务的时候，自动初始化gorm.DB实例，注册在ServiceContext中
* 通过goctl api文件初始化im_auth模块，自动生成相关路由，controller和service和配置文件
* 编写各级表结构(用户表，群组表，信息表，验证表，聊天表)

### Fix

* 修复baseurl输错的问it
* 修复swagger无法请求的问题
* 修复无法创建表结构的问题

### Init

* 项目初始化
* 项目初始化

