package adaptor

import "fmt"

type DiscordAdaptor struct{}

func (s *DiscordAdaptor) SendMesage(message string) {
	fmt.Print("Send")
}