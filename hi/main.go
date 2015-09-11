package main

import (
  "bufio"
  "os"
  "fmt"
  c "github.com/gabeart10/go-1/colors"
)
func main() {
  r := bufio.NewReader(os.Stdin)
  fmt.Print(c.Clear)
  fmt.Println(c.Yellow + "What is your name?" + c.Reset)
  fmt.Print(c.Blue)
  line, err := r.ReadString('\n')
  fmt.Print(c.Reset)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  fmt.Print(c.Red + "Hello " + line + c.Reset)

}
