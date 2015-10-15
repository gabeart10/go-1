package main
import (
	"fmt"
	"math/rand"
	"input"
	"strings"
	c "github.com/gabeart10/go-1/colors"
)

func main() {
	for {
		fmt.Print(c.Clear)
		text := input.Ask("What do you want to know?\n")
		answers := []string{
			"It is certain",
			"It is decidedly so",
			"Without a doubt",
			"Yes definitely",
			"You may rely on it",
			"As I see it yes",
			"Most likely",
			"Outlook good",
			"Yes",
			"Signs point to yes",
			"Reply hazy try again",
			"Ask again later",
			"Better not tell you now",
			"Cannot predict now",
			"Concentrate and ask again",
			"Don't count on it",
			"My reply is no",
			"My sources say no",
			"Outlook not so good",
			"Very doubtful",
		}
		color := c.Random_color()
		if strings.Contains(text, "die") {
			fmt.Println(color + "Yes you will die!" + c.Reset)
		} else if strings.Contains(text, "love") {
			fmt.Println(color + "Love you wish!" + c.Reset)
		} else if text == 'bye' {
			input.Bye()
		} else {	
			fmt.Println(color + answers[rand.Intn(len(answers))] + c.Reset)
		}
		input.Ask("")
	}
}
