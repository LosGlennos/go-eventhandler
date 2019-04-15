# gull
A small eventhandling library for Go. 

Will register all events, and fire them as goroutines when `EmitEvent` is called.

# Usage

```go
import "github.com/losglennos/gull"

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
