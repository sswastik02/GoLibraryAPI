# Setup
Clone the project and install required go packages
```
git clone https://github.com/sswastik02/Books-API
go mod tidy
```

---
**Note** : `go mod tidy` executes `go get` on the imports in main.go, you can also `go get` each library individually

---

# Run
```
go run main.go
```


# Steps followed to build the Project

## Initialise Module for the project
We will add a `go.mod` file with link to our repo with the command in the root directory:
```
go mod init github.com/sswastik02/Books-API
```

## Basic Directory Structure
Created the files and folders manually, forming the following directory tree

```
    Books-API
    ├── go.mod
    ├── main.go
    ├── .env
    ├── models/
    └── storage/
```

## Writing to main.go
Next we start writing to the main.go file for : 
* Import Neccessary Go Libraries(Such as Fiber and Gorm)
* Importing .env
* Initialise Fiber Framework
* Write routes and functions

---
### Resources

[GoLang Setup and Tutorial](https://youtu.be/yyUHQIec83I)
[Postgres Setup for Ubuntu 20.04](https://www.cherryservers.com/blog/how-to-install-and-setup-postgresql-server-on-ubuntu-20-04)
[Playlist with GO-Fiber and PostgresQL](https://youtube.com/playlist?list=PL5dTjWUk_cPaKHFvmMct_VG5vIU4piYv4)

---