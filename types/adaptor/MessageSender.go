package adaptor

type MessageSender interface {
	SendMessage(message string)
}