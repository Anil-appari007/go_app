package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	DB_HOST     = os.Getenv("DB_HOST")
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_PORT_STR = os.Getenv("DB_PORT")
	DB_NAME     = os.Getenv("DB_NAME")
	// DB_PORT     int
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

func dbList() []inventory {
	db := openDbConn()
	rows, err := db.Query("SELECT * FROM inventory;")
	checkError(err, "dbList select")
	var dbList []inventory
	for rows.Next() {
		var item inventory
		err = rows.Scan(&item.Id, &item.Name, &item.Price, &item.Sales, &item.Stock)
		checkError(err, "dbList rows scan")
		dbList = append(dbList, item)
	}
	defer db.Close()
	return dbList
}

func dbinventoryList(c *gin.Context) {
	dbinventoryList := dbList()
	c.IndentedJSON(200, dbinventoryList)
}

func dbItem(c *gin.Context) {
	var getItem inventory
	ri := c.Param("name")
	QUUERY := "SELECT * FROM inventory WHERE name = '" + ri + "';"

	db := openDbConn()
	rows, err := db.Query(QUUERY)
	checkError(err, "dbItem query")
	if rows.Next() {
		err = rows.Scan(&getItem.Id, &getItem.Name, &getItem.Price, &getItem.Sales, &getItem.Stock)
		checkError(err, "dbItem row scan")
		c.IndentedJSON(200, getItem)
	} else {
		errorMessage := "item " + ri + " not found"
		c.IndentedJSON(404, gin.H{"error": errorMessage})
	}

}

func dbAddItem(c *gin.Context) {
	// Read the input, store in var
	// check db if item name exists
	// if not exists, add to db

	var ai inventory
	err := c.BindJSON(&ai)
	checkError(err, "dbAddItem bindjson")
	fmt.Printf("\nItem to add %s", ai.Name)

	QUERY := "SELECT * FROM inventory WHERE name = '" + ai.Name + "';"
	db := openDbConn()
	rows, err := db.Query(QUERY)
	checkError(err, "dbAddItem Query")
	if rows.Next() {
		message := "item " + ai.Name + " already exists"
		c.IndentedJSON(400, gin.H{"error": message})
	} else {
		/*
					INSERT INTO inventory (name, price, sales, stock)
			VALUES ('papayya', 33, 45, 90);
		*/
		AddQuery := "INSERT INTO inventory (name, price, sales, stock) VALUES ($1, $2, $3, $4);"
		// _, err = db.Exec(AddQuery)
		_, err = db.Exec(AddQuery, ai.Name, ai.Price, ai.Sales, ai.Stock)
		checkError(err, "dbAddItem exec query")
		message := "item " + ai.Name + " added"
		c.IndentedJSON(200, message)
	}
}

func dbv2(c *gin.Context) {
	db := openDbConn()
	rows, err := db.Query("SELECT version();")
	checkError(err, "dbv2 select version")
	var dbV2 string
	for rows.Next() {
		err = rows.Scan(&dbV2)
		checkError(err, "dbv2 row scan")
	}
	fmt.Println(dbV2)
	defer db.Close()
	c.IndentedJSON(200, gin.H{"dbv2": dbV2})
}

func openDbConn() *sql.DB {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT_STR, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)
	checkError(err, "SQL OPEN")
	err = db.Ping()
	checkError(err, "sql ping")
	return db
}

func checkError(err error, method string) {
	if err != nil {
		fmt.Printf("\nError During %s", method)
		fmt.Println(err)
		os.Exit(2)
	}
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
		// verify whether item details are empty or not
		if newItem.Id == 0 || newItem.Name == "" {
			// fmt.Println("invalid item details")
			c.IndentedJSON(400, gin.H{"error": "invalid item details"})
		} else {
			inventoryList = append(inventoryList, newItem)
			c.IndentedJSON(http.StatusCreated, newItem)
		}
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
	// fmt.Println(itemToDelete)

	var message string
	var status string
	var statusCode int
	var updatedInv []inventory
	var deleteIndex int

	// fmt.Printf("deleteIndex %d", deleteIndex)
	for n, each := range inventoryList {
		// fmt.Println("checking with", each, n)
		// if itemToDelete.Name == each.Name {
		// 	fmt.Println("Item ", each.Name, " matched with ", itemToDelete.Name)
		// 	status = "success"
		// 	message = "item " + itemToDelete.Name + " deleted"
		// 	statusCode = 200
		// 	deleteIndex = n
		// } else {
		// 	status = "error"
		// 	statusCode = 400
		// 	message = "item " + itemToDelete.Name + " to delete does not exist"
		// }
		// updatedInv = append(updatedInv, each)
		if itemToDelete.Name == each.Name {
			deleteIndex = n + 1
		}
		if deleteIndex == 0 {
			status = "error"
			statusCode = 400
			message = "item " + itemToDelete.Name + " to delete does not exist"
		} else {
			status = "success"
			message = "item " + itemToDelete.Name + " deleted"
			statusCode = 200
		}
	}

	if statusCode == 200 {
		// inventoryList = updatedInv
		for n, each := range inventoryList {
			if n != (deleteIndex - 1) {
				updatedInv = append(updatedInv, each)
			}
		}
		inventoryList = updatedInv
	}
	c.IndentedJSON(statusCode, gin.H{status: message})

}

func main() {

	if DB_HOST == "" || DB_USER == "" || DB_PASSWORD == "" || DB_PORT_STR == "" || DB_NAME == "" {
		fmt.Println("Db creds are invalid")
		os.Exit(2)
	}
	DB_PORT, err := strconv.Atoi(DB_PORT_STR)
	if err != nil {
		fmt.Println("Error During Atoi")
		fmt.Println(err)
	}
	fmt.Println(DB_PORT)
	// checkDbConnection(DB_HOST, DB_USER, DB_PASSWORD, DB_PORT, DB_NAME)
	// getDbVersion()

	fmt.Println("\nStarting GO Backend API Server")
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

	router.GET("dbv2", dbv2)
	router.GET("/dbinventoryList", dbinventoryList)
	router.GET("/dbinventoryList/:name", dbItem)
	router.POST("/dbAddItem", dbAddItem)
	// router.SetTrustedProxies(nil)
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Run("localhost:8888")
}
