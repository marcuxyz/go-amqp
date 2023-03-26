package main

import (
	"fmt"

	"github.com/marcuxyz/go-apis/messages"
)

func main() {
	message := messages.NewMessage()
	// message.Add(messages.Message{
	// 	Title:   "Hello",
	// 	Content: "Send message by server",
	// })

	fmt.Println(message.All())
}
