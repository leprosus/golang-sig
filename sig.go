package sig

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitSig(callback func(), sigs ...os.Signal) {
	var (
		sig  = make(chan os.Signal, 1)
		done = make(chan struct{})
	)

	if len(sigs) == 0 {
		sigs = append(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	}

	signal.Notify(sig, sigs...)

	go func() {
		<-sig

		callback()

		done <- struct{}{}
	}()

	<-done
}
