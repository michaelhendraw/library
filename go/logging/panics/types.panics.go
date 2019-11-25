package panics

import (
	"context"
	"net/http"

	"github.com/michaelhendraw/library/go/router"
	nsq "github.com/nsqio/go-nsq"
)

// A Panics interface provides logging for panic situation
type Panics interface {
	HTTPHandler(http.Handler) http.Handler
	RouterHandler(context.Context, *http.Request, *router.HandlerParam, func(context.Context, *router.HandlerParam) router.HandlerResult) router.HandlerResult
	Cron(func()) func()
	NSQ(nsq.HandlerFunc) nsq.HandlerFunc
	/* Capture is to be used in common function which does not have entry point from any of the above.
	One common usage of Capture() is from go routine function.
	If you run any go routine, panic will not be captured in any of above wrapper.
	So you need to use this Capture() function to log it.
	Usage :
	go func(){
		defer p.Capture() // This function must be deferred otherwise will not work
		// Your code here...
	}
	*/
	Capture()
}
