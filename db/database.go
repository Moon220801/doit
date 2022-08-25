package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Cuser *mongo.Collection
var Ctodolist *mongo.Collection
var Cdaily *mongo.Collection

func init() {
	LoadTheEnv()
	CreateDBInstance()
}

func LoadTheEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
func CreateDBInstance() {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collUser := os.Getenv("DB_COLLECTION_USER")
	collToDoList := os.Getenv("DB_COLLECTION_TODOLIST")
	collDaily := os.Getenv("DB_COLLECTION_DAILY")
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	Cuser = client.Database(dbName).Collection(collUser)
	Ctodolist = client.Database(dbName).Collection(collToDoList)
	Cdaily = client.Database(dbName).Collection(collDaily)

	fmt.Println("Collection instance created!")
}

func NewClient() *redis.Client { // 實體化redis.Client 並返回實體的位址
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}
