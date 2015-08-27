package main

import (
  "bufio"
  "fmt"
  "os"
)
func main() {
  r := bufio.NewReader(os.Stdin)

  println("What is your name?:")
  line, err := r.ReadString('\n')
  if err != nil {
    println(err)
    os.Exit(1)
  }
  fmt.Println(line)

}
