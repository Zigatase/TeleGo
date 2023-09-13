package example

import (
	tgClient "github.com/Zigatase/telego"
	event_consumer "github.com/Zigatase/telego/consumer/event-consumer"
	"github.com/Zigatase/telego/events/telegram"
	"log"
)

const (
	batchSize = 100
)

func main() {
	eventPc := telegram.NewProcessor(tgClient.New("6394482620:AAFJHmuK_21DSnAzNt2_eQEeuuq1VQ5G5n4"))

	log.Print("start")

	consumer := event_consumer.NewConsumer(eventPc, eventPc, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatalf("bot is stopped", err)
	}
}
