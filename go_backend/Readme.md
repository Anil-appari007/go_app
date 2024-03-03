Aim:

Create a crud api with MySQL/Postgres DB
- Get inventory list
- update inventory list
- delete item in inventory
- add an item to inventory

item data

id  name    price   sales   stock
int string  float   int     int
i   apple   1       11      22
2   banana  2.5     244     432
3   mango   3.99    342     321
4   orange  1.98    34      311



go test --tags=integration -coverprofile=coverage.out


## Integration Test

export DB_HOST=localhost && export DB_USER=postgres && export DB_PASSWORD=postgres1234 && export DB_PORT=5432 && export DB_NAME=inventory && go test --tags=integration -coverprofile=coverage.out