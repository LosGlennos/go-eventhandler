# gull
A small eventhandling library for Go. 

Will register all events, and fire them as goroutines when `EmitEvent` is called.

# Usage

```go
import "github.com/losglennos/gull"

func SubscribeEvent() {
	eventhandler.Subscribe("hello", hello)
}

func EmitEvent() {
	eventhandler.EmitEvent("hello")
}

func hello() {
	//run this function on event
}
```
