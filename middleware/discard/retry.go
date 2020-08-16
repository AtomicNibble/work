package discard

import (
	"github.com/taylorchu/work"
)

// MaxRetry discards a job if reaches max retry
func MaxRetry(maxRetries int64) work.HandleMiddleware {
	return func(f work.HandleFunc) work.HandleFunc {
		return func(c work.ContextMap, job *work.Job, opt *work.DequeueOptions) error {
			err := f(c, job, opt)
			if err != nil {
				if job.Retries >= maxRetries {
					return work.ErrUnrecoverable
				}

				return err
			}
			return nil
		}
	}
}
