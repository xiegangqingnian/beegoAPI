// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"zlt/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/info",
			beego.NSInclude(
				&controllers.GetInfoController{},
			),
		),
		beego.NSNamespace("/block",
			beego.NSInclude(
				&controllers.GetBlockController{},
			),
		),
		beego.NSNamespace("/blocks",
			beego.NSInclude(
				&controllers.GetBlocksController{},
			),
		),
		beego.NSNamespace("/trx",
			beego.NSInclude(
				&controllers.GetTrxController{},
			),
		),
		beego.NSNamespace("/addr",
			beego.NSInclude(
				&controllers.GetAddrController{},
			),
		),
		beego.NSNamespace("/trxjsontobin",
			beego.NSInclude(
				&controllers.TrxJsonToBinController{},
			),
		),
		beego.NSNamespace("/newaddr",
			beego.NSInclude(
				&controllers.NewAddrController{},
			),
		),
		beego.NSNamespace("/sendtrx",
			beego.NSInclude(
				&controllers.SendTrxController{},
			),
		),
		beego.NSNamespace("/trxheight",
			beego.NSInclude(
				&controllers.TrxHeightController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
