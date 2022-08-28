package syncterval

import (
	"errors"
	"testing"
	"time"
)

func Test(t *testing.T) {
	ranFn := false
	fn := func(){
		ranFn = true
	}

	New(10 * time.Millisecond, fn)

	time.Sleep(100 * time.Millisecond)

	if !ranFn {
		t.Error(errors.New("Function Did Not Run"))
	}
}
