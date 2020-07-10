package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type GetTrxNum struct {
	TotalTrxNumber int64 `json:"total_trx_number"`
}

func Search() interface{}{

	clientOptions := options.Client().ApplyURI("mongodb://47.114.165.27:27018")

	client,err :=mongo.Connect(context.TODO(),clientOptions)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("zltdb").Collection("transactions")
	fmt.Println("***************1***************!")
	count, err:= collection.CountDocuments(context.TODO(),bson.D{})
	if(err != nil){
		log.Fatal(err)
	}
	//fmt.Println("*************************************!")
	//fmt.Println("height!")
	fmt.Println(count)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Connection to MongoDB closed.")
	res :=GetTrxNum{count}
	return res
}