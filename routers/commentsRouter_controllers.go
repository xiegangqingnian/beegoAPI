package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["zlt/controllers:GetAddrController"] = append(beego.GlobalControllerRouter["zlt/controllers:GetAddrController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"POST"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zlt/controllers:GetBlockController"] = append(beego.GlobalControllerRouter["zlt/controllers:GetBlockController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"POST"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zlt/controllers:GetBlocksController"] = append(beego.GlobalControllerRouter["zlt/controllers:GetBlocksController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"POST"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zlt/controllers:GetInfoController"] = append(beego.GlobalControllerRouter["zlt/controllers:GetInfoController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"GET"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zlt/controllers:GetTrxController"] = append(beego.GlobalControllerRouter["zlt/controllers:GetTrxController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"POST"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zlt/controllers:NewAddrController"] = append(beego.GlobalControllerRouter["zlt/controllers:NewAddrController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"POST"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zlt/controllers:ObjectController"] = append(beego.GlobalControllerRouter["zlt/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zlt/controllers:ObjectController"] = append(beego.GlobalControllerRouter["zlt/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zlt/controllers:ObjectController"] = append(beego.GlobalControllerRouter["zlt/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zlt/controllers:ObjectController"] = append(beego.GlobalControllerRouter["zlt/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zlt/controllers:ObjectController"] = append(beego.GlobalControllerRouter["zlt/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zlt/controllers:SendTrxController"] = append(beego.GlobalControllerRouter["zlt/controllers:SendTrxController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"POST"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zlt/controllers:TrxHeightController"] = append(beego.GlobalControllerRouter["zlt/controllers:TrxHeightController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"GET"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zlt/controllers:TrxJsonToBinController"] = append(beego.GlobalControllerRouter["zlt/controllers:TrxJsonToBinController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"POST"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zlt/controllers:UserController"] = append(beego.GlobalControllerRouter["zlt/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zlt/controllers:UserController"] = append(beego.GlobalControllerRouter["zlt/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zlt/controllers:UserController"] = append(beego.GlobalControllerRouter["zlt/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:uid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zlt/controllers:UserController"] = append(beego.GlobalControllerRouter["zlt/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:uid",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zlt/controllers:UserController"] = append(beego.GlobalControllerRouter["zlt/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:uid",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zlt/controllers:UserController"] = append(beego.GlobalControllerRouter["zlt/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zlt/controllers:UserController"] = append(beego.GlobalControllerRouter["zlt/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: "/logout",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
