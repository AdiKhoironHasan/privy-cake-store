# privy-cake-store

I use existing libs :
- Echo Router
- Viper, for config management
- Sqlx, for database connection
- Testify, for unit testing

I'am build a service to manage data articles using golang, using mysql for database. 
This service has database migration, logging, and unit testing.

I created 5 endpoints :
1. [POST] /cakes to add a new cake
2. [GET] /cakes to get the list of cakes. Sort the cakes by highest rating first with optional query parameters:
   - a. query: to search keywords in the cake title and description
   - b. author: filter by rating points.
3. [GET] /cakes/:id to get detail of cake by :id
4. [PUT] /cakes/:id to update an cake by :id
5. [DELETE] /cakes/:id to delete an cake by :id

# How to run in local
## clone the repo:
- ```$ git clone https://github.com/AdiKhoironHasan/privy-cake-store.git```

## Setup after cloning the repo:
Run this command on your terminal to prepare dependencies:
- ```$ cd privy-cake-store```
- ```$ go get all```
- ```$ go mod tidy```

## Configure environment:
Do this following actions for set up your configuration:
- copy config/config-dev.yaml.example to config/config-dev.yaml
- complete the necessary credentials such as mysql databases according to the existing format

## Database migration :
I use mysql for database.
you can create database tables by migration, but before that you have to create a new database on your RDBMS.
- ```$ migrate -database "mysql://userDB:passwordDB@tcp(hostDB:portDB)/yourDB" -path pkg/database/migrations up```

## Run service :
You can run the service by using the following command, after that the service is ready to use.
 - ```$ go run main.go```

# How to run with docker
You can run service in docker with this command:
- ```$ docker-compose -f docker-compose.yml up -d```

# How to use service
You can deploy a service with consumption to an already created API endpoint, you can use the postman tool. To make it easier to use, I've created a workspace for it, and it's ready to go. don't forget to use postman with desktop agent if the service is running on lokalhost.
## Postman Link: 
- ```https://www.postman.com/lively-comet-875863/workspace/privy-cake-store/collection/18402968-dd5be257-5653-44e8-b21a-202e76fab65c?action=share&creator=18402968```

# How to run unit test
This step to run unit test:
1. First you must go to service folder
 ```cd internal/services```
2. ```go test -v -cover```
You can see the coverage testing in each package by open the project with vscode, choose the testing file, right click then choose "Go:Toogle Test Coverage in Current Package"

## summary of unit test
> === RUN   TestAddNewCake_Success\
> --- PASS: TestAddNewCake_Success (0.00s)\
> === RUN   TestAddNewCake_Error\
> --- PASS: TestAddNewCake_Error (0.00s)\
> === RUN   TestShowAllCake_Success\
> --- PASS: TestShowAllCake_Success (0.00s)\
> === RUN   TestShowAllCake_Error\
> --- PASS: TestShowAllCake_Error (0.00s)\
> === RUN   TestShowAllCake_ErrorQuery\
> --- PASS: TestShowAllCake_ErrorQuery (0.00s)\
> === RUN   TestShowAllCake_ErrorRating\
> --- PASS: TestShowAllCake_ErrorRating (0.00s)\
> === RUN   TestShowCakeByID_NotFound\
> --- PASS: TestShowCakeByID_NotFound (0.00s)\
> === RUN   TestShowCakeByID_Success\
> --- PASS: TestShowCakeByID_Success (0.00s)\
> === RUN   TestUpdateCake_Success\
> --- PASS: TestUpdateCake_Success (0.00s)\
> === RUN   TestUpdateCake_Error\
> --- PASS: TestUpdateCake_Error (0.00s)\
> === RUN   TestDeleteCake_Success\
> --- PASS: TestDeleteCake_Success (0.00s)\
> === RUN   TestDeleteCake_Error\
> --- PASS: TestDeleteCake_Error (0.00s)\
> PASS\
> coverage: 94.4% of statements

# Thank You
If there are problems or want to know more information about me, please contact via linkedin via the following link https://www.linkedin.com/in/adi-khoiron-hasan or by sending an email to adikhoironhasan@gmail.com. Thank You :)
