package ws

type Room struct {
	ID      string
	Name    string
	Clients map[string]*Client
}

type Hub struct {
	Rooms map[string]*Room
}

func NewHub() *Hub {
	return &Hub{Rooms: make(map[string]*Room)}
}
