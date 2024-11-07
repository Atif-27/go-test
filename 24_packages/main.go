package main

import (
	"fmt"

	"github.com/Atif-27/podcast/auth"
	"github.com/Atif-27/podcast/user"
	"github.com/fatih/color"
)
func main() {
	auth.LoginWithAuth("Atif","1234567")
	newUser:= user.User{
		Username: "Atif",
		Password: "1234567",
	}
	color.Red("Username: %s", newUser.Username)
	session:=auth.GetSession()
	fmt.Println(session,newUser)
}