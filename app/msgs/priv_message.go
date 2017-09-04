package msgs

import (
	"fmt"
)

type PrivMessage struct {
	FromUser  string
	ToChannel string
	Text      string
}

func (pm *PrivMessage) FormatResponse() string {
	atUser := fmt.Sprintf("@%s ", pm.FromUser)
	t := fmt.Sprintf("%s%s", atUser, pm.Text)

	return "PRIVMSG " + pm.ToChannel + " :" + t
}
