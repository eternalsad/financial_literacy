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
    curl -X POST localhost:8080/add/transaction -H 'Content-Type: application/json' -d '{"transaction_name":"first", "amount": 1.5, "isIncome": true, "comment": "yahabalya", "category_id": 3}'
### for reading transaction you can use:

### for updating transaction you can use:
    curl -X POST localhost:8080/update/transaction -H 'Content-Type: application/json' -d '{"transaction_name":"first", "amount": 1.5, "isIncome": true, "comment": "yahabalya", "category_id": 3}'