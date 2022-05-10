# About Project

## Simple Library Management System
### Features
```
Add Book
Get Book Info
Get All Books Info
Remove Book
Users
```
### Technologies
```
Backend Framework   : Go-Fiber
Database            : Postresql
ORM                 : GORM
Password Hashing    : HS256
Session Management  : JWT Token
```


# Setup
Clone the project and install required go packages
```
git clone https://github.com/sswastik02/Books-API
go mod tidy
```

---
**Note** : `go mod tidy` ensures that the go.mod file matches the source code in the module, you can also `go get` each library individually

---

# Run
```
go run main.go
```


# Steps followed to build the Project

## <u> Initialise Module for the project</u>
We will add a `go.mod` file with link to our repo with the command in the root directory:
```
go mod init github.com/sswastik02/Books-API
```

## <u> Basic Directory Structure</u>
Created the files and folders manually, forming the following directory tree

```
    Books-API
    ├── go.mod
    ├── main.go
    ├── .env
    ├── models/
    └── storage/
```

## <u> Writing to main.go</u>
Next we start writing to the main.go file for : 
* Import Neccessary Go Libraries(Such as Fiber and Gorm)
* Importing .env
* Connect with Database
* Initialise Fiber Framework
* Write routes and functions

## <u> Writing Models</u>
We write the book model required to store in the database in the models folder along with the migrate method

## <u> Writing storage</u>
This Step includes making configuration to connect with postgresSQL

## <u> Writing .env file </u>

You need to create a postgres database on your localhost before writing the env file

<span style="color:grey">In the following, keep in mind the < angular brackets > are to denote variables you have to set yourself. In the actual, omit the <> </span>

The basic structure of the file looks like : 

```

DB_HOST=localhost
DB_PORT=5432
DB_USER=<username>
DB_PASS=<password>
DB_NAME=<dbname>
DB_SSLMODE=disable
JWT_SECRET=<jwtsecret>

```

---
### Resources

[GoLang Setup and Tutorial](https://youtu.be/yyUHQIec83I)<br>
[Postgres Setup for Ubuntu 20.04](https://www.cherryservers.com/blog/how-to-install-and-setup-postgresql-server-on-ubuntu-20-04)<br>
[Playlist with GO-Fiber and PostgresQL](https://youtube.com/playlist?list=PL5dTjWUk_cPaKHFvmMct_VG5vIU4piYv4)
[GORM models](https://gorm.io/docs/models.html)<br>
[Posgresql Basic create user and database](https://medium.com/coding-blocks/creating-user-database-and-adding-access-on-postgresql-8bfcd2f4a91e)<br>
[Implementing Password Authentication ](https://www.sohamkamani.com/golang/password-authentication-and-storage/)<br>
[Implementing jwt in Fiber](https://github.com/gofiber/jwt)<br>

---