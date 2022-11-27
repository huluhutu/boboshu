package router

import "boboshu/Test"

func init() {
	testGroup := Engine.Group("/test")
	{
		testGroup.GET("/mongo/insertOne", Test.InsetOne)
		testGroup.GET("/mongo/insertMany", Test.InsertMany)

	}
}
