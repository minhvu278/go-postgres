package main

import (
	"fmt"
	"go-postgres/driver"
	"go-postgres/repository/repoimpl"
	models "go-postgres/model"
)

const (
	host = "localhost"
	port = "5432"
	user = "postgres"
	password = "root"
	dbname = "postgres"
)

func main() {
	db := driver.Connect(host, port, user, password, dbname)

	err := db.SQL.Ping()
	if err != nil {
		panic(err)
	}

	userRepo := repoimpl.NewUserRepo(db.SQL)
	mv := models.User{
		ID: 1,
		Name: "Minh Vu",
		Gender: "Male",
		Email: "mv@gmail.com",
	}
	dt := models.User{
		ID: 2,
		Name: "Dan Truong",
		Gender: "Male",
		Email: "dt@gmail.com",
	}
	users, err := userRepo.Select()
	if err != nil {
		fmt.Println("Error when select user: ", err.Error())
	}

	err = userRepo.Insert(mv)
	if err != nil {
		fmt.Println("Error when insert first record: ", err.Error())
	}

	err = userRepo.Insert(dt)
	if err != nil {
		fmt.Println("Error when insert second record: ", err.Error())
	}



	for i := range users{
		fmt.Println(users[i])
	}
}
