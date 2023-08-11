package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// connect to the database
	// db = ConnectD() // used init to initialize the database connection
	// defer CloseDB(db)
	// removed the db parameter and accessed it globally
	// since it will be initialized by init before CloseDB or any of the
	// routes will be hit else we call FLog=Log+os.Exit, Log=log.Println
	defer CloseDB()

	// handlers
	router := gin.Default()
	router.POST("/", InsertHandler)
	router.GET("/", SelectAllHandler)
	router.GET("/:id", SelectOneByIdHandler)
	router.PUT("/:id", UpdateByIdHandler)
	router.DELETE("/:id", DeleteByIdHandler)

	// listen on port 8080
	router.Run("localhost:8080")

}
