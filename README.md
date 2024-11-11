# im_server
即时通信程序后端

# 系统服务端口

im_server
  api
    im_gateway   网关            127.0.0.1:9000
    im_auth      认证服务         127.0.0.1:20021
    im_user      用户服务         127.0.0.1:20022
    im_chat      对话服务         127.0.0.1:20023
    im_group     群聊服务         127.0.0.1:20024
    im_file      文件服务         127.0.0.1:20025
    im_settings  系统服务         127.0.0.1:20026
  rpc
    im_user      用户rpc服务         127.0.0.1:30022
    im_chat      对话rpc服务         127.0.0.1:30023
    im_file      文件rpc服务         127.0.0.1:30024
    im_group     群聊rpc服务         127.0.0.1:30025
    