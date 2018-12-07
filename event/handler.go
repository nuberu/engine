package event

type Handler struct {
	parent *Emitter
}

func (eh *Handler) Once(cb func(interface{}, *Args)) *Listener {
	ec := &Listener{
		strategy: strategyOnce,
		callback: cb,
	}

	eh.parent.AddListener(ec)

	return ec
}

func (eh *Handler) Always(cb func(interface{}, *Args)) *Listener {
	ec := &Listener{
		strategy: strategyAlways,
		callback: cb,
	}

	eh.parent.AddListener(ec)

	return ec
}

func (eh *Handler) UnListen(listener *Listener) {
	eh.parent.RemoveListener(listener)
}