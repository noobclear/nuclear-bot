package messages

type Handler interface {
	Handle(ctx *Context, m Message) (success bool)
}