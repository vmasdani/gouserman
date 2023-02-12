package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from gouserman!\n"))
}

var DB *gorm.DB

func main() {
	fmt.Println("Starting...")
	time.Sleep(time.Second * 10)
	fmt.Println("Starting done...")

	jwtSecret, err := GetJwtSecret()

	if err != nil {
		panic(err)
	}
	fmt.Println("Got jwt secret:" + jwtSecret)

	r := mux.NewRouter()

	Route(r)

	// DB stuff
	dsn := "host=" + os.Getenv("POSTGRES_HOST") + " user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " port=" + os.Getenv("POSTGRES_PORT") + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB = db

	db.AutoMigrate(&User{})

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":5430", r))

}
