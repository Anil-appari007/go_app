package main

import (
	"database/sql"
	"fmt"

	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	DB_HOST     = os.Getenv("DB_HOST")
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_PORT     = os.Getenv("DB_PORT")
	DB_NAME     = os.Getenv("DB_NAME")
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
	c.JSON(200, dbinventoryList)
}

func dbItem(c *gin.Context) {
	var getItem inventory
	ri := c.Param("name")
	QUUERY := "SELECT * FROM inventory WHERE name = '" + ri + "';"

	db := openDbConn()
	defer db.Close()
	rows, err := db.Query(QUUERY)
	checkError(err, "dbItem query")
	if rows.Next() {
		err = rows.Scan(&getItem.Id, &getItem.Name, &getItem.Price, &getItem.Sales, &getItem.Stock)
		checkError(err, "dbItem row scan")
		c.JSON(200, getItem)
	} else {
		errorMessage := "item " + ri + " not found"
		c.JSON(404, gin.H{"error": errorMessage})
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
	defer db.Close()
	rows, err := db.Query(QUERY)
	checkError(err, "dbAddItem Query")
	if rows.Next() {
		message := "item " + ai.Name + " already exists"
		c.JSON(400, gin.H{"error": message})
	} else {
		AddQuery := "INSERT INTO inventory (name, price, sales, stock) VALUES ($1, $2, $3, $4);"
		_, err = db.Exec(AddQuery, ai.Name, ai.Price, ai.Sales, ai.Stock)
		checkError(err, "dbAddItem exec query")
		message := "item " + ai.Name + " added"
		c.JSON(200, message)
	}
}
func dbUpdateItem(c *gin.Context) {
	var ui inventory
	err := c.BindJSON(&ui)
	// checkError(err, "dbUpdateItem bindjson")

	if err != nil {
		fmt.Printf("Invalid Item Details %+v\n", ui)
		message := "item " + ui.Name + " has invalid details"
		c.JSON(400, gin.H{"error": message})
		return
	}
	// check the item exists or not
	// update if exists
	db := openDbConn()
	SEARCH_Q := "SELECT name FROM inventory WHERE name = '" + ui.Name + "';"
	rows, err := db.Query(SEARCH_Q)
	checkError(err, "dbUpdateItem query")
	if rows.Next() {
		// update inventory SET price = 12, sales = 20, stock = 49 WHERE name = 'apple';
		U_Q := "update inventory SET price = $2, sales = $3, stock = $4 WHERE name = $1;"
		_, err = db.Exec(U_Q, ui.Name, ui.Price, ui.Sales, ui.Stock)
		checkError(err, "dbUpdateItem exec")
		message := "item " + ui.Name + " updated successfully"
		c.JSON(200, gin.H{"success": message})
	} else {
		message := "item " + ui.Name + " not found"
		c.JSON(400, gin.H{"error": message})
	}
}

func dbDeleteItem(c *gin.Context) {
	var di inventory
	err := c.BindJSON(&di)
	// checkError(err, "Error in BindJson")
	if err != nil {
		fmt.Printf("Invalid Json details %+v \n", di)
		c.JSON(400, gin.H{"error": "Invalid Json details"})
		return
	}

	db := openDbConn()
	defer db.Close()

	rows, err := db.Query("SELECT name FROM inventory WHERE name = $1;", di.Name)
	checkError(err, "dbDeleteItem select query")
	if rows.Next() {
		// Delete
		DELETE_Q := "DELETE FROM inventory WHERE name = $1;"
		_, err = db.Exec(DELETE_Q, di.Name)
		checkError(err, "dbDeleteItem delete query exec")
		c.JSON(200, gin.H{"success": "item deleted"})
	} else {
		// Item Not Found
		c.JSON(400, gin.H{"error": "item not found"})
	}
}

func openDbConn() *sql.DB {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)
	checkError(err, "SQL OPEN")
	err = db.Ping()
	checkError(err, "sql ping")
	return db
}

func checkError(err error, method string) {
	if err != nil {
		fmt.Printf("\nError During %s\n", method)
		fmt.Println(err)
		return
		// os.Exit(2)
	}
}

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello2"})
}

func main() {

	if DB_HOST == "" || DB_USER == "" || DB_PASSWORD == "" || DB_PORT == "" || DB_NAME == "" {
		fmt.Println("Db creds are invalid")
		os.Exit(2)
	}

	fmt.Println("\nStarting GO Backend API Server")
	router := gin.Default()
	// router.Use(cors)
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.GET("/hello", sayHello)

	router.GET("/inventoryList", dbinventoryList)
	router.GET("/inventoryList/:name", dbItem)

	router.POST("/addItem", dbAddItem)

	router.PUT("/updateItem", dbUpdateItem)

	router.DELETE("/deleteItem", dbDeleteItem)
	// router.SetTrustedProxies(nil)
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Run("localhost:8888")
}
