package eventsystem

import "engo.io/engo"

var (
	Mailbox        *engo.MessageManager
)

func init() {
	Mailbox = &engo.MessageManager{}
}
