package controllers

import (
	"context"
	"doit/db"
	"doit/middleware"
	"doit/models"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var todolist []models.ToDoList

func GetTasks(ctx iris.Context) {
	email := middleware.AuthToken(ctx)
	filter := bson.D{{"useremail", bson.D{{"$eq", email}}}}
	fmt.Println(filter)

	var results []primitive.M

	d, err := db.Ctodolist.Find(context.Background(), filter)

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

func CreateTasks(ctx iris.Context) {
	email := middleware.AuthToken(ctx)
	tasks := new(models.ToDoList)
	err := ctx.ReadJSON(tasks)
	tasks.Useremail = email
	tasks.CreatedAt = time.Now()
	tasks.UpdateAt = time.Now()
	tasks.Status = "0"
	insertResult, err := db.Ctodolist.InsertOne(context.Background(), tasks)
	t := time.Time(tasks.EnddateAt).Sub(tasks.CreatedAt)
	it64 := int64(t)
	tasks.Timeleft = strconv.FormatInt(it64, 10)

	fmt.Println(tasks.Timeleft)

	if err != nil {
		panic(err.Error())
	}

	ctx.JSON(insertResult)

}

func UpdoTasks(ctx iris.Context) {

	tasks := new(models.ToDoList)
	err := ctx.ReadJSON(tasks)

	email := middleware.AuthToken(ctx)
	id := ctx.Params().Get("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	d := db.Ctodolist.FindOne(context.Background(), filter)
	var result bson.M
	e := d.Decode(&result)
	if e != nil {
		log.Fatal(e)
	}

	if email == result["useremail"] {
		id := ctx.Params().Get("id")
		objectId, _ := primitive.ObjectIDFromHex(id)
		filter := bson.M{"_id": objectId}
		update := bson.M{"$set": tasks}

		if err != nil {
			fmt.Println(err.Error())
		}
		tasks.UpdateAt = time.Now()

		insertResult, err := db.Ctodolist.UpdateOne(context.Background(), filter, update)

		if err != nil {
			panic(err.Error())
		}

		ctx.JSON(insertResult)

	} else {
		ctx.WriteString("不要亂偷人家任務拉！ಠ_ಠ\n")
	}

}

func DeleteOneTasks(ctx iris.Context) {
	tasks := new(models.ToDoList)
	err := ctx.ReadJSON(tasks)

	if err != nil {
		fmt.Println(err.Error())
	}

	email := middleware.AuthToken(ctx)
	id := ctx.Params().Get("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	d := db.Ctodolist.FindOne(context.Background(), filter)
	var result bson.M
	e := d.Decode(&result)
	if e != nil {
		log.Fatal(e)
	}
	if email == result["useremail"] {
		id := ctx.Params().Get("id")
		objectId, _ := primitive.ObjectIDFromHex(id)

		filter := bson.M{"_id": objectId}
		_, err := db.Ctodolist.DeleteOne(context.Background(), filter)
		if err != nil {
			fmt.Println(err.Error())
		}

	} else {
		ctx.WriteString("不要亂刪人家任務拉！ಠ_ಠ\n")
	}

}

func DeleteAllTasks(ctx iris.Context) {

	email := middleware.AuthToken(ctx)
	filter := bson.D{{"useremail", bson.D{{"$eq", email}}}}
	fmt.Println(filter)

	d, err := db.Ctodolist.DeleteMany(context.Background(), filter)
	if err != nil {
		fmt.Println(err.Error())
	}
	ctx.JSON(d)

	ctx.WriteString("刪光光拉\n")

}

func SearchTasks(ctx iris.Context) {
	id := ctx.Params().Get("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	email := middleware.AuthToken(ctx)

	filter := bson.M{"_id": objectId}
	d := db.Ctodolist.FindOne(context.Background(), filter)
	var result bson.M
	e := d.Decode(&result)

	if e != nil {
		log.Fatal(e)
	}

	if email == result["email"] {
		filter := bson.M{"_id": objectId}
		d := db.Ctodolist.FindOne(context.Background(), filter)

		var results []primitive.M
		var result bson.M
		e := d.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)

		ctx.JSON(results)
	}

}

func SearchTypeTasks(ctx iris.Context) {
	taskstype := ctx.Params().Get("type")
	email := middleware.AuthToken(ctx)

	d, _ := db.Ctodolist.Find(context.Background(), bson.M{"$and": []bson.M{{"type": taskstype}, {"useremail": email}}})
	fmt.Println(d)

	var results []primitive.M

	for d.Next(context.Background()) {
		var result bson.M
		e := d.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		// fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)

	}

	ctx.JSON(results)

}

func CompleteTask(ctx iris.Context) {
	tasks := new(models.ToDoList)
	err := ctx.ReadJSON(tasks)

	if err != nil {
		fmt.Println(err.Error())
	}

	email := middleware.AuthToken(ctx)
	id := ctx.Params().Get("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	d := db.Ctodolist.FindOne(context.Background(), filter)
	var result bson.M
	e := d.Decode(&result)
	if e != nil {
		log.Fatal(e)
	}

	if email == result["email"] {

		id := ctx.Params().Get("id")
		objectId, _ := primitive.ObjectIDFromHex(id)
		filter := bson.M{"_id": objectId}

		tasks.CompletedAt = time.Now()
		update := bson.M{"$set": bson.M{"complete": true, "completedAt": tasks.CompletedAt, "state": "3"}}
		result, err := db.Ctodolist.UpdateOne(context.Background(), filter, update)
		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(result)
		ctx.WriteString("任務完成！")

	} else {
		ctx.WriteString("你怎麼知道人家的任務(´･ᆺ･`)\n")

	}

}

func UpdoTasksDate(ctx iris.Context) {
	tasks := new(models.ToDoList)
	err := ctx.ReadJSON(tasks)

	if err != nil {
		fmt.Println(err.Error())
	}

	email := middleware.AuthToken(ctx)
	filter := bson.D{{"useremail", bson.D{{"$eq", email}}}}
	fmt.Println(filter)
	//result, err := db.Ctodolist.UpdateMany(context.Background(), filter, bson.D{{"$set", bson.D{{"size.uom", "cm"}, {"status", "P"}}}})

}
