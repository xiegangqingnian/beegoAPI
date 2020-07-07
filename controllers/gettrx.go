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
type authorization struct {
	Actor      string `json:"actor"`
	Permission string `json:"permission"`
}
type Data struct {
	From     string `json:"from"`
	Memo     string `json:"memo"`
	Quantity string `json:"quantity"`
	To       string `json:"to"`
}
type Act struct {
	Account       string          `json:"account"`
	Authorization []authorization `json:"authorization"`
	Data          Data            `json:"data"`
	HexData       string          `json:"hex_data"`
	Name          string          `json:"name"`
}
type Receipt struct {
	AbiSequence    int            `json:"abi_sequence"`
	ActDigest      string         `json:"act_digest"`
	AuthSequence   []AuthSequence `json:"auth_sequence"`
	CodeSequence   int            `json:"code_sequence"`
	GlobalSequence int            `json:"global_sequence"`
	Receiver       string         `json:"receiver"`
	RecvSequence   int            `json:"recv_sequence"`
}
type AuthSequence struct {
}

type resGetTrx struct {
	TrxID       string `json:"trx_id"`
	BlockHeight int    `json:"block_height"`
	Time        string `json:"time"`
	Data        Data   `json:"data"`
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
			case string:
				//fmt.Println(k,"is string",v2)
			case int:
			//	fmt.Println(k,"is int",v2)
			case bool:
			//	fmt.Println(k,"is bool",v2)
			case []interface{}:
				//fmt.Println(k,"is an array:")
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

	par, _ := json.Marshal(res2)
	json.Unmarshal(par, &result)

	if result.TrxID == "" {
		err := &JSONStruct{201, "No this address"}
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
