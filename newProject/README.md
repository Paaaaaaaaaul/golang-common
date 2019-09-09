# newProject

这是一个工具包，用于快速生成一套通用的web或者rpc框架代码

使用方法：
```
选取对应系统的可执行文件(或自己编译)
newProject.exe适用windows系统
newProject适用linux系统

去到项目要创建的路径上，执行 newProject --projectName yourProjectName 即可
```


使用该工具创建的项目结构如下：

```
 --- ProjectName
     --- config
     --- data
         --- ini
         --- json
     --- exception
     --- gRpcHandler
     --- handler
     --- model
     --- router
     --- service
     main.go
```