# newProject

这是一个工具包，用于快速生成一套通用的web或者rpc框架代码



# 相关依赖
```go
github.com/becent/golang-common
```



# 使用方法：
```
以windows为例

先编译执行文件 go build -o newProject.exe main.go

去到项目要创建的路径上，执行 newProject.exe --projectName yourProjectName 即可
```


使用该工具创建的项目结构如下：

```
 --- ProjectName
     --- config        // 配置文件管理
     --- data          // 数据文件夹，例如一些配置文件、业务数据文件
         --- ini
         --- json
     --- exception     // 异常管理
     --- gRpcHandler   // gRpc对应的Controller实现
     --- handler       // http对应的Controller实现
     --- model         // model层
     --- router        // 路由层
     --- service       // 服务层
     main.go
```

# HelloWorld

下面我们快速开发一个hello world接口

1. 先注册路由，修改router/router.go文件

```go
engine.POST("/hello", handler.HelloWorldAction)
```

该行代码注册了一个 /hello 的路由，当用户访问该路由时，将由handler.HelloWorldAction(*gin.Context)函数响应

2. 实现handler.HelloWorldAction，编辑handler/example.go文件，添加如下代码

```go
func HelloWorldAction(c *gin.Context) {
	h := gin_handler.DefaultHandler(c)

	var (
		name = h.StringParam("name")
	)

	h.SuccessResponse(fmt.Sprintf("hello %v", name))
}
```
该代码实现了HelloWorldAction函数，函数接收一个叫做name的参数，并以hello $name作为输出

3. 编译文件，在main.go目录下执行 go build -o main.exe main.go获取main.exe

4. 修改配置文件

```go
[system]
# App名
app_name = helloWorld

# 服务器监听的web端口
http_listen_port = 8080


[logs]
# 日志目录
dir = /data/logs/hello-world/

# 日志名称
file_name = hello-world.log

# 日志保留时间，单位天
keep_days = 30

# 日志切割间隔，单位小时
rate_hours = 24
```

5. 启动main.exe

6. 在浏览器上发起请求 http://127.0.0.1:8080/hello?name=bob，看到响应 hello bob，至此hello world服务开发完成


