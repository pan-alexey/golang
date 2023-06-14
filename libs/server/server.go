package server

import (
	"net/http"
	"strings"
	"veva/libs/connect"
)

func Serve(srv *http.Server) (string, error) {
	cErr := make(chan error)
	cOk := make(chan bool)

	addr := strings.Split(srv.Addr, ":")

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			cErr <- err
		}
	}()

	go func() {
		var _, err = connect.CheckTCP(addr[0], addr[1])
		if err != nil {
			cErr <- err
		} else {
			cOk <- true
		}
	}()

	select {
	case err := <-cErr:
		close(cErr)
		close(cOk)
		return srv.Addr, err
	case <-cOk:
		close(cErr)
		close(cOk)
		return srv.Addr, nil
	}
}
