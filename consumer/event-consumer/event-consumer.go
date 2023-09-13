package event_consumer

import (
	"github.com/Zigatase/telego/events"
	"log"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func NewConsumer(fetcher events.Fetcher, processor events.Processor, baseSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: baseSize,
	}
}

func (c Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[EventConsumer -> Start] consumer: %s", err.Error())

			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		if err := c.handleEvents(gotEvents); err != nil {
			log.Printf("[Consumer -> start -> handleEvents] Error: %s", err)

			continue
		}
	}
}

/*
TODO: Проблемы
 1. Потеря событий, Возвращения в хранилище, Фоллбек, Потверждение сдвигов оффсетов
 2. Обработка всей пачки евентов разом (варианты решения останавливатся после 1 ошибки, или добавить счетчик)
 3. !Параллельная обработка! (sync.WaitGroup{})
*/
func (c Consumer) handleEvents(events []events.Event) error {
	for _, event := range events {
		log.Printf("[Consumer] Got new event-consumer: %s", event.Text)

		// --- TODO: пароблема в логике, если евент не обработался, то его пропустит и не обработает ---
		if err := c.processor.Process(event); err != nil {
			log.Printf("[Consumer -> handleEvents] Can't handle event-consumer: %s", err.Error())

			continue
		}
	}

	return nil
}
