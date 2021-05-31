# beego_report_reader

This is a record about how to use Beego framework
## Bee
### Install

### Run

```
bee run
```

### pack

```
bee pack -be GOOS=windows
```
## 路由

### 获取url以及参数的方式

> 以下都全默认在controller下执行
获取当前请求的referer
```go
fmt.Println(this.Ctx.Request.Referer())
```
输出：http://localhost:8080/swagger/

获取当前uri:

```go
fmt.Println(this.Ctx.Request.RequestURI)
```
输出： /v1/weather/?longitude=1&latitude=2

获取query参数，形如 /?longitude=1&latitude=2

```go
fmt.Println(this.Ctx.Input.Query("longitude"))
fmt.Println(this.Ctx.Input.Query("latitude"))
```
ps:正常情况下，Query的key不应当以:开头，以免和Param里的key冲突

获取path参数， 形如http://localhost:8080/userinfo/{uid}这种
```go
fmt.Println(u.GetString(":uid"))
```
或者this.Ctx.Input.Param(":uid")

这里是字符串，如果是其他类型参考
GetString(key string) string
GetStrings(key string) []string
GetInt(key string) (int64, error)
GetBool(key string) (bool, error)
GetFloat(key string) (float64, error)

直接解析到 struct
### 注解路由

从 beego 1.3 版本开始支持了注解路由，用户无需在 router 中注册路由，只需要 Include 相应地 controller，然后在 controller 的 method 方法上面写上 router 注释（// @router）就可以了，详细的使用请看下面的例子：

// CMS API
type CMSController struct {
    beego.Controller
}

func (c *CMSController) URLMapping() {
    c.Mapping("StaticBlock", c.StaticBlock)
    c.Mapping("AllBlock", c.AllBlock)
}


// @router /staticblock/:key [get]
func (this *CMSController) StaticBlock() {

}

// @router /all/:key [get]
func (this *CMSController) AllBlock() {

}
可以在 router.go 中通过如下方式注册路由：

beego.Include(&CMSController{})
beego 自动会进行源码分析，注意只会在 dev 模式下进行生成，生成的路由放在 "/routers/commentsRouter.go" 文件中。

这样上面的路由就支持了如下的路由：

GET /staticblock/:key
GET /all/:key
其实效果和自己通过 Router 函数注册是一样的：

beego.Router("/staticblock/:key", &CMSController{}, "get:StaticBlock")
beego.Router("/all/:key", &CMSController{}, "get:AllBlock")
同时大家注意到新版本里面增加了 URLMapping 这个函数，这是新增加的函数，用户如果没有进行注册，那么就会通过反射来执行对应的函数，如果注册了就会通过 interface 来进行执行函数，性能上面会提升很多。

## 参数

c.GetString("id")
c.GetBool("id")
c.GetInt("id")
c.GetStrings("id")
## Issue

panic: err: chdir : The system cannot find the file specified.: stderr:

> 没有打包完整，go.mod