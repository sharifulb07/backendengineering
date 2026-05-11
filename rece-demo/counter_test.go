package counter

import (
	"sync"
	"testing"
)

func TestCounterRace(t *testing.T) {

	c:=&Counter{}
	var wg sync.WaitGroup

	for i:=0; i<1000; i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			c.increment()
		}()
	}

	wg.Wait()

	t.Log("Final Value: ", c.value())
}