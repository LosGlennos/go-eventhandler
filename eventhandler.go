//Package go_eventhandler provides an easy way to inject the behaviour of events in your application
package go_eventhandler

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

type event struct {
	s string
	f func()
}

type fn func()

var eventHandlers []event

//EmitEvent will fire the event specified and the functions that has been registered with that event.
//The events will be fired as go-routines.
func EmitEvent(s string) (int, error){
	var selectedEvents []fn
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

//Subscribe will add the event specified together with a function that later can be fired using EmitEvent.
//You will not be able to add the same combination multiple times.
func Subscribe(s string, f fn) error {
	err := checkIfSameEventAlreadyRegistered(s, f)
	if (err != nil) {
		return err
	}

	eventHandlers = append(eventHandlers, event{s, f})
	return nil
}

func checkIfSameEventAlreadyRegistered(s string, f fn) error {
	for _, event := range eventHandlers {
		if event.s == s && getFunctionName(event.f) == getFunctionName(f) {
			return errors.New(fmt.Sprintf("event %s with function %s is already registered", event.s, getFunctionName(f)))
		}
	}

	return nil
}

func getFunctionName (f fn) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
