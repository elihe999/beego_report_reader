package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["example.com/m/controllers:DeviceController"] = append(beego.GlobalControllerRouter["example.com/m/controllers:DeviceController"],
        beego.ControllerComments{
            Method: "Index",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["example.com/m/controllers:DeviceController"] = append(beego.GlobalControllerRouter["example.com/m/controllers:DeviceController"],
        beego.ControllerComments{
            Method: "IndexAbout",
            Router: "/about",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
