package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"

)
var chain_url="http://47.92.208.227:8888"

type GetInfoController struct {
	beego.Controller
}

type GetBlockController struct {
	beego.Controller
}

type GetTrxController struct {
	beego.Controller
}

// @Title GetInfo
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [GET]
func (u *GetInfoController) Get() {
	resInfo :=getInfo()
	u.Data["json"]= resInfo
	u.ServeJSON()
}

type blockParams struct {
	BlockNumOrId string `json:"block_num_or_id"`
}

type trxParams struct {
	ID string `json:"id"`
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
	err := json.Unmarshal(data,&blockNumber)
	fmt.Println(blockNumber.BlockNumOrId)
	if err != nil {
		blo.Ctx.WriteString("Parameter abnormality ")
		blo.Ctx.ResponseWriter.WriteHeader(500)
		return
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
	err  := json.Unmarshal(data,&trxId)
	if err != nil {
		trx.Ctx.WriteString("Parameter abnormality")
		trx.Ctx.ResponseWriter.WriteHeader(500)
		return
	}
	fmt.Println(trxId.ID)

	res :=getTrx(trxId.ID)
	fmt.Printf("************************\n")
	fmt.Printf("%T",res)


	trx.Data["json"] = res

	trx.Ctx.ResponseWriter.WriteHeader(200)
	trx.ServeJSON()
}