package conc_utils

import (
	"github.com/serenefiregroup/ffa_server/pkg/log"
	"github.com/sourcegraph/conc"
)

// GoSafe runs the given fn using another goroutine, recovers if fn panics.
func GoSafe(fn func()) {
	wg := conc.NewWaitGroup()
	defer func() {
		recoverPanic := wg.WaitAndRecover()
		if recoverPanic != nil {
			log.Error(recoverPanic.String())
		}
	}()

	wg.Go(fn)
}
