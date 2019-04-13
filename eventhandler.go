package gull

import (
	"errors"
	"fmt"
)

type event struct {
	s string
	f func()
}

type fn func()

var eventHandlers = []event{}

func EmitEvent(s string) (int, error){
	var selectedEvents = []fn{}
	for _, event := range eventHandlers{
		if event.s == s {
			selectedEvents = append(selectedEvents, event.f)
		}
	}

	if len(selectedEvents) <= 0 {
		return 0, errors.New(fmt.Sprintf("No event registered for: %s", s) )
	}

	for _, fn := range selectedEvents {
		go fn()
	}

	return len(selectedEvents), nil
}

func Subscribe(s string, f fn) {
	eventHandlers = append(eventHandlers, event{s, f})
}
