package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"zlt/models"
)

var chain_url = "http://121.196.15.20:8888"
var wallet_url= "http://121.196.15.20:8900"
type GetInfoController struct {
	beego.Controller
}

type GetBlockController struct {
	beego.Controller
}

type GetBlocksController struct {
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

type NewAddrController struct {
	beego.Controller
}

type TrxHeightController struct {
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
	NumberOrHash string `json:"number_or_hash"`
}

type blocksParams struct {
	Start int `json:"start"`
	End   int `json:"end"`
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

// @Title GetTrxNumber
// @Description search
// @Param	body		body 		true  " body for user"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [GET]
func (t *TrxHeightController) Get() {
	res:=models.Search()
	t.Data["json"]=res
	t.ServeJSON()
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
	fmt.Println(blockNumber.NumberOrHash)
	if err != nil {
		error := JSONStruct{201, "Abnormal data format"}
		blo.Data["json"] = error
		blo.ServeJSON()
	}
	bloInfo := getBlockInfo(blockNumber.NumberOrHash)

	blo.Data["json"] = bloInfo
	blo.Ctx.ResponseWriter.WriteHeader(200)
	blo.ServeJSON()
}

// @Title GetBlocks
// @Description search blocks
// @Param	body	body 	blockNumber	 	"body for blockNumber"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [POST]
func (blo *GetBlocksController) Post() {
	var blocksNumber blocksParams
	data := blo.Ctx.Input.RequestBody
	json.Unmarshal(data, &blocksNumber)
	fmt.Println("11111111111111111111111111111111")
	fmt.Println(blocksNumber.Start)
	fmt.Println(blocksNumber.End)
	if(blocksNumber.Start>blocksNumber.End){
		err := JSONStruct{500,"Abnormal data format"}
		blo.Data["json"] = err
		blo.ServeJSON()
	}
	bloInfo := getBlocksInfo(blocksNumber)
	blo.Data["json"] = bloInfo
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
	addr.Data["json"] = getAddr(account.Addr)
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
// @Title NewAddrJsonToBin
// @Description create users
// @Param	body	body 	addr	true		"body for new addr"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [POST]
func (n *NewAddrController) Post() {
	var params NewAddrParams
	data := n.Ctx.Input.RequestBody
	json.Unmarshal(data, &params)
	n.Data["json"] = signNewAddr(params)
	n.ServeJSON()
}


// @Title SendTrx
// @Description send trx
// @Param	body	body 	addr	true	"body for data"
// @Success 200 {int}
// @Failure 403 body is empty
// @router / [POST]
func (trx *SendTrxController) Post() {
	var params ArgsNewTrx
	data := trx.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &params)
	if err != nil {
		error := JSONStruct{500, "Abnormal data format"}
		trx.Data["json"] = error
		trx.ServeJSON()
	}

	if checkAddrExist(params.From) && checkAddrExist(params.To) {
		res := sendtrx(params)
		trx.Data["json"] = res
		trx.ServeJSON()
		return
	} else {
		error := JSONStruct{500, "addr error"}
		trx.Data["json"] = error
		trx.ServeJSON()
	}
}
