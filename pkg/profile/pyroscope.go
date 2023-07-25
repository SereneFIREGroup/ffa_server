package profile

import (
	"github.com/pyroscope-io/client/pyroscope"
	"github.com/serenefiregroup/ffa_server/pkg/config"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
)

// Profile starts pyroscope to profile the application
func Profile() error {
	addr := config.String("pyroscope", "")
	if addr == "" {
		return nil
	}
	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: "ffa",
		ServerAddress:   "http://localhost:4040",
		Logger:          pyroscope.StandardLogger,
	})
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}
