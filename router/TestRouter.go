package router

import "boboshu/test"

func init() {
	testGroup := Engine.Group("/test")
	{
		testGroup.GET("/mongo/insertOne", test.InsetOne)
		testGroup.GET("/mongo/insertMany", test.InsertMany)
		testGroup.GET("/mongo/updateOne", test.UpdateOne)
		testGroup.GET("/mongo/findOne", test.FindOne)
		testGroup.GET("/mongo/deleteOne", test.DeleteOne)
	}
}
