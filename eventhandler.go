package eventhandler

type event struct {
	s string
	f func()
}

type fn func()

var eventHandlers = []event{}

func EmitEvent(s string) {
	var selectedEvents = []fn{}
	for _, event := range eventHandlers{
		if event.s == s {
			selectedEvents = append(selectedEvents, event.f)
		}
	}

	for _, fn := range selectedEvents {
		go fn()
	}
}

func Subscribe(s string, f fn) {
	eventHandlers = append(eventHandlers, event{s, f})
}
