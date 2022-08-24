package controllers

import (
	"context"
	"doit/db"
	"doit/models"
	"fmt"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
)

type UserController struct {
}

var user []models.User

func handleErr(ctx iris.Context, err error) {
	ctx.JSON(iris.Map{"response": err.Error()})
}

func Login(ctx iris.Context) {

	//user := models.User{}
	var (
		email    = ctx.FormValue("email")
		password = ctx.FormValue("password")
	)
	//err := ctx.ReadJSON(&user)
	fmt.Println(email, password)
	//filter := bson.M{"email": user.Email}

	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	if email == "" {
		ctx.WriteString("請記得輸入你的email\n")
	} else if password == "" {
		ctx.WriteString("請記得輸入你的密碼\n")
	} else if email == "" && password == "" {
		ctx.WriteString("啥都沒有呢\n")
	}
	var result bson.M
	d := db.Cuser.FindOne(context.Background(), bson.M{"email": email})
	fmt.Println(d)
	e := d.Decode(&result)

	if result["password"] == password {
		//ctx.WriteString("登錄成功\n歡迎回來")
		//ctx.JSON(result)
		//m.getTokenHandler(ctx)
		data := iris.Map{
			"Title":      "Page Title",
			"FooterText": "Footer contents",
			"Message":    "Main contents",
		}

		ctx.View("tasks", data)
	} else {
		ctx.WriteString("登錄失敗\n")
	}

	if e != nil {

		ctx.WriteString("找不到您的email唷\n")
	}

}

func Regiter(ctx iris.Context) {

	//user := models.User{}
	var (
		name     = ctx.FormValue("name")
		email    = ctx.FormValue("email")
		password = ctx.FormValue("password")
	)
	var users models.Users
	err := ctx.ReadJSON(&users)
	if err != nil {
		fmt.Println(err.Error())
	}

	var result bson.M
	d := db.Cuser.FindOne(context.Background(), bson.M{"email": email})
	e := d.Decode(&result)

	if e != nil {
		insertResult, err := db.Cuser.InsertOne(context.Background(), bson.M{"name": name, "email": email, "password": password})

		fmt.Println(insertResult)

		if err != nil {
			panic(err.Error())
		}
		data := iris.Map{
			"Title":      "Page Title",
			"FooterText": "Footer contents",
			"Message":    "Main contents",
		}

		ctx.View("tasks", data)
	} else {
		data := iris.Map{
			"Title":      "出錯紀錄",
			"FooterText": "Footer contents",
			"Message":    "已有此帳戶",
		}
		ctx.View("error", data)
	}

	//users.Email_verified_at = time.Now()

}
