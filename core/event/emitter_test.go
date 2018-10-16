package event

import (
	"testing"
)

func TestEmitter_Emit(t *testing.T) {
	emitter := NewEvent()

	testCount := 0
	test := func(emitter interface{}, args *Args) {
		testCount++
	}

	tcb := emitter.GetHandler().Always(test)

	// +1
	emitter.Emit(nil, nil)
	emitter.GetHandler().Always(test)

	// +2
	emitter.Emit(nil, nil)
	emitter.GetHandler().UnListen(tcb)

	// +1
	emitter.Emit(nil, nil)

	if testCount != 4 {
		t.Errorf("Event emission count must be %d but is %d", 4, testCount)
	}
}