package msgs

type PingMessage struct {
	Host string
}

func (pm *PingMessage) FormatResponse() string {
	return "PONG " + pm.Host
}
