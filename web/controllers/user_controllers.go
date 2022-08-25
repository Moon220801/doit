package controllers

import (
	"context"
	"doit/db"
	"doit/middleware"
	"doit/models"
	"fmt"
	"time"

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

	user := models.User{}

	err := ctx.ReadJSON(&user)

	if err != nil {
		fmt.Println(err.Error())
	}

	if user.Email == "" && user.Password == "" {
		ctx.WriteString("啥都沒有呢\n")
	} else if user.Password == "" {
		ctx.WriteString("請記得輸入你的密碼\n")
	} else if user.Email == "" {
		ctx.WriteString("請記得輸入你的email\n")

	}
	var result bson.M
	d := db.Cuser.FindOne(context.Background(), bson.M{"email": user.Email})

	e := d.Decode(&result)

	if result["password"] == user.Password {
		ctx.WriteString("登錄成功\n歡迎回來")
		ctx.JSON(result)
		middleware.GetTokenHandler(ctx, user.Email)

	} else {
		ctx.WriteString("登錄失敗\n")
	}

	if e != nil {

		ctx.WriteString("找不到您的email唷\n")
	}

}

func Regiter(ctx iris.Context) {

	users := models.Users{}

	err := ctx.ReadJSON(&users)

	if err != nil {
		fmt.Println(err.Error())
	}

	var result bson.M
	d := db.Cuser.FindOne(context.Background(), bson.M{"email": users.Email}).Decode(&result)
	users.Email_verified_at = time.Now()
	if d != nil {

		insertResult, err := db.Cuser.InsertOne(context.Background(), users)
		fmt.Println(insertResult)
		if err != nil {
			panic(err.Error())
		}
		ctx.WriteString("創建成功，歡迎加入！！\n")
		ctx.JSON(users)
	} else {
		ctx.WriteString("已有此帳戶\n")
	}
}

func Logout(ctx iris.Context) {
	email := middleware.AuthToken(ctx)
	middleware.DeleteToken(ctx, email)
	ctx.WriteString("拜拜 下次見  ʕ •̀ o •́ ʔ\n")

}
