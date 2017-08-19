package messages

type PrivMessage struct {
	FromUser  string
	ToChannel string
	Text      string
}

func (pm *PrivMessage) FormatResponse() string {
	return "PRIVMSG " + pm.ToChannel + " :" + pm.Text
}
