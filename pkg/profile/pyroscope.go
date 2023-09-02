package profile

import (
	"runtime"

	"github.com/grafana/pyroscope-go"
	"github.com/serenefiregroup/ffa_server/pkg/config"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
)

// Profile starts pyroscope to profile the application
func Profile() error {
	addr := config.String("pyroscope", "")
	if addr == "" {
		return nil
	}
	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(5)
	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: "ffa",
		ServerAddress:   "http://localhost:4040",
		Logger:          pyroscope.StandardLogger,
		ProfileTypes: []pyroscope.ProfileType{
			// these profile types are enabled by default:
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,

			// these profile types are optional:
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}
