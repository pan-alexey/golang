package os

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func WaitExit() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println()
		done <- true
	}()

	fmt.Println("Wait SIGINT")
	<-done
}
