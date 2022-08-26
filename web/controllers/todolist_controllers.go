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

	t := time.Time(tasks.EnddateAt).Sub(tasks.CreatedAt)
	it64 := int64(t)
	tasks.Timeleft = strconv.FormatInt(it64, 10)

	fmt.Println(tasks.Timeleft)
	insertResult, err := db.Ctodolist.InsertOne(context.Background(), tasks)
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
		fmt.Println(insertResult)

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
		deleteResult, err := db.Ctodolist.DeleteOne(context.Background(), filter)
		if err != nil {
			fmt.Println(err.Error())
		}
		ctx.WriteString("掰掰拉( ˚ཫ˚ )\n")
		ctx.JSON(deleteResult)

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
	fmt.Println(id)
	d := db.Ctodolist.FindOne(context.Background(), filter)
	var result bson.M
	e := d.Decode(&result)

	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(e)

	if email == result["useremail"] {
		id := ctx.Params().Get("id")
		objectId, _ := primitive.ObjectIDFromHex(id)
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
		ctx.WriteString("小精靈查到你需要的任務了！٩(◦`꒳´◦)۶")

	} else {
		ctx.WriteString("你怎麼知道人家的任務꒰´•͈⌔•͈⑅꒱\n")

	}

}

func SearchTypeTasks(ctx iris.Context) {
	taskstype := ctx.Params().Get("type")
	email := middleware.AuthToken(ctx)

	d, _ := db.Ctodolist.Find(context.Background(), bson.M{"$and": []bson.M{{"type": taskstype}, {"useremail": email}}})
	fmt.Println(*d)

	var results []primitive.M

	for d.Next(context.Background()) {
		var result bson.M
		e := d.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		fmt.Println(e)
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

	if email == result["useremail"] {

		id := ctx.Params().Get("id")
		objectId, _ := primitive.ObjectIDFromHex(id)
		filter := bson.M{"_id": objectId}

		tasks.Timeleft = "0"

		tasks.CompletedAt = time.Now()
		update := bson.M{"$set": bson.M{"complete": true, "completedAt": tasks.CompletedAt, "status": "3"}}
		result, err := db.Ctodolist.UpdateOne(context.Background(), filter, update)
		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(result)
		ctx.WriteString("任務完成！你太棒了ʕ⸝⸝⸝˙Ⱉ˙ʔ")

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
	id := ctx.Params().Get("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	d := db.Ctodolist.FindOne(context.Background(), filter)
	var result bson.M
	e := d.Decode(&result)
	if e != nil {
		log.Fatal(e)
	}
	//a, _ := db.Ctodolist.Find(context.Background(), bson.M{"enddateAt": bson.M{"$gt": "createdAt"}})
	//b := a.Decode(&result)
	//fmt.Println(b)

	if email == result["useremail"] {

		id := ctx.Params().Get("id")
		objectId, _ := primitive.ObjectIDFromHex(id)
		filter := bson.M{"_id": objectId, "enddateat": bson.M{"$lt": time.Now()}}
		fmt.Println(filter)
		t := time.Time(tasks.EnddateAt).Sub(time.Now())
		it64 := int64(t)
		tasks.Timeleft = strconv.FormatInt(it64, 10)

		fmt.Println(tasks.Timeleft)

		update := bson.M{"$set": bson.M{"status": "1", "timeleft": tasks.Timeleft}}
		result, err := db.Ctodolist.UpdateOne(context.Background(), filter, update)
		fmt.Println(result)
		if err != nil {
			log.Fatal(err)
		}
		if result.MatchedCount != 0 {
			ctx.JSON(result)
			ctx.WriteString("任務逾期了拉！！")
		} else {
			filter2 := bson.M{"_id": objectId, "enddateat": bson.M{"$gt": time.Now()}, "startdateat": bson.M{"$lt": time.Now()}}

			update2 := bson.M{"$set": bson.M{"status": "2", "timeleft": tasks.Timeleft}}
			result2, err := db.Ctodolist.UpdateOne(context.Background(), filter2, update2)
			if err != nil {
				log.Fatal(err)
			}

			ctx.JSON(result2)
			ctx.WriteString("很棒，任務正在進行(*°ω°*ฅ)*")

		}

	} else {
		ctx.WriteString("你怎麼知道人家的任務(´･ᆺ･`)\n")

	}
	//result, err := db.Ctodolist.UpdateMany(context.Background(), filter, bson.D{{"$set", bson.D{{"size.uom", "cm"}, {"status", "P"}}}})

}
