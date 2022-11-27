package Test

import (
	"boboshu/config"
	"context"
	"github.com/gin-gonic/gin"
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
