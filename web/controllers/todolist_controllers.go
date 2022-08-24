package controllers

import (
	"context"
	"doit/db"
	"doit/models"
	"fmt"
	"log"
	"time"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var todolist []models.ToDoList

func GetTasks(ctx iris.Context) {
	results := []models.ToDoList{}
	cur, err := db.Ctodolist.Find(context.Background(), bson.D{{}})
	fmt.Println(*cur)
	if err != nil {
		log.Fatal(err)
	}
	var elem models.ToDoList

	for cur.Next(context.Background()) {
		e := cur.Decode(&elem)
		if e != nil {
			log.Fatal(e)
		}
		// fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, elem)

	}
	for i := range results {
		fmt.Println(results[i].Title)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	//ctx.JSON(results)

	/*data := iris.Map{
		"qq":      results,
		"Title":   result["title"],
		"Content": result["content"],
	}*/
	ctx.ViewData("Title", "hi")
	ctx.ViewData("qq", results)
	ctx.ViewData("author", "Iris expert")
	ctx.ViewData("chapterCount", "40")

	ctx.View("tasks")

	/*ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty"`
	Content     string             `json:"content,omitempty"`
	Type        string             `json:"type,omitempty"`
	Remark      string             `json:"remark,omitempty"`
	CreatedAt   time.Time          `json:"created_at,omitempty"`
	UpdateAt    time.Time          `json:"update_at,omitempty"`
	CompletedAt time.Time          `json:"completed_at,omitempty"`
	StartdateAt time.Time          `json:"startdate_at,omitempty"`
	EnddateAt   time.Time          `json:"enddate_at,omitempty"`
	Status      bool               `json:"status,omitempty"`*/

}

func CreateTasks(ctx iris.Context) {
	var (
		title   = ctx.FormValue("title")
		content = ctx.FormValue("content")
	)
	var tasks models.ToDoList

	//tasks := new(models.ToDoList)
	//err := ctx.ReadJSON(tasks)

	tasks.CreatedAt = time.Now()
	tasks.UpdateAt = time.Now()

	_, err := db.Ctodolist.InsertOne(context.Background(), bson.M{"title": title, "content": content})

	if err != nil {
		panic(err.Error())
	}

	//ctx.JSON(insertResult)
	//GetTasks(ctx)
	ctx.View("tasks")

}

func UpdoTasks(ctx iris.Context) {
	tasks := new(models.ToDoList)
	err := ctx.ReadJSON(tasks)

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

}

func DeleteOneTasks(ctx iris.Context) {
	method := ctx.FormValue("method")
	if method == "delete" {
		id := ctx.Params().Get("id")
		objectId, _ := primitive.ObjectIDFromHex(id)

		filter := bson.M{"_id": objectId}
		_, err := db.Ctodolist.DeleteOne(context.Background(), filter)
		if err != nil {
			fmt.Println(err.Error())
		}

	}
	GetTasks(ctx)

}

func DeleteAllTasks(ctx iris.Context) {

	d, err := db.Ctodolist.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		fmt.Println(err.Error())
	}
	ctx.JSON(d)

	ctx.WriteString("刪光光拉\n")

}

func SearchTasks(ctx iris.Context) {
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

}
