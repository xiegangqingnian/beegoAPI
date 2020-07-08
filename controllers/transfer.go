package controllers

import (
	"encoding/json"
)

type SendTransaction struct {
	Compression string      `json:"compression"`
	Transaction Transaction `json:"transaction"`
	Signatures  []string    `json:"signatures"`
}

type Transaction struct {
	Expiration            string   `json:"expiration"`
	RefBlockNum           int      `json:"ref_block_num"`
	RefBlockPrefix        int      `json:"ref_block_prefix"`
	ContextFreeActions    []string `json:"context_free_actions"`
	Actions               []args_2 `json:"actions"`
	TransactionExtensions []string `json:"transaction_extensions"`
}

type transfer struct {
	Code   string `json:"code"`
	Action string `json:"action"`
	Args   struct {
		From     string `json:"from"`
		To       string `json:"to"`
		Quantity string `json:"quantity"`
		Memo     string `json:"memo"`
	} `json:"args"`
}

type args_0 struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Quantity string `json:"quantity"`
	Memo     string `json:"memo"`
}

type args_1 struct {
	From           string `json:"from"`
	To             string `json:"to"`
	Quantity       string `json:"quantity"`
	Memo           string `json:"memo"`
	RefBlockNum    int    `json:"ref_block_num"`
	RefBlockPrefix int    `json:"ref_block_prefix"`
	Expiration     string `json:"expiration"`
	Data           string `json:"data"`
	Signatures     string `json:"signatures"`
}

type args_2 struct {
	Account       string          `json:"account"`
	Name          string          `json:"name"`
	Authorization []Authorization `json:"authorization"`
	Data          string          `json:"data"`
}

type binargs struct {
	Binargs string `json:"binargs"`
}

//********************************

type PushRespon struct {
	TransactionID string    `json:"transaction_id"`
	Processed     Processed `json:"processed"`
}
type receipt struct {
	Status string `json:"status"`
	Cs     int    `json:"cs"`
	Ns     int    `json:"ns"`
}

type InlineTraces struct {
	Receipt         Receipt       `json:"receipt"`
	Act             Act           `json:"act"`
	ContextFree     bool          `json:"context_free"`
	Elapsed         int           `json:"elapsed"`
	Console         string        `json:"console"`
	TrxID           string        `json:"trx_id"`
	BlockNum        int           `json:"block_num"`
	BlockTime       string        `json:"block_time"`
	ProducerBlockID interface{}   `json:"producer_block_id"`
	Deltas          []interface{} `json:"deltas"`
	Except          interface{}   `json:"except"`
	InlineTraces    []interface{} `json:"inline_traces"`
}
type ActionTraces struct {
	Receipt         Receipt        `json:"receipt"`
	Act             Act            `json:"act"`
	ContextFree     bool           `json:"context_free"`
	Elapsed         int            `json:"elapsed"`
	Console         string         `json:"console"`
	TrxID           string         `json:"trx_id"`
	BlockNum        int            `json:"block_num"`
	BlockTime       string         `json:"block_time"`
	ProducerBlockID interface{}    `json:"producer_block_id"`
	Deltas          []interface{}  `json:"deltas"`
	Except          interface{}    `json:"except"`
	InlineTraces    []InlineTraces `json:"inline_traces"`
}
type Processed struct {
	ID              string         `json:"id"`
	BlockNum        int            `json:"block_num"`
	BlockTime       string         `json:"block_time"`
	ProducerBlockID interface{}    `json:"producer_block_id"`
	Receipt         receipt        `json:"receipt"`
	Elapsed         int            `json:"elapsed"`
	TS              int            `json:"t_s"`
	Scheduled       bool           `json:"scheduled"`
	ActionTraces    []ActionTraces `json:"action_traces"`
	Except          interface{}    `json:"except"`
}

func trxjsontobin(data transfer) interface{} {

	params, _ := json.Marshal(data)
	body := HttpPost(string(params), "chain", "abi_json_to_bin")
	var res binargs
	json.Unmarshal(body, &res)

	if res.Binargs == "" {
		err := JSONStruct{500, "try it again"}
		return err
	} else {
		return res
	}
}
func _trxjsontobin(data transfer) string {

	params, _ := json.Marshal(data)
	body := HttpPost(string(params), "chain", "abi_json_to_bin")
	var res binargs
	json.Unmarshal(body, &res)
	return res.Binargs
}

func sendtrx(params args_1) interface{} {

	actionData := []args_2{{"zltio.token", "transfer", []Authorization{{params.From, "active"}},
		params.Data}}
	tranData := Transaction{params.Expiration, params.RefBlockNum, params.RefBlockPrefix,
		[]string{}, actionData, []string{}}
	sendData := SendTransaction{"none", tranData, []string{params.Signatures}}
	sendDataJson, err := json.Marshal(sendData)
	if err != nil {
		errRes := JSONStruct{500, "toJson abnormal"}
		return errRes
	}
	body := HttpPost(string(sendDataJson), "chain", "push_transaction")
	//fmt.Println(string(body))
	//fmt.Println("**************555*****************")

	var result PushRespon
	var resGetTrx resGetTrx
	var error DuplicateTrxErr

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
			res := JSONStruct{500, "other abnormal"}
			return res
		}
	} else {
		resGetTrx.TrxID = result.TransactionID
		resGetTrx.Time = result.Processed.BlockTime
		resGetTrx.BlockHeight = result.Processed.BlockNum
		resGetTrx.Data = result.Processed.ActionTraces[0].Act.Data

		return resGetTrx
	}
	return nil
}
