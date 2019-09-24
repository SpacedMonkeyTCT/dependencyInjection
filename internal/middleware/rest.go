package middleware

import "fmt"

type Rester interface {
	Rest()
}

type RestMiddleware struct {
	r Rester
}

func NewRESTMiddleware(r Rester) RestMiddleware {
	return RestMiddleware{r: r}
}

func (rm RestMiddleware) CallRester() {
	fmt.Println("Calling Rester interface")
	rm.r.Rest()
}
