package logrus

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/taylorchu/work"
)

func TestHandleFuncLogger(t *testing.T) {
	job := work.NewJob()
	opt := &work.DequeueOptions{
		Namespace: "n1",
		QueueID:   "q1",
	}
	h := HandleFuncLogger(func(work.ContextMap, *work.Job, *work.DequeueOptions) error {
		return nil
	})

	c := work.ContextMap{}
	err := h(c, job, opt)
	require.NoError(t, err)

	h = HandleFuncLogger(func(work.ContextMap, *work.Job, *work.DequeueOptions) error {
		return errors.New("no reason")
	})
	err = h(c, job, opt)
	require.Error(t, err)
}

func TestEnqueueFuncLogger(t *testing.T) {
	job := work.NewJob()
	opt := &work.EnqueueOptions{
		Namespace: "n1",
		QueueID:   "q1",
	}
	h := EnqueueFuncLogger(func(*work.Job, *work.EnqueueOptions) error {
		return nil
	})

	err := h(job, opt)
	require.NoError(t, err)

	h = EnqueueFuncLogger(func(*work.Job, *work.EnqueueOptions) error {
		return errors.New("no reason")
	})
	err = h(job, opt)
	require.Error(t, err)
}
