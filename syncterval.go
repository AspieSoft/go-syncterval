package syncterval

import (
	"time"

	"github.com/alphadose/haxmap"
)

type funcInterval struct {
	fn *func()
	t int64
	last int64
}

var funcIntervalListFast *haxmap.Map[uintptr, *funcInterval] = haxmap.New[uintptr, *funcInterval]()
var funcIntervalListSlow *haxmap.Map[uintptr, *funcInterval] = haxmap.New[uintptr, *funcInterval]()

func init(){
	lastMin := int64(0)

	go (func(){
		for {
			now := time.Now().UnixMilli()
	
			funcIntervalListFast.ForEach(func(i uintptr, fnInterval *funcInterval) bool {
				if now - fnInterval.last >= fnInterval.t {
					fnInterval.last = now
					(*fnInterval.fn)()
				}
				return true
			})

			if now - lastMin >= int64(time.Minute.Milliseconds()) {
				lastMin = now
				funcIntervalListSlow.ForEach(func(i uintptr, fnInterval *funcInterval) bool {
					if now - fnInterval.last >= fnInterval.t {
						fnInterval.last = now
						(*fnInterval.fn)()
					}
					return true
				})
			}
	
			time.Sleep(10 * time.Millisecond)
		}
	})()
}

func New(t time.Duration, fn func()){
	fnRef := &fn

	fnInterval := funcInterval{fnRef, int64(t.Milliseconds()), 0}

	if t >= time.Minute {
		funcIntervalListSlow.Set(funcIntervalListSlow.Len(), &fnInterval)
	}else{
		funcIntervalListFast.Set(funcIntervalListSlow.Len(), &fnInterval)
	}
}
