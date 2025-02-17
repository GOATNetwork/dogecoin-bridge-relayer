package pool

import (
	"testing"

	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/stretchr/testify/assert"
)

type Case struct {
	ID   uint64
	data string
}

func (c *Case) Type() int {
	// TODO implement me
	panic("implement me")
}

func (c *Case) TaskID() uint64 {
	return c.ID
}

func TestPool(t *testing.T) {
	p := NewTaskPool[uint64]()
	assert.True(t, p.IsEmpty())

	cases := []Case{
		{
			ID:   30,
			data: "30",
		},
		{
			ID:   3,
			data: "3",
		},
		{
			ID:   1,
			data: "1",
		},
		{
			ID:   4,
			data: "4",
		},
		{
			ID:   2,
			data: "2",
		},
		{
			ID:   10,
			data: "10",
		},
		{
			ID:   7,
			data: "7",
		},
	}

	for _, c := range cases {
		p.Add(&c)
	}

	tasks := p.GetTopN(4)
	t.Log(utils.FormatJSON(tasks))
	assert.Equal(t, len(tasks), 4)

	tasks = p.GetTopN(10)
	assert.Equal(t, len(tasks), len(cases))
	t.Log(utils.FormatJSON(tasks))

	task := p.First()
	assert.Equal(t, task.TaskID(), uint64(1))

	task = p.Last()
	assert.Equal(t, task.TaskID(), uint64(30))

	assert.True(t, !p.IsEmpty())
	assert.True(t, p.IsExist(task.TaskID()))
	assert.True(t, !p.IsExist(uint64(55)))

	assert.True(t, p.IsExist(uint64(30)))
	p.Remove(uint64(30))
	assert.True(t, !p.IsExist(uint64(30)))

	p.RemoveTopN(5)
	p.Println()
	assert.True(t, !p.IsEmpty())
	assert.Equal(t, 2, p.Len())
	p.RemoveTopN(100)
	p.Println()
	assert.True(t, p.IsEmpty())
}
