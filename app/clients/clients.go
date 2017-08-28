package clients

type Clients struct {
	Wit WitClient
}

func NewClients() *Clients {
	return &Clients{
		Wit: WitClient{},
	}
}