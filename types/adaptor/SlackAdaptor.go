package adaptor

import "fmt"

type SlackAdaptor struct{}

func (s *SlackAdaptor) SendMesage(message string) {
	fmt.Print("Send")
}