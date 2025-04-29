package telegram

import "github.com/yowie645/ReadItLaterBot/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	// storage
}

func New(client *telegram.Client) {

}
