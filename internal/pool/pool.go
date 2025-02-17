package pool

import (
	"cmp"
	"fmt"
	"slices"
	"sync"
)

type Task[E any] interface {
	Type() int
	TaskID() E
}

type Pool[E cmp.Ordered] struct {
	ids   []E
	items map[E]Task[E]
	sync.RWMutex
}

func NewTaskPool[E cmp.Ordered]() *Pool[E] {
	return &Pool[E]{
		ids:   make([]E, 0),
		items: make(map[E]Task[E]),
	}
}

func (t *Pool[E]) First() Task[E] {
	items := t.GetTopN(1)
	if len(items) == 0 {
		return nil
	}

	return items[0]
}

func (t *Pool[E]) Last() Task[E] {
	l := t.Len()

	items := t.GetTopN(int64(l))
	if len(items) == 0 {
		return nil
	}

	return items[l-1]
}

func (t *Pool[E]) Add(item Task[E]) {
	t.Lock()
	defer t.Unlock()

	_, ok := t.items[item.TaskID()]
	if !ok {
		t.ids = append(t.ids, item.TaskID())
	}

	t.items[item.TaskID()] = item
}

func (t *Pool[E]) Remove(id E) {
	t.Lock()
	defer t.Unlock()
	delete(t.items, id)
}

func (t *Pool[E]) Get(id E) Task[E] {
	t.RLock()
	defer t.RUnlock()

	return t.items[id]
}

func (t *Pool[E]) BatchGet(ids []E) []Task[E] {
	t.RLock()
	defer t.RUnlock()

	tasks := make([]Task[E], 0)

	for _, id := range ids {
		if t.items[id] != nil {
			tasks = append(tasks, t.items[id])
		}
	}

	return tasks
}

func (t *Pool[E]) IsExist(id E) bool {
	t.RLock()
	defer t.RUnlock()
	_, ok := t.items[id]

	return ok
}

func (t *Pool[E]) Len() int {
	t.RLock()
	defer t.RUnlock()

	return len(t.items)
}

func (t *Pool[E]) IsEmpty() bool {
	t.RLock()
	defer t.RUnlock()

	return len(t.items) == 0
}

func (t *Pool[E]) Println() {
	t.RLock()
	defer t.RUnlock()

	fmt.Println(t.ids)
}

func (t *Pool[E]) Clear() {
	t.Lock()
	defer t.Unlock()
	t.clear()
}

func (t *Pool[E]) clear() {
	t.ids = make([]E, 0)
	t.items = make(map[E]Task[E])
}

func (t *Pool[E]) Reset() {
	t.Clear()
}

func (t *Pool[E]) GetTopN(N int64) []Task[E] {
	t.Lock()

	is := slices.IsSorted(t.ids)
	if !is {
		slices.Sort(t.ids)
	}
	t.Unlock()
	t.RLock()
	defer t.RUnlock()

	keys := make(map[E]Task[E], 0)

	items := make([]Task[E], 0, N)
	for _, id := range t.ids {
		item, ok := t.items[id]
		_, is := keys[id]

		if ok && !is {
			keys[item.TaskID()] = item

			items = append(items, item)
			if len(items) == int(N) {
				break
			}
		}
	}

	return items
}

func (t *Pool[E]) RemoveTopN(levelId E) {
	t.Lock()
	defer t.Unlock()

	if len(t.items) == 0 {
		t.ids = make([]E, 0)
		return
	}

	is := slices.IsSorted(t.ids)
	if !is {
		slices.Sort(t.ids)
	}

	if levelId < t.ids[0] {
		return
	}

	if levelId >= t.ids[len(t.ids)-1] {
		t.clear()
		return
	}

	ids := make([]E, len(t.ids))
	copy(ids, t.ids)

	levelIndex := 0
L:
	for index, id := range ids {
		if id <= levelId {
			delete(t.items, id)
			levelIndex = index
		} else {
			break L
		}
	}

	t.ids = t.ids[levelIndex+1:]
}
