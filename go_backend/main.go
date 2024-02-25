package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type inventory struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Sales int     `json:"sales"`
	Stock int     `json:"stock"`
}

var inventoryList = []inventory{
	{Id: 1, Name: "apple", Price: 1, Sales: 11, Stock: 22},
	{Id: 2, Name: "banana", Price: 2.5, Sales: 244, Stock: 432},
	{Id: 3, Name: "mango", Price: 3.99, Sales: 342, Stock: 321},
	{Id: 4, Name: "orange", Price: 1.98, Sales: 34, Stock: 311},
}

func getList(c *gin.Context) {
	c.IndentedJSON(200, inventoryList)
}
func sayHello(c *gin.Context) {
	// c.JSON(200, gin.H{"message": "Hello"})
	c.IndentedJSON(200, gin.H{"message": "hello2"})
}

func addItem(c *gin.Context) {
	var newItem inventory
	err := c.BindJSON(&newItem)
	if err != nil {
		return
	}
	// don't add if element already exists with id
	check := 0
	for _, each := range inventoryList {
		if newItem.Id == each.Id {
			check = 1
			c.IndentedJSON(400, gin.H{"error": "id already exists"})
			fmt.Println("id already exists", newItem.Id)

			break
		} else {
			check = 0
		}
	}
	if check == 0 {
		inventoryList = append(inventoryList, newItem)
		c.IndentedJSON(http.StatusCreated, newItem)
	}

}
func getItemById(c *gin.Context) {
	for _, each := range inventoryList {
		if c.Param("name") == each.Name {
			c.IndentedJSON(200, each)
			return
		}
	}
}

func updateItem(c *gin.Context) {
	var updatedData inventory
	c.BindJSON(&updatedData)
	var message string
	var status string
	var statusCode int
	for index, eachItem := range inventoryList {
		if updatedData.Name == eachItem.Name {
			inventoryList[index] = updatedData
			status = "success"
			message = "item " + updatedData.Name + "updated to inventory"
			statusCode = 201
			break
		} else {
			status = "error"
			message = "item " + updatedData.Name + "not found"
			statusCode = 400
		}
	}
	c.IndentedJSON(statusCode, gin.H{status: message})

}

func deleteItem(c *gin.Context) {
	var itemToDelete inventory
	c.BindJSON(&itemToDelete)

	var message string
	var status string
	var statusCode int
	var updatedInv []inventory
	for _, each := range inventoryList {
		if itemToDelete.Name == each.Name {
			status = "success"
			message = "item " + itemToDelete.Name + " deleted"
			statusCode = 200
			continue
		} else {
			status = "error"
			statusCode = 400
			message = "item " + itemToDelete.Name + " to delete does not exist"
		}
		updatedInv = append(updatedInv, each)
	}
	if statusCode == 200 {
		inventoryList = updatedInv
	}
	c.IndentedJSON(statusCode, gin.H{status: message})

}

func main() {
	fmt.Println("Starting the backend")
	router := gin.Default()
	// router.Use(cors)
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.GET("/hello", sayHello)
	router.GET("/inventoryList", getList)
	router.GET("/inventoryList/:name", getItemById)

	router.POST("/addItem", addItem)

	router.PUT("/updateItem", updateItem)

	router.DELETE("/deleteItem", deleteItem)

	// router.SetTrustedProxies(nil)
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Run("localhost:8888")
}
