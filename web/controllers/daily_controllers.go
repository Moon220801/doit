package controllers

import (
	"context"
	"doit/db"
	"doit/models"
	"log"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var daily []models.Daily

func GetDaily(ctx iris.Context) {
	var results []primitive.M

	d, err := db.Cdaily.Find(context.Background(), bson.D{{}})

	if err != nil {

	}
	for d.Next(context.Background()) {
		var result bson.M
		e := d.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		// fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)

	}

	if err := d.Err(); err != nil {
		log.Fatal(err)
	}
	ctx.JSON(results)

}

func CreateDaily(ctx iris.Context) {
	daily := new(models.Daily)
	err := ctx.ReadJSON(daily)

	insertResult, err := db.Cdaily.InsertOne(context.Background(), daily)
	if err != nil {
		panic(err.Error())
	}
	ctx.JSON(insertResult)

}

func DeleteOneDaily(ctx iris.Context) {

	id := ctx.Params().Get("id")
	objectId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objectId}
	deleteResult, err := db.Cdaily.DeleteOne(context.Background(), filter)
	if err != nil {

	}
	ctx.WriteString("掰掰拉( ˚ཫ˚ )\n")
	ctx.JSON(deleteResult)

}
