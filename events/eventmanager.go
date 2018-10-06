package events

var (
	Mailbox        *EventManager
	CaptureChannel chan BaseEvent
)

func init() {
	Mailbox = &EventManager{}
	CaptureChannel = make(chan BaseEvent)
}

type EventMap map[string][]EventHandler

func (em EventMap) clearKey(key string) {
	em[key] = em[key][:0]
}

type EventHandler func(msg BaseEvent)

// MessageManager manages messages and subscribed handlers
type EventManager struct {
	listeners         EventMap
	onceListeners     EventMap
}

// Dispatch sends a message to all subscribed handlers of the message's type
func (mm *EventManager) Dispatch(message BaseEvent) {
	CaptureChannel <- message

	onceHandlers := mm.onceListeners[message.Type()]
	onceHandlerCopy := make([]EventHandler, len(onceHandlers))
	copy(onceHandlerCopy, onceHandlers)
	if mm.onceListeners != nil {
		mm.onceListeners.clearKey(message.Type())
	}
	for _, handler := range onceHandlerCopy {
		handler(message)
	}


	handlers := mm.listeners[message.Type()]
	for _, handler := range handlers {
		handler(message)
	}
}

// Listen subscribes to the specified message type and calls the specified handler when fired
func (mm *EventManager) Listen(messageType string, handler EventHandler) {
	if mm.listeners == nil {
		mm.listeners = make(EventMap)
	}
	mm.listeners[messageType] = append(mm.listeners[messageType], handler)
}

// Listen subscribes to the specified message type and calls the specified handler when fired
func (mm *EventManager) ListenOnce(messageType string, handler EventHandler) {
	if mm.onceListeners == nil {
		mm.onceListeners = make(EventMap)
	}
	mm.onceListeners[messageType] = append(mm.onceListeners[messageType], handler)
}
