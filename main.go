package main

import (
    "fmt"
    "os"
    "./git"
)

func main() {
    username, err := git.GetUserName()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    file, err := os.Create("data.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()
    file.WriteString(username + "\n")
}