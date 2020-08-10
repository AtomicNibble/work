package recovery

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/taylorchu/work"
)

func TestCatchPanic(t *testing.T) {
	job := work.NewJob()
	opt := &work.DequeueOptions{
		Namespace: "n1",
		QueueID:   "q1",
	}
	h := CatchPanic(func(work.ContextMap, *work.Job, *work.DequeueOptions) error {
		panic("fatal error")
	})

	c := work.ContextMap{}
	err := h(c, job, opt)
	require.Error(t, err)
}
