package pool

import (
	"math/big"
	"slices"
	"sync"
)

type BigIntPool struct {
	ids   []*big.Int
	items map[*big.Int]Task[*big.Int]
	sync.RWMutex
}

func NewBigIntPool() *BigIntPool {
	return &BigIntPool{
		ids:   make([]*big.Int, 0),
		items: make(map[*big.Int]Task[*big.Int]),
	}
}

func (t *BigIntPool) First() Task[*big.Int] {
	items := t.GetTopN(1)
	if len(items) == 0 {
		return nil
	}

	return items[0]
}

func (t *BigIntPool) Last() Task[*big.Int] {
	l := t.Len()

	items := t.GetTopN(int64(l))
	if len(items) == 0 {
		return nil
	}

	return items[l-1]
}

func (t *BigIntPool) Add(item Task[*big.Int]) {
	//if t.IsExist(item.TaskID()) {
	//	return
	//}
	t.Lock()
	defer t.Unlock()
	t.ids = append(t.ids, item.TaskID())
	t.items[item.TaskID()] = item
}

func (t *BigIntPool) Remove(id *big.Int) {
	t.Lock()
	defer t.Unlock()
	delete(t.items, id)
}

func (t *BigIntPool) Get(id *big.Int) Task[*big.Int] {
	t.RLock()
	defer t.RUnlock()

	return t.items[id]
}

func (t *BigIntPool) BatchGet(ids []*big.Int) []Task[*big.Int] {
	t.RLock()
	defer t.RUnlock()

	tasks := make([]Task[*big.Int], 0)

	for _, id := range ids {
		if t.items[id] != nil {
			tasks = append(tasks, t.items[id])
		}
	}

	return tasks
}

func (t *BigIntPool) IsExist(id *big.Int) bool {
	t.RLock()
	defer t.RUnlock()
	_, ok := t.items[id]

	return ok
}

func (t *BigIntPool) Len() int {
	t.RLock()
	defer t.RUnlock()

	return len(t.items)
}

func (t *BigIntPool) IsEmpty() bool {
	t.RLock()
	defer t.RUnlock()

	return len(t.ids) == 0
}

func (t *BigIntPool) Clear() {
	t.Lock()
	defer t.Unlock()
	t.clear()
}

func (t *BigIntPool) clear() {
	t.ids = t.ids[:0]
	t.items = make(map[*big.Int]Task[*big.Int])
}

func (t *BigIntPool) Reset() {
	t.Clear()
}

func (t *BigIntPool) GetTopN(N int64) []Task[*big.Int] {
	t.Lock()

	is := slices.IsSortedFunc(t.ids, func(a, b *big.Int) int { return a.Cmp(b) })
	if !is {
		slices.SortFunc(t.ids, func(a, b *big.Int) int { return a.Cmp(b) })
	}
	t.Unlock()
	t.RLock()
	defer t.RUnlock()

	items := make([]Task[*big.Int], 0, N)
	for _, id := range t.ids {
		item, ok := t.items[id]
		if ok {
			items = append(items, item)
			if len(items) == int(N) {
				break
			}
		}
	}

	return items
}

func (t *BigIntPool) RemoveTopN(rightId *big.Int) {
	t.Lock()
	defer t.Unlock()

	if len(t.items) == 0 {
		t.ids = make([]*big.Int, 0)
		return
	}

	is := slices.IsSortedFunc(t.ids, func(a, b *big.Int) int { return a.Cmp(b) })
	if !is {
		slices.SortFunc(t.ids, func(a, b *big.Int) int { return a.Cmp(b) })
	}

	if rightId.Cmp(t.ids[0]) == -1 {
		return
	}

	if rightId.Cmp(t.ids[0]) == 1 {
		t.clear()
		return
	}

	ids := t.ids[0:]
L:
	for index, id := range ids {
		if id.Cmp(rightId) <= 0 {
			delete(t.items, id)
			if id == rightId {
				t.ids = t.ids[index+1:]
				break L
			}
		}
	}
}
