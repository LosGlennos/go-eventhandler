# go-eventhandler
A small eventhandling library for Go. 

Will register all events, and fire them as goroutines when `EmitEvent` is called.

# Usage

```go
import "github.com/losglennos/go-eventhandler"

func SubscribeEvent() {
	eventhandler.Subscribe("foo", bar)
}

func EmitEvent() {
	i, e := eventhandler.EmitEvent("foo")
}

func bar() {
	//run this function on event
}
```

# Installation

```
go get github.com/losglennos/go-eventhandler
```
