package main

import (
  "bufio"
  "os"
  "fmt"
)
func main() {
  r := bufio.NewReader(os.Stdin)

  fmt.println("What is your name?")
  line, err := r.ReadString('\n')
  if err != nil {
    fmt.println(err)
    os.Exit(1)
  }
  fmt.Print("Hello " + line)

}
