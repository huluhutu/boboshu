package test

import (
	"boboshu/config"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	_ "gorm.io/gorm/clause"
	"log"
)

type User struct {
	Name string
	Age  int
}

func InsetOne(ctx *gin.Context) {
	db := config.ConnectToMongoDB()
	userCollection := db.Database("test").Collection("users")
	user1 := &User{
		Name: "Lance",
		Age:  23,
	}
	defer config.DisConnectMongo(db)
	_, err := userCollection.InsertOne(context.TODO(), user1)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(200, user1)
}

func InsertMany(ctx *gin.Context) {
	db := config.ConnectToMongoDB()
	defer config.DisConnectMongo(db)
	userCollection := db.Database("test").Collection("users")
	user2 := []interface{}{User{Name: "HULUHUTU", Age: 23}, User{Name: "GaoFei", Age: 21}}
	_, err := userCollection.InsertMany(context.TODO(), user2)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(200, user2)
}

func UpdateOne(ctx *gin.Context) {
	db := config.ConnectToMongoDB()
	defer config.DisConnectMongo(db)
	userCollection := db.Database("test").Collection("users")
	filter := bson.D{{"name", "Lance"}}
	update := bson.D{{"$set", bson.D{{"name", "HULUHUTU"}}}}
	_, err := userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(200, "更新成功")
}

func FindOne(ctx *gin.Context) {
	db := config.ConnectToMongoDB()
	defer config.DisConnectMongo(db)
	userCollection := db.Database("test").Collection("users")
	filter := bson.D{{"name", "Lance"}}
	var user3 User
	err := userCollection.FindOne(context.TODO(), filter).Decode(&user3)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(200, user3)
}

func DeleteOne(ctx *gin.Context) {
	db := config.ConnectToMongoDB()
	defer config.DisConnectMongo(db)
	userCollection := db.Database("test").Collection("users")
	filter := bson.D{{"name", "Dean"}}
	one, err := userCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(200, one)
}
