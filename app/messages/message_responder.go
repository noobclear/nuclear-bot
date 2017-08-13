package messages

type Responder interface {
	Respond(msg string, ctx *Context)
}

type MessageResponder struct {
	Handlers []Handler
}

func (mr *MessageResponder) Respond(msg string, ctx *Context) {
	m := NewMessage(msg)

	for _, h := range mr.Handlers {
		if h.Handle(ctx, m) {
			return
		}
	}
}

func (mr *MessageResponder) RegisterHandler(h Handler) {
	mr.Handlers = append(mr.Handlers, h)
}

func NewMessageResponder() Responder {
	mr := &MessageResponder{}
	mr.RegisterHandler(NewPingPongHandler())
	mr.RegisterHandler(NewIgnoreSelfHandler())
	mr.RegisterHandler(NewPrivMsgHandler())
	return mr
}