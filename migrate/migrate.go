package main

import ("github.com/go-crud/initializers"
		"github.com/go-crud/models")

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}
	
func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}