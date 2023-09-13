package telegram

import (
	"log"
	"strconv"
	"strings"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("Got new command '%s' from '%s' (ChatID -> %s)", text, username, strconv.Itoa(chatID))

	// TODO: Сделать нормальную реализацую у Роуторов
	// --- Init Command ---
	if text == "/start" {
		return p.sendHello(chatID)
	}

	// --- TODO: TEST ---
	err := p.tgClient.SendMessageText(chatID, "Unknown Command")
	if err != nil {
		return err
	}

	return nil
}

func (p *Processor) sendHello(chatID int) error {
	return p.tgClient.SendMessageText(chatID, "msgHello")
}
