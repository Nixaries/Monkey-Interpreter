package main

import (
  "fmt"
  "os"
  "os/user"
  "interpreter/repl"
)

func main() {
  user, err := user.Current()
  if err != nil {
    panic(err)
  }
  fmt.Printf("Hello %s! This is the Monkey programming language!", user.Username)
  fmt.Printf("Type commands below")
  repl.Start(os.Stdin, os.Stdout)
}
