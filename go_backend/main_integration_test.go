// !integration

package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/magiconair/properties/assert"
)

var API_URL = "http://localhost:8888"
var u_hello = API_URL + "/hello"
var u_addItem = API_URL + "/addItem"
var u_getItems = API_URL + "/inventoryList"
var u_updateItem = API_URL + "/updateItem"
var u_dbDeleteItem = API_URL + "/deleteItem"

func TestITSayHello(t *testing.T) {
	res, err := http.Get(u_hello)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	bodyByte, _ := io.ReadAll(res.Body)
	actualResult := string(bodyByte)
	expResult := `{"message":"hello2"}`
	assert.Equal(t, expResult, actualResult)
	assert.Equal(t, 200, res.StatusCode)
}

func TestAddItem(t *testing.T) {
	ui := inventory{
		Id:    22,
		Name:  "testitem",
		Price: 33,
		Sales: 333,
		Stock: 33,
	}
	jsonData, _ := json.Marshal(ui)
	contentType := "application/json"

	res, err := http.Post(u_addItem, contentType, bytes.NewBuffer(jsonData))
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()
	resBodyByte, _ := io.ReadAll(res.Body)
	respData := string(resBodyByte)
	// expData := `{"error":"item testitem already exists"}`
	expData := `"item testitem added"`

	assert.Equal(t, respData, expData)
	assert.Equal(t, res.StatusCode, 200)

}

func TestDbItem(t *testing.T) {
	itemName := "testitem"
	itemUrl := u_getItems + "/" + itemName
	res, err := http.Get(itemUrl)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	respBody, _ := io.ReadAll(res.Body)
	actual_resp := string(respBody)

	db := openDbConn()
	QUERY := "SELECT * FROM inventory WHERE name = '" + itemName + "';"
	rows, err := db.Query(QUERY)
	checkTestErr(err, t)
	var i_id int
	var i_name string
	var i_price float32
	var i_sales int
	var i_stock int

	for rows.Next() {
		err := rows.Scan(&i_id, &i_name, &i_price, &i_sales, &i_stock)
		if err != nil {
			checkTestErr(err, t)
		}
	}
	i_data := inventory{
		Id:    i_id,
		Name:  i_name,
		Price: i_price,
		Sales: i_sales,
		Stock: i_stock,
	}
	json_data, _ := json.Marshal(i_data)
	exp_resp := string(json_data)

	assert.Equal(t, actual_resp, exp_resp)
	assert.Equal(t, res.StatusCode, 200)
}

func TestDbUpdateItem(t *testing.T) {
	// Get Item Details from Db
	itemName := "testitem"

	db := openDbConn()
	QUERY := "SELECT * FROM inventory WHERE name = '" + itemName + "';"
	rows, err := db.Query(QUERY)
	checkTestErr(err, t)
	var i_id int
	var i_name string
	var i_price float32
	var i_sales int
	var i_stock int

	for rows.Next() {
		err := rows.Scan(&i_id, &i_name, &i_price, &i_sales, &i_stock)
		if err != nil {
			checkTestErr(err, t)
		}
	}
	i_data := inventory{
		Id:    i_id,
		Name:  i_name,
		Price: i_price,
		Sales: i_sales,
		Stock: i_stock,
	}
	json_data, _ := json.Marshal(i_data)

	req, err := http.NewRequest(http.MethodPut, u_updateItem, bytes.NewBuffer(json_data))
	checkTestErr(err, t)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	checkTestErr(err, t)
	defer res.Body.Close()
	exp_resp := `{"success":"item testitem updated successfully"}`

	res_byte, _ := io.ReadAll(res.Body)
	actual_resp := string(res_byte)

	assert.Equal(t, actual_resp, exp_resp)
	assert.Equal(t, res.StatusCode, 200)

}

func TestDbDeleteItem(t *testing.T) {
	itemName := "testitem"

	db := openDbConn()
	QUERY := "SELECT * FROM inventory WHERE name = '" + itemName + "';"
	rows, err := db.Query(QUERY)
	checkTestErr(err, t)
	var i_id int
	var i_name string
	var i_price float32
	var i_sales int
	var i_stock int

	for rows.Next() {
		err := rows.Scan(&i_id, &i_name, &i_price, &i_sales, &i_stock)
		if err != nil {
			checkTestErr(err, t)
		}
	}
	i_data := inventory{
		Id:    i_id,
		Name:  i_name,
		Price: i_price,
		Sales: i_sales,
		Stock: i_stock,
	}
	json_data, _ := json.Marshal(i_data)

	req, err := http.NewRequest(http.MethodDelete, u_dbDeleteItem, bytes.NewBuffer(json_data))

	if err != nil {
		t.Error(err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	checkTestErr(err, t)

	defer res.Body.Close()
	respBody, _ := io.ReadAll(res.Body)
	actual_resp := string(respBody)

	exp_resp := `{"success":"item deleted"}`
	assert.Equal(t, actual_resp, exp_resp)
	assert.Equal(t, res.StatusCode, 200)
}
func checkTestErr(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}
