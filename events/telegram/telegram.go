package telegram

import (
	"errors"
	"fmt"
	"github.com/Zigatase/telego"
	"github.com/Zigatase/telego/e"
	events "github.com/Zigatase/telego/events"
	"github.com/Zigatase/telego/types"
)

type Processor struct {
	tgClient *telego.Client
	offset   int
	// storage
}

var ErrUnknownEventType = errors.New("Unknown Message")

func NewProcessor(client *telego.Client) *Processor {
	return &Processor{
		tgClient: client,
	}
}

func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tgClient.Updates(p.offset, limit)
	if err != nil {
		return nil, e.Wrap("[Events/Telegram -> Fetch] Can't get events: ", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	response := make([]events.Event, 0, len(updates))
	for _, update := range updates {
		response = append(response, event(update))
	}

	p.offset = updates[len(updates)-1].ID + 1

	return response, nil
}

func (p *Processor) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		err := p.processMessage(event)
		if err != nil {
			return err
		}

	default:
		return e.Wrap("[Events/Telegram -> Process] Can't process message", ErrUnknownEventType)
	}

	return nil
}

func (p *Processor) processMessage(event events.Event) error {
	if event.ChatID == 0 && event.UserName == "" {
		return fmt.Errorf("[Telegram -> processMessage] Event data Null")
	}

	// TODO: Сделать нормальную реализацую у инцилизации комманд
	// --- Init Command ---
	if err := p.doCmd(event.Text, event.ChatID, event.UserName); err != nil {
		return e.Wrap("[Telegram -> proccessMessage] Can't process message: ", err)
	}

	return nil
}

func event(update types.Update) events.Event {
	updateType := fetchType(update)

	res := events.Event{
		Type: updateType,
		Text: fetchText(update),
	}

	if updateType == events.Message {
		res.ChatID = update.Message.Chat.Id
		res.UserName = update.Message.From.UserName
	}

	return res
}

func fetchType(update types.Update) events.Type {
	if update.Message == nil {
		return events.Unknown
	}

	return events.Message
}

func fetchText(update types.Update) string {
	if update.Message == nil {
		return ""
	}
	return update.Message.Text
}
