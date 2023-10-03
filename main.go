package main

import (
	"apiTest.com/controllers"
	"apiTest.com/initializers"
	"apiTest.com/repositories"
	"apiTest.com/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

var (
	server *gin.Engine
)

func main() {
	//http.HandleFunc("/", hello) //設定存取的路由
	//log.Fatal(http.ListenAndServe(":5963", nil)) //設定監聽的埠

	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	testRepository := repositories.NewTestRepository(initializers.DB)
	testService := services.NewTestService(testRepository)
	testController := controllers.NewTestController(testService)

	router := server.Group("/api")

	router.GET("healthchecker", func(context *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres Example"
		context.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	testGroup := router.Group("/tests")
	{
		//testGroup.POST("", controllers.CreateTest)
		//testGroup.GET("", controllers.FindTests)
		//testGroup.GET("/:id", controllers.FindTests)
		//testGroup.PATCH("/:id", controllers.UpdateTest)
		//testGroup.DELETE("/:id", controllers.DeleteTest)

		testGroup.POST("", testController.CreateTest)
		testGroup.GET("", testController.FindTests)
		testGroup.GET("/:id", testController.FindOneTest)
		testGroup.PATCH("/:id", testController.UpdateTest)

	}

	log.Fatal(server.Run(":" + config.ServerPort))
}

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
	server = gin.Default()
}

func hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析參數，預設是不會解析的
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello! golang docker test") //這個寫入到 w 的是輸出到客戶端的
}
