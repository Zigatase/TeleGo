package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Event struct {
	// --- Base Data ---
	Type Type
	Text string

	// --- Tg Data ---
	ChatID   int
	UserName string
}
