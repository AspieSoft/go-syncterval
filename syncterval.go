package syncterval

import (
	"time"

	"github.com/alphadose/haxmap"
)

type FuncInterval struct {
	fn *func()
	t int64
	last int64
}

var funcIntervalListFast *haxmap.HashMap[uintptr, *FuncInterval] = haxmap.New[uintptr, *FuncInterval]()
var funcIntervalListSlow *haxmap.HashMap[uintptr, *FuncInterval] = haxmap.New[uintptr, *FuncInterval]()

func init(){
	lastMin := int64(0)

	go (func(){
		for {
			now := time.Now().UnixMilli()
	
			funcIntervalListFast.ForEach(func(i uintptr, fnInterval *FuncInterval) {
				if now - fnInterval.last >= fnInterval.t {
					fnInterval.last = now
					(*fnInterval.fn)()
				}
			})

			if now - lastMin >= int64(time.Minute.Milliseconds()) {
				lastMin = now
				funcIntervalListSlow.ForEach(func(i uintptr, fnInterval *FuncInterval) {
					if now - fnInterval.last >= fnInterval.t {
						fnInterval.last = now
						(*fnInterval.fn)()
					}
				})
			}
	
			time.Sleep(10 * time.Millisecond)
		}
	})()
}

func New(t time.Duration, fn func()){
	fnRef := &fn

	fnInterval := FuncInterval{fnRef, int64(t.Milliseconds()), 0}

	if t >= time.Minute {
		funcIntervalListSlow.Set(funcIntervalListSlow.Len(), &fnInterval)
	}else{
		funcIntervalListFast.Set(funcIntervalListSlow.Len(), &fnInterval)
	}
}
