package controllers

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)


type NewAddrParams struct {
	Name   string `json:"name"`
	Key    string `json:"key"`
}

type NewAccount struct {
	Code   string 		`json:"code"`
	Action string 	    `json:"action"`
	Args   interface{}  `json:"args"`
}
type Args  struct {
	Creator string 		`json:"creator"`
	Name    string 		`json:"name"`
	Owner   interface{} `json:"owner"`
	Active  interface{} `json:"active"`
}
type Owner  struct {
	Threshold int 		   `json:"threshold"`
	Keys     []Key 		   `json:"keys"`
	Accounts []interface{} `json:"accounts"`
	Waits    []interface{} `json:"waits"`
}
type Key  struct {
	Key    string `json:"key"`
	Weight int    `json:"weight"`
}
type Active struct {
	Threshold int          `json:"threshold"`
	Keys     []Key		   `json:"keys"`
	Accounts []interface{} `json:"accounts"`
	Waits    []interface{} `json:"waits"`
}
//*******************************
type ArgsNewAddr struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	RefBlockNum    int    `json:"ref_block_num"`
	RefBlockPrefix int    `json:"ref_block_prefix"`
	Expiration     string `json:"expiration"`
	Data           string `json:"data"`
	Signatures     string `json:"signatures"`
}

type Sign struct {
	RefBlockNum       int           `json:"ref_block_num"`
	RefBlockPrefix    int           `json:"ref_block_prefix"`
	Expiration        string        `json:"expiration"`
	Actions           []args_2      `json:"actions"`
	Signatures        []string      `json:"signatures"`
}

type SignRespon struct {
	Expiration         string        `json:"expiration"`
	RefBlockNum        int           `json:"ref_block_num"`
	RefBlockPrefix     int           `json:"ref_block_prefix"`
	ContextFreeActions []interface{} `json:"context_free_actions"`
	Actions            []struct {
		Account       string `json:"account"`
		Name          string `json:"name"`
		Authorization []struct {
			Actor      string `json:"actor"`
			Permission string `json:"permission"`
		} `json:"authorization"`
		Data string `json:"data"`
	} `json:"actions"`
	TransactionExtensions []interface{} `json:"transaction_extensions"`
	Signatures            []string      `json:"signatures"`
	ContextFreeData       []interface{} `json:"context_free_data"`
}

func signNewAddr(params NewAddrParams) interface{} {

	var height resInfo
	var resblock  resBlockInfo
	var binData binargs
	var signRes SignRespon

	if(len(params.Key)==53 && !checkAddrExist(params.Name)) {
		owner := Owner{1, []Key{{params.Key, 1},}, []interface{}{}, []interface{}{}}
		active := Active{1, []Key{{params.Key, 1},}, []interface{}{}, []interface{}{}}
		data := NewAccount{"zltio", "newaccount", Args{"zltio", params.Name, owner, active}}
		jsonData, _ := json.Marshal(data)
	//fmt.Println(string(jsonData))
		body := HttpPost(string(jsonData), "chain", "abi_json_to_bin")
		json.Unmarshal(body, &binData)
		fmt.Println("json to bin****************")
	fmt.Println(binData.Binargs)
		//return res.Binargs
	}else{
		err := JSONStruct{500,"invaild Key or addr already  exist"}
		return err
	}
	//1.getinfo
	temp := getInfo()
	info,_:= json.Marshal(temp)
	json.Unmarshal(info,&height)

	//getblock
	temp2   := getBlockInfo(height.HeadBlockID)
	info2,_ := json.Marshal(temp2)
	json.Unmarshal(info2,&resblock)

	datatime := strings.Replace(resblock.HeadBlockTime,"T"," ",-1)

	//日期转换成时间戳
	timeLayout := "2006-01-02 15:04:05.000"  //模板
	loc, _:= time.LoadLocation("Local") //获取时区
	tmp, _:= time.ParseInLocation(timeLayout, datatime, loc)
	timestamp_0 := tmp.Unix() ////转化为时间戳 类型是int64
	//fmt.Println(timestamp_0)

	//时间戳加2分钟，后转回
	timestamp_1 :=timestamp_0+1200
	timestamp_2:= time.Unix(timestamp_1, 0).Format(timeLayout)
	timestamp:=strings.Replace(timestamp_2," ","T",-1)
	//fmt.Println(timestamp)

	//3.sign
	actions := []args_2{{"zltio","newaccount",[]Authorization{{"zltio","active"},},binData.Binargs}}
	context := Sign{resblock.HeadBlockNum,resblock.RefBlockPrefix,timestamp,actions,[]string{}}
	content_0 := []interface{}{context,[]string{"ZLT6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV"},
		"ed15a6a8a2bdb136cea6327a73cbc18304d7464a389238e338c9020e9a83f0d0"}
	content_1,_:=json.Marshal(content_0)
	body := HttpPost(string(content_1),"wallet","sign_transaction")
	json.Unmarshal(body,&signRes)

	//4.psuh

	tranData := Transaction{context.Expiration, context.RefBlockNum, context.RefBlockPrefix,
		[]string{}, actions, []string{}}
	sendData := SendTransaction{"none", tranData, signRes.Signatures}
	sendDataJson,_ := json.Marshal(sendData)
	fmt.Println("*************this post data************")
	fmt.Println(string(sendDataJson))
	body2 := HttpPost(string(sendDataJson),"chain","push_transaction")

	return pushNewAddr(params.Name,body2)
}

func pushNewAddr(name string,body []byte)interface{}{

	var result PushRespon
	var error DuplicateTrxErr
	var resGetTrx  resGettrxaddr
	var resGetAddr resGetAddr

	data := getAddr(name)
	temp,_ :=json.Marshal(data)
	json.Unmarshal(temp,&resGetAddr)

	json.Unmarshal(body, &result)
	if result.TransactionID == "" {
		json.Unmarshal(body, &error)
		switch error.Error.Code {
		case 3040005:
			res := JSONStruct{500, "expired_tx_exception"}
			return res
		case 3090003:
			res := JSONStruct{500, "authorization abnormal"}
			return res
		case 3010010:
			res := JSONStruct{500, "Invalid packed transaction，data abnormal"}
			return res
		case 3040007:
			res := JSONStruct{500, "Invalid packed transaction，ref_block  abnormal"}
			return res
		case 3040006:
			res := JSONStruct{500, "Invalid packed transaction，expired_tx too far abnormal"}
			return res
		default:
			fmt.Println(result)
			res := JSONStruct{500, "other abnormal"}
			return res
		}
	} else {

		resGetTrx.TrxID = result.TransactionID
		resGetTrx.Time = result.Processed.BlockTime
		resGetTrx.BlockHeight = result.Processed.BlockNum
		resGetTrx.Addr = resGetAddr.Addr
		resGetTrx.CurrentBalance = resGetAddr.CurrentBalance
		return resGetTrx
	}
}