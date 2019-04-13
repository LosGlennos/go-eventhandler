package gull

import (
	"testing"
	"time"
)

var a = 0
var b = 0

func TestEmitEvent(t *testing.T) {
	Subscribe("test", happen)
	Subscribe("test", happen2)
	_, _ = EmitEvent("test")

	time.Sleep(10 * time.Millisecond)

	aGot := a
	bGot := b
	if aGot != 1 {
		t.Errorf("a = %d, want 1", aGot)
	}

	if bGot != 2 {
		t.Errorf("b = %d, want 2", bGot)
	}
}

func TestEmitEventCount(t *testing.T) {
	Subscribe("testcount", happen)
	Subscribe("testcount", happen2)
	i, _ := EmitEvent("testcount")

	if i != 2 {
		t.Errorf("Should been 2, was: %d", i)
	}
}

func TestEmitEventError(t *testing.T) {
	_, e := EmitEvent("noevent")

	if e == nil {
		t.Error("Should return error, didn't")
	}
}

func happen() {
	a = 1
}

func happen2() {
	b = 2
}
