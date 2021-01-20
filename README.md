# goim

基于gnet框架完成的IM框架

1. api-gateway 网关
2. file_server 文件服务器
3. im-client golang命令行客户端
4. file_server 文件服务器
5. im-connect im连接层
6. im-core im核心代码
7. logic-user user业务
8. protocols protobuf协议文件
9. utils 通用工具包

启动：
1. 直接执行make all指令，生成可执行文件，最后启动网关即可；
2. 运行im-client项目，这是一个Java完成的客户端，地址：https://github.com/lsm1998/im-client

[客户端](https://github.com/lsm1998/im-client) 使用Java语言编写，基于swing+SpringBoot
