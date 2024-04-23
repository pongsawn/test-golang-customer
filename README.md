# test-golang-customer

1.Curl เส้น Create
curl --location 'localhost:8888/customers/create' \
--header 'Content-Type: application/json' \
--data '{
    "id": 100,
    "name": "Potae",
    "age": 27
}'

2.Curl เส้น Read
curl --location 'localhost:8888/customers/read'

3.Curl เส้น Update
curl --location --request PUT 'localhost:8888/customers/update/100' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Boing",
    "age": 24
}'

4.Curl เส้น Delete  (ตรงเลข 100 คือหมายเลข id)
curl --location --request DELETE 'localhost:8888/customers/delete/100'


