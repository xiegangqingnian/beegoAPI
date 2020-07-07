package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

var chain_url = "http://47.92.208.227:8888"

type GetInfoController struct {
	beego.Controller
}

type GetBlockController struct {
	beego.Controller
}

type GetTrxController struct {
	beego.Controller
}

type GetAddrController struct {
	beego.Controller
}

type TrxJsonToBinController struct {
	beego.Controller
}
type SendTrxController struct {
	beego.Controller
}

type JSONStruct struct {
	Code int
	Msg  string
}

type blockParams struct {
	BlockNumOrId string `json:"block_num_or_id"`
}

// @Title GetInfo
// @Description create users
// @Param	body		body 		true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [GET]
func (u *GetInfoController) Get() {
	resInfo := getInfo()
	u.Data["json"] = resInfo
	u.ServeJSON()
}

// @Title GetBlock
// @Description create users
// @Param	body		body 	blockNumber	true		"body for blockNumber"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [POST]
func (blo *GetBlockController) Post() {
	var blockNumber blockParams
	data := blo.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &blockNumber)
	fmt.Println(blockNumber.BlockNumOrId)
	if err != nil {
		error := JSONStruct{201, "Abnormal data format"}
		blo.Data["json"] = error
		blo.ServeJSON()
	}
	bloInfo := getBlockInfo(blockNumber.BlockNumOrId)

	blo.Data["json"] = bloInfo
	blo.Ctx.ResponseWriter.WriteHeader(200)
	blo.ServeJSON()
}

// @Title GetTrx
// @Description create users
// @Param	body		body 	trxId	true		"body for trxId"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [POST]
func (trx *GetTrxController) Post() {
	var trxId trxParams
	data := trx.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &trxId)
	if err != nil {
		error := JSONStruct{201, "Abnormal data format"}
		trx.Data["json"] = error
		trx.ServeJSON()
	}
	//fmt.Println(trxId.ID)
	trx.Data["json"] = getTrx(trxId.ID)
	trx.Ctx.ResponseWriter.WriteHeader(200)
	trx.ServeJSON()
}

// @Title GetAddr
// @Description create users
// @Param	body		body 	addr	true		"body for trxId"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [POST]
func (addr *GetAddrController) Post() {

	var account Addr
	data := addr.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &account)
	if err != nil {
		res := JSONStruct{500, "Abnormal data format"}
		addr.Data["json"] = res
	}
	addr.Data["json"] = getAddr(account.AccountName)
	addr.ServeJSON()
}

// @Title JsonToBin
// @Description create users
// @Param	body		body 	addr	true		"body for trxId"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [POST]
func (addr *TrxJsonToBinController) Post() {
	var params transfer
	data := addr.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &params)
	if err != nil {
		err := JSONStruct{500, "try it again"}
		addr.Data["json"] = err
		addr.ServeJSON()
	}
	addr.Data["json"] = trxjsontobin(params)
	addr.ServeJSON()
}

// @Title SendTrx
// @Description send trx
// @Param	body	body 	addr	true	"body for data"
// @Success 200 {int}
// @Failure 403 body is empty
// @router / [POST]
func (trx *SendTrxController) Post() {
	var params args_1
	data := trx.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &params)
	if err != nil {
		error := JSONStruct{500, "Abnormal data format"}
		trx.Data["json"] = error
		trx.ServeJSON()
	}

	//fmt.Println(params.From,params.Signatures)

	if checkAddrExist(params.From) && checkAddrExist(params.To) {
		error := JSONStruct{500, "addr error"}
		trx.Data["json"] = error
		trx.ServeJSON()
	} else {

		sendtrx(params)
		//fmt.Println(res)
		//error := JSONStruct{500,"No this from or to address"}
		//trx.Data["json"] = error
		//trx.ServeJSON()
		//return
	}

	//addr.ServeJSON()
}
