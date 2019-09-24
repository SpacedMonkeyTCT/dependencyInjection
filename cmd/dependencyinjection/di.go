package main

import (
	"github.com/SpacedMonkeyTCT/dependencyInjection/internal/middleware"
	"github.com/SpacedMonkeyTCT/dependencyInjection/pkg/rest"
)

func main() {
	s := rest.NewSleeper()
	rm := middleware.NewRESTMiddleware(s)
	rm.CallRester()
}
