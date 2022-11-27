package router

import "boboshu/Test"

func init() {
	testGroup := Engine.Group("/test")
	{
		testGroup.GET("/mongo/insertOne", Test.InsetOne)
		testGroup.GET("/mongo/insertMany", Test.InsertMany)
		testGroup.GET("/mongo/updateOne", Test.UpdateOne)
		testGroup.GET("/mongo/findOne", Test.FindOne)
		testGroup.GET("/mongo/deleteOne", Test.DeleteOne)
	}
}
