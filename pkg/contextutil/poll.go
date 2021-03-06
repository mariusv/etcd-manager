package contextutil

import (
	"context"
	"time"

	"k8s.io/klog"
)

func Sleep(ctx context.Context, duration time.Duration) error {
	timer := time.NewTimer(duration)

	select {
	case <-ctx.Done():
		return ctx.Err()

	case <-timer.C:
		return nil
	}
}

func Forever(ctx context.Context, interval time.Duration, f func()) {
	for {
		if ctx.Err() != nil {
			klog.Infof("context cancelled; exiting loop")
			return
		}

		f()

		Sleep(ctx, interval)
	}

}
