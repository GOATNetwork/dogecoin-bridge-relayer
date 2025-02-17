package eventbus

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestEventBus(t *testing.T) {
	utils.SkipCI(t)

	bus := NewBus()

	t.Log("test eventbus begin")

	testLen := 1000
	exist := make(chan struct{}, testLen)
	wg := sync.WaitGroup{}
	count := atomic.Uint64{}

	for i := 0; i < testLen; i++ {
		c := i
		if i%2 == 0 {
			c = 0
		}

		unknown := bus.SubscribeWithLen(EventUnknown{}, c)

		wg.Add(1)

		go func() {
			exist <- struct{}{}

			result := <-unknown
			t.Logf("subtest:index = %d, result = %v", i, result)
			count.Add(1)

			wg.Done()
		}()
	}

	<-exist
	bus.Publish(EventUnknown{}, "OK")
	t.Log("eventbus publish end")
	wg.Wait()
	assert.Equal(t, count.Load(), uint64(bus.SubscriberLen(EventUnknown{})))
	t.Log("test eventbus end")
}

func TestDefaultBus(t *testing.T) {
	type TssMsg1 struct{}

	type TssMsg2 struct{}

	assert.Equal(t, TssMsg1{}, TssMsg1{})
	assert.NotEqual(t, TssMsg1{}, TssMsg2{})

	var msg TssMsg1

	var msg1 TssMsg1

	var msg2 TssMsg2

	assert.Equal(t, TssMsg1{}, msg)
	assert.Equal(t, msg1, msg)
	assert.NotEqual(t, msg1, msg2)
}
