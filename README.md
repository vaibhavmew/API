# API
A Simple API made using two golang frameworks. (Gin and Gorm). The database used is MySQL.

### Features
A CRUD Application to create, get, update and delete users from the database.

### Instructions

### Open the MySQL Command Line Client and run the following commands to create the database.

```sh
1. create database example;
```
```sh
2. use example;
```
```sh
3. CREATE TABLE `users` (
	`id` int AUTO_INCREMENT,
	`name` varchar(40) NOT NULL,
	`dob` date NOT NULL,
	`address` varchar(150) NOT NULL,
	`city` varchar(30) NOT NULL,
	`state` varchar(30) NOT NULL,
	`pincode` varchar(6) NOT NULL,
	PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

### Enter the main directory and run the following commands

```sh
1. go get .
```
```sh
2. go run .
```

### Successful Output
```sh
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /users                    --> main.GetUsers (3 handlers)
[GIN-debug] GET    /user/:id                 --> main.GetUser (3 handlers)
[GIN-debug] POST   /create                   --> main.CreateUser (3 handlers)
[GIN-debug] PATCH  /user/:id                 --> main.UpdateUser (3 handlers)
[GIN-debug] DELETE /user/:id                 --> main.DeleteUser (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```

### API
```
1. Get all the users from the database.
   Request - GET
   Endpoint - http://localhost:8080/users
```
JSON Output
```sh
[
    {
        "id": 3,
        "name": "Vaibhav",
        "dob": "2021-10-10",
        "address": "Neela Apt",
        "city": "Mumbai",
        "state": "Maharashtra",
        "pincode": "400061"
    },
    {
        "id": 4,
        "name": "Raj",
        "dob": "2010-04-05",
        "address": "Vishwadeep Apt",
        "city": "Mumbai",
        "state": "Maharashtra",
        "pincode": "400067"
    }
]
```
```
2. Get the user based on the ID.
   Request - GET
   Endpoint - http://localhost:8080/user/<id>
```
JSON Output
```sh
{
        "id": 12,
        "name": "Ameya",
        "dob": "2021-10-10",
        "address": "Neela Apt",
        "city": "Mumbai",
        "state": "Maharashtra",
        "pincode": "400061"
}
```
```
3. Create a new user.
   Request - POST
   Endpoint - http://localhost:8080/create
```
Payload
```sh
{
        "name": "Vaibhav",
        "dob": "2021-10-10",
        "address": "Neela Apt",
        "city": "Mumbai",
        "state": "Maharashtra",
        "pincode": "400061"
}
```
```
4. Update any of the user details.
   Request - PATCH
   Endpoint - http://localhost:8080/user/<id>
```
Payload
```sh
{
        "address": "Pachasheel",
        "city": "Pune",
}
```
```
5. Delete the user from the database.
   Request - DELETE
   Endpoint - http://localhost:8080/user/<id>
```
