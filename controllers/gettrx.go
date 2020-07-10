package controllers

import (
	"encoding/json"
)

type trxParams struct {
	ID string `json:"id"`
}
type TrxId struct {
	ID string `json:"id"`
}
type TrxResult struct {
	Act             Act           `json:"act"`
	BlockNum        int           `json:"block_num"`
	BlockTime       string        `json:"block_time"`
	Console         string        `json:"console"`
	ContextFree     bool          `json:"context_free"`
	Deltas          []interface{} `json:"deltas"`
	Elapsed         int           `json:"elapsed"`
	Except          interface{}   `json:"except"`
	InlineTraces    []interface{} `json:"inline_traces"`
	ProducerBlockID interface{}   `json:"producer_block_id"`
	Receipt         Receipt       `json:"receipt"`
	TrxID           string        `json:"trx_id"`
}
type Authorization struct {
	Actor      string `json:"actor"`
	Permission string `json:"permission"`
}
type Data struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Memo     string `json:"memo"`
	Quantity string `json:"quantity"`
}
type Act struct {
	Account       string          `json:"account"`
	Name          string          `json:"name"`
	Authorization []Authorization `json:"authorization"`
	Data          Data            `json:"data"`
	HexData       string          `json:"hex_data"`
}
type Receipt struct {
	Receiver       string         `json:"receiver"`
	ActDigest      string         `json:"act_digest"`
	GlobalSequence int            `json:"global_sequence"`
	RecvSequence   int            `json:"recv_sequence"`
	AuthSequence   []AuthSequence `json:"auth_sequence"`
	CodeSequence   int            `json:"code_sequence"`
	AbiSequence    int            `json:"abi_sequence"`
}
type AuthSequence struct {
}

type resGetTrx struct {
	TrxID       string `json:"trx_id"`
	BlockHeight int    `json:"block_height"`
	Time        string `json:"time"`
	Data        Data   `json:"data"`
}

type resGettrxaddr struct {
	TrxID         string 		`json:"trx_id"`
	BlockHeight    int    		`json:"block_height"`
	Time           string 		`json:"time"`
	Addr       	   string	    `json:"addr"`
	CurrentBalance string       `json:"current_balance"`
}

func getTrx(trxId string) interface{} {

	var res interface{}
	var res2 interface{}
	var result TrxResult
	var resGetTrx resGetTrx

	//var resGettrx resGetTrx
	TrxID := TrxId{
		trxId,
	}
	params, _ := json.Marshal(TrxID)
	body := HttpPost(string(params), "history", "get_transaction")
	json.Unmarshal(body, &res)

	data, ok := res.(map[string]interface{})
	if ok {
		for _, v := range data {
			switch v2 := v.(type) {
			case []interface{}:
				for i, iv := range v2 {
					if i == 1 {
						res2 = iv
						//fmt.Println(i, iv)
					}
				}
			default:
				//fmt.Println(k,"is another type not handle yet")
			}
		}
	}
	temp, _ := json.Marshal(res2)
	json.Unmarshal(temp, &result)

	if result.TrxID == "" {
		err := &JSONStruct{201, "Invalid transaction ID"}
		return err
	} else {
		resGetTrx.TrxID = result.TrxID
		resGetTrx.BlockHeight = result.BlockNum
		resGetTrx.Time = result.BlockTime
		resGetTrx.Data = result.Act.Data
	}
	//fmt.Println(result.Act.Data.From)
	return resGetTrx
}
