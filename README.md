![Build status](https://ci.appveyor.com/api/projects/status/6vv0gorlpag2m1dk?svg=true)

# go-eventhandler
A small eventhandling library for Go. 

Will register all events, and fire them as goroutines when `EmitEvent` is called.

# Usage

```go
import "github.com/losglennos/go-eventhandler"

func SubscribeEvent() {
	err := eventhandler.Subscribe("foo", bar)
}

func EmitEvent() {
	i, err := eventhandler.EmitEvent("foo")
}

func bar() {
	//run this function on event
}
```

# Installation

```
go get github.com/losglennos/go-eventhandler
```
