package main
import (
	"fmt"
	"math/rand"
	"bufio"
	"os"
	"strings"
	c "github.com/gabeart10/go-1/colors"
)

func main() {
	for {
		r := bufio.NewReader(os.Stdin)
		fmt.Print(c.Clear)
		fmt.Println(c.Yellow + "What do you want to know?" + c.Reset)
		fmt.Print(c.Red)
		text, err := r.ReadString('\n')
		fmt.Print(c.Reset)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		text = strings.TrimSpace(text)
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
		} else {	
			fmt.Println(color + answers[rand.Intn(len(answers))] + c.Reset)
		}
		_, err_new := r.ReadString('\n')
		if err_new != nil {
			fmt.Println(err_new)
			os.Exit(1)
		}
	}
}
