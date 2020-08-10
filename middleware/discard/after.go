package discard

import (
	"time"

	"github.com/taylorchu/work"
)

// After discards a job if it is already stale.
func After(d time.Duration) work.HandleMiddleware {
	return func(f work.HandleFunc) work.HandleFunc {
		return func(c work.ContextMap, job *work.Job, opt *work.DequeueOptions) error {
			err := f(c, job, opt)
			if err != nil {
				if time.Now().Sub(job.CreatedAt) > d {
					return work.ErrUnrecoverable
				}
				return err
			}
			return nil
		}
	}
}
