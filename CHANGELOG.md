
<a name="v0.1.2"></a>
## [v0.1.2](https://github.com/meowrain/im_server/compare/v0.1.1...v0.1.2) (2024-10-25)

### Feat

* 可以使用脚本调用数据库创建相关表结构

### Style

* auth_api.api引起的代码格式化


<a name="v0.1.1"></a>
## [v0.1.1](https://github.com/meowrain/im_server/compare/v0.1.0...v0.1.1) (2024-10-25)

### Feat

* 添加jwt封装功能: 可以生成和解析token
* 编写密码加密和比对功能，用于数据库存储用户信息
* Changelog功能添加


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

