# About Project

## Simple Library Management System
### Features
```

Book : 
    Add Book          (Administrator)
    Get Book Info
    Get All Books Info
    Remove Book       (Administrator)

Users : 
    User Signup
    User Signin
    Refresh JWT token pair
    
    Admin Signup      (Administrator)
    Give Admin Role   (Administrator)
    Revoke Admin Role (Administrator) [Blacklisting JWT access token for its lifetime]
    Delete any User   (Administrator) [Blacklisting JWT access token for its lifetime]

Issue : 
    Add Issue    (Administrator)
    Get Issues
    Delete Issue (Administrator)

```
### Technologies
```
Backend Framework   : Go-Fiber
Database            : Postresql
ORM                 : GORM
Password Hashing    : HS256
Session Management  : JWT Token
Cache               : Redis
```


# Setup

## Local

Clone the project and install required go packages 
(replace `.env` contents with `.env.dev`)

You need postgres to be installed on the system
```
git clone https://github.com/sswastik02/Books-API
go mod tidy # or go mod download
```

---
**Note** : `go mod tidy` ensures that the go.mod file matches the source code in the module, you can also `go get` each library individually

---

### Create Admin
```
go run main.go -createadmin
```

### Run Server
```
go run main.go -runserver
```

## Docker

Alternatively, docker container can be built from the `docker-compose.yml` file
(replace `.env` contents with `.env.prod`)

You will need docker and docker-compose for this

### Create Admin
```
docker exec -it <go_fiber_container_name> go run main.go -createadmin
```
<span style="color:grey"> go fiber container name is the name of the container running the go fiber backend that is running on your system. You can view it by running `docker ps`</span>

### Run Server

```
docker-compose up --build
```


<span style="color:grey">Note : You need to make sure nothing is already running on the ports 8000 and 5432. If there is you need to stop them</span>

# Steps followed to build the Project Initially

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

<span style="color:grey"> The above .env file is outdated as it was during the inital steps of the project. You can have a look at `env.dev` or `env.prod` for reference</span>

---
### Resources

[GoLang Setup and Tutorial](https://youtu.be/yyUHQIec83I)<br>
[Postgres Setup for Ubuntu 20.04](https://www.cherryservers.com/blog/how-to-install-and-setup-postgresql-server-on-ubuntu-20-04)<br>
[Playlist with GO-Fiber and PostgresQL](https://youtube.com/playlist?list=PL5dTjWUk_cPaKHFvmMct_VG5vIU4piYv4)
[GORM models](https://gorm.io/docs/models.html)<br>
[Posgresql Basic create user and database](https://medium.com/coding-blocks/creating-user-database-and-adding-access-on-postgresql-8bfcd2f4a91e)<br>
[Implementing Password Authentication ](https://www.sohamkamani.com/golang/password-authentication-and-storage/)<br>
[Implementing JWT in Fiber](https://github.com/gofiber/jwt)<br>
[Implemeting JWT with refresh token](https://medium.com/monstar-lab-bangladesh-engineering/jwt-auth-in-go-part-2-refresh-tokens-d334777ca8a0)
[Dockerizing Fiber with Postgresql](https://levelup.gitconnected.com/dockerized-crud-restful-api-with-go-gorm-jwt-postgresql-mysql-and-testing-61d731430bd8)
[Setting up Redis on Ubuntu](https://www.digitalocean.com/community/tutorials/how-to-install-and-secure-redis-on-ubuntu-18-04)
[Overview of the GORM object relational mapper](https://www.youtube.com/watch?v=nVD9acHituc)


---