package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sswastik02/GoLibraryAPI/api"
	database "github.com/sswastik02/GoLibraryAPI/database"
	"github.com/sswastik02/GoLibraryAPI/models"
)

func main(){

	// ---------------------------------- Command Line Arguments -------------------------------------

	CREATEADMIN := flag.Bool("createadmin",false,"Create Admin User")
	RUNSERVER := flag.Bool("runserver",false,"Run the server")
	flag.Parse()

	if !*CREATEADMIN && !*RUNSERVER {
		log.Fatal(
			"\n\nUsage : go run main.go -[COMMAND] or ./[binary] -[COMMAND]",
			"\n\nPossible Commands are : ",
			"\n runserver \t- To Run the Server",
			"\n createadmin \t- To enter a admin into the database",
		)
	}

	// ----------------------------------- Load the dotenv file ----------------- 

	err:=godotenv.Load(".env")
	if(err != nil){
		log.Fatal("Could not import .env file")
	}

	// -------------------------------- Load Configuration from env file -----------------------

	config:= &database.Config{ // & creates a pointer to a structure with data as declared
		Host: os.Getenv("DB_HOST") ,
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:os.Getenv("DB_NAME"),
		SSLMode:os.Getenv("DB_SSLMODE"),
	}

	api.InitializeJwtSecret(os.Getenv("JWT_SECRET"))

	db, err := database.NewConnection(config)

	if(err != nil){
		log.Fatal("Could not load database")
	}

	// -------------------------------- Migrate the Database -----------------------

	err = models.Migrate(db)
	if(err != nil){
		log.Fatal("Could not Migrate Database")
	}

	// -----------------------------Initialise a Repository containing the Database--------------------

	r := api.Repository{
		DB: db,
	}

	//------------------------------------- Run Server and Create Admin ------------------------------------

	if(*CREATEADMIN) {
		createadmin(&r)
	}

	if(*RUNSERVER) {
		runserver(&r)
	}
	
	
}

func runserver(r *api.Repository) {
// ------------------------- Initialise Fiber Framework and routes using the repository-------------------
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8000")
}

func createadmin(r *api.Repository) {
	var username string
	var password string

	fmt.Print("\nEnter username : ")
	fmt.Scanf("%s",&username)
	fmt.Print("\nEnter password : ")
	fmt.Scanf("%s",&password)

	fmt.Printf("\nCreating user with\n Username : %s \n Password : %s\n",username,password)
	user := models.User{
		Username: username,
		Password: password,
		Admin: true,
	}

	_, err := r.CreateUser(&user)

	if err != nil {
		log.Fatalf("\nCould not create Admin\n Error : %s",err.Error())
	}
}