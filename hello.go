package main

import (
	"os/exec"
	"runtime"

	"example.com/m/controllers"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func openBrowser(url string) error {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	return exec.Command(cmd, append(args, url)...).Start()
}

func main() {
	openBrowser("http://localhost:8888")
	web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	// web.Router("/", &MainController{})
	web.Router("/", &controllers.MainController{})
	web.Router("/sip/", &controllers.SipController{})
	web.Router("/sip/:id:string", &controllers.SipController{}, "get:GetSipContextDetail")
	web.Router("/device/", &controllers.DeviceController{})
	web.Router("/device/:id", &controllers.DeviceController{}, "get:ListDeviceStep")
	web.Router("/userinfo/", &controllers.UserInfoController{})
	web.Router("/file/down/:id", &controllers.DeviceController{}, "get:FileDown")
	web.SetStaticPath("/static", "static")
	web.Run()
}
