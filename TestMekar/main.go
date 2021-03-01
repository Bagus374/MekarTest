package main

import (
	"TestMekar/driver"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"

	ur "TestMekar/users/repository"
	uh "TestMekar/users/handler"
	us "TestMekar/users/service"
)

func init()  {
	gotenv.Load()
}

func main()  {
	port := os.Getenv("PORT")
	db :=driver.Connect()
	defer db.Close()
	driver.InitTable(db)

	router := mux.NewRouter().StrictSlash(true)

	userRepo := ur.CreateRepoImpl(db)
	userService := us.CreateServiceImpl(userRepo)
	uh.CreateUserHandler(router, userService)

	fmt.Println("Started Web Server at : ", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal()
	}
}
