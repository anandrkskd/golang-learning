package main

import (
    "fmt"
    "os"
    "os/user"
)

func main() {

    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    // Current User
    fmt.Println("Hi " + user.Name + " (id: " + user.Uid + ")")
    fmt.Println("Username: " + user.Username)
    fmt.Println("Home Dir: " + user.HomeDir)

    fmt.Println("Real User: " + os.Getenv("SUDO_USER"))
    if  (os.Getenv("SUDO_USER")==""){
	fmt.Print("hello bro")
    }
}
