package controllers

import (
	"encoding/json"
	"fmt"
)

type TrxId struct {
	ID string `json:"id"`
}

func getTrx(trxId string) interface{}{

	var res interface{}
	var res2 []interface{}
	TrxID := TrxId{
		trxId,
		}
	params,_:=json.Marshal(TrxID)

	body :=HttpPost(string(params),"history","get_transaction")
	json.Unmarshal(body,&res)

	fmt.Printf("%T",res)

	data, ok := res.(map[string]interface{})
	if ok {
		for k,v := range data{
			switch v2 := v.(type) {
			case string :
				fmt.Println(k,"is string",v2)
			case int :
				fmt.Println(k,"is int",v2)
			case bool :
				fmt.Println(k,"is bool",v2)
			case []interface{}:
				fmt.Println(k,"is an array:")
				for i, iv := range v2{
					res2 = append(res2,v2)
					fmt.Println(i,iv)
				}
			default:
				fmt.Println(k,"is another type not handle yet")
			}
		}
	}


	fmt.Printf("*******222222222222222222222****************\n")
	return res
}