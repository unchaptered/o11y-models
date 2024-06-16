package outboundadaptor

type MessageSender interface {
	SendMessage(message string)
}