package discard

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/taylorchu/work"
)

func TestMaxRetry(t *testing.T) {
	job := work.NewJob()
	opt := &work.DequeueOptions{
		Namespace:    "n1",
		QueueID:      "q1",
		InvisibleSec: 10,
	}

	retryErr := fmt.Errorf("error")
	d := MaxRetry(4)
	h := d(func(work.ContextMap, *work.Job, *work.DequeueOptions) error {
		return retryErr
	})

	var err error
	for i := 0; i < 4; i++ {
		err := h(work.ContextMap{}, job, opt)
		require.Error(t, err)
		require.Equal(t, retryErr, err)

		job.Retries++
	}

	err = h(work.ContextMap{}, job, opt)
	require.Error(t, err)
	require.Equal(t, work.ErrUnrecoverable, err)
}
