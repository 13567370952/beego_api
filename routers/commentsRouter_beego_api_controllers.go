package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["beego_api/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_api/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["beego_api/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_api/controllers:UserController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beego_api/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_api/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beego_api/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_api/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["beego_api/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_api/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
