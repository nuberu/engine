package event

type Emitter struct {
	handler *Handler
	listeners []*Listener
}

func NewEvent() *Emitter {
	ev := &Emitter{
		handler: &Handler{},
		listeners: make([]*Listener, 0),
	}

	ev.handler.parent = ev
	return ev
}

func (eh *Emitter) eventExists(ev *Listener) bool {
	for _, a := range eh.listeners {
		if a == ev {
			return true
		}
	}
	return false
}

func (eh *Emitter) GetHandler() *Handler {
	return eh.handler
}

func (eh *Emitter) AddListener(listener *Listener) {
	if !eh.eventExists(listener) {
		eh.listeners = append(eh.listeners[0:], listener)
	}
}

func (eh *Emitter) RemoveListener(listener *Listener) {
	for i, a := range eh.listeners {
		if a == listener {
			eh.listeners[i] = eh.listeners[len(eh.listeners)-1]
			eh.listeners = eh.listeners[:len(eh.listeners)-1]
		}
	}
}

func (eh *Emitter) Emit(emitter interface{}, args *Args) {
	var expired []*Listener

	for _, a := range eh.listeners {
		a.callback(emitter, args)

		if a.strategy == strategyOnce {
			expired = append(expired, a)
		}
	}

	for _, list := range expired {
		eh.RemoveListener(list)
	}
}