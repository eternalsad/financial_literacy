# How to build:
    make

# How to run:
    ./app

# How to test functionality?
    you can use tomcat
    but in our school we can't so I used regular cURL
### for adding category use:
    curl -X POST localhost:8080/add/category -H 'Content-Type: application/json' -d '{"name":"name1"}'
### for reading category use:
    curl -X POST localhost:8080/read/category/1
### for updating category use: 
    curl -X POST localhost:8080/update/category -H 'Content-Type: application/json' -d '{"id": 3, "name":"new2"}'
### for deleting category use:
    curl -X POST localhost:8080/delete/category/4

### for adding transaction you can use:
    curl -X POST localhost:8080/add/transaction -H 'Content-Type: application/json' -d '{"transaction_name":"first", "amount": 1.5, "isIncome": true, "comment": "gagarinskiy dom 3", "category_id": 3}'
### for reading transaction you can use:
    curl -X POST localhost:8080/read/transaction/1
### for updating transaction you can use:
    curl -X POST localhost:8080/update/transaction -H 'Content-Type: application/json' -d '{"id": 1,"transaction_name":"changed", "amount": 1.5, "isIncome": true, "comment": "changed_comment", "category_id": 3}'
### for deleting transaction:
    curl -X POST localhost:8080/delete/transaction/1

#### How does it work?
    for creating or updating transaction or category you need to pass a json
    for reading or deleting just pass an id of entry