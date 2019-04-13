package eventhandler

import "testing"


func TestEmitEvent(t *testing.T) {
	Subscribe("test", Happen)
	Subscribe("test", Happen2)
	EmitEvent("test")
}

func Happen() {
	a := 1
	a = a+2
}

func Happen2() {
	a := 1
	a = a+2
}
