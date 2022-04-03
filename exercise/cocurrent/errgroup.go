package cocurrent

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func ServerStart() {
	eg, _ := errgroup.WithContext(context.Background())
	eg.Go(
		func() error {
			return http.ListenAndServe(":8080", nil)
		})

	eg.Go(func() error {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGTERM)
		<-c
		return errors.New("receive sig term")
	})
	err := eg.Wait()
	if err != nil {
		log.Fatal("http start error", err.Error())
		os.Exit(1)
	}
}
