#!/bin/bash

REACT_APP_API_URL="localhost"

curl -s ${REACT_APP_API_URL}:8888/hello 
curl -s ${REACT_APP_API_URL}:8888/inventoryList


curl -X POST ${REACT_APP_API_URL}:8888/addItem \
   -H 'Content-Type: application/json' \
   -d '{"id":5,"name":"guava","price":22,"sales":22,"stock":22}'

curl -s ${REACT_APP_API_URL}:8888/inventoryList


curl -X POST ${REACT_APP_API_URL}:8888/addItem \
   -H 'Content-Type: application/json' \
   -d '{"id":5,"name":"guava","price":22,"sales":22,"stock":22}'



curl -s ${REACT_APP_API_URL}:8888/inventoryList/mango


curl -X PUT ${REACT_APP_API_URL}:8888/updateItem \
   -H 'Content-Type: application/json' \
   -d '{"id":5,"name":"guava","price":333,"sales":22,"stock":22}'

curl -s ${REACT_APP_API_URL}:8888/inventoryList



curl -X DELETE ${REACT_APP_API_URL}:8888/deleteItem \
   -H 'Content-Type: application/json' \
   -d '{"id":5,"name":"guava","price":333,"sales":22,"stock":22}'

curl -s ${REACT_APP_API_URL}:8888/inventoryList

curl -X DELETE ${REACT_APP_API_URL}:8888/deleteItem \
   -H 'Content-Type: application/json' \
   -d '{"id":5,"name":"guava","price":333,"sales":22,"stock":22}'

curl -s ${REACT_APP_API_URL}:8888/inventoryList