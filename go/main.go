package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	cli := NewClient(os.Getenv("OPENAI_APIKEY"))

	for {
		fmt.Printf("Send: ")
		reader := bufio.NewReader(os.Stdin)
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		if msg == "\n" {
			continue
		}

		reply, err := cli.Say("user", msg)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("From %s\n%+s\n", reply.Role, reply.Message)
		fmt.Printf("\n\n")
	}
}
