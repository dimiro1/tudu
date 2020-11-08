package graceful

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// ListenAndServe starts the HTTP server and blocks until it is shutdown.
func ListenAndServe(server *http.Server, timeout time.Duration) error {
	srvExitedCtx, srvExitedCancel := context.WithCancel(context.Background())

	go func() {
		defer srvExitedCancel()

		if err := server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			}
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGUSR1)
	defer signal.Stop(signalChan)

	select {
	case <-signalChan:
	case <-srvExitedCtx.Done():
		return nil
	}

	timeoutCtx, cancelTimeout := context.WithTimeout(context.Background(), timeout)
	defer cancelTimeout()

	return server.Shutdown(timeoutCtx)
}
