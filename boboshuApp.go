package main

import "boboshu/router"

func main() {
	engine := router.GetEngine()
	engine.Run(":8808").Error()
}
