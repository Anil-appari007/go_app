#!/bin/bash

curl -s localhost:8888/hello 
curl -s localhost:8888/inventoryList


curl -X POST localhost:8888/addItem \
   -H 'Content-Type: application/json' \
   -d '{"id":5,"name":"guava","price":22,"sales":22,"stock":22}'

curl -s localhost:8888/inventoryList


curl -X POST localhost:8888/addItem \
   -H 'Content-Type: application/json' \
   -d '{"id":5,"name":"guava","price":22,"sales":22,"stock":22}'



curl -s localhost:8888/inventoryList/mango


curl -X PUT localhost:8888/updateItem \
   -H 'Content-Type: application/json' \
   -d '{"id":5,"name":"guava","price":333,"sales":22,"stock":22}'

curl -s localhost:8888/inventoryList



curl -X DELETE localhost:8888/deleteItem \
   -H 'Content-Type: application/json' \
   -d '{"id":5,"name":"guava","price":333,"sales":22,"stock":22}'

curl -s localhost:8888/inventoryList

curl -X DELETE localhost:8888/deleteItem \
   -H 'Content-Type: application/json' \
   -d '{"id":5,"name":"guava","price":333,"sales":22,"stock":22}'

curl -s localhost:8888/inventoryList