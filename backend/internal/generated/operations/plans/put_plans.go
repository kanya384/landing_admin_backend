// Code generated by go-swagger; DO NOT EDIT.

package plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutPlansHandlerFunc turns a function with the right signature into a put plans handler
type PutPlansHandlerFunc func(PutPlansParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutPlansHandlerFunc) Handle(params PutPlansParams) middleware.Responder {
	return fn(params)
}

// PutPlansHandler interface for that can handle valid put plans params
type PutPlansHandler interface {
	Handle(PutPlansParams) middleware.Responder
}

// NewPutPlans creates a new http.Handler for the put plans operation
func NewPutPlans(ctx *middleware.Context, handler PutPlansHandler) *PutPlans {
	return &PutPlans{Context: ctx, Handler: handler}
}

/* PutPlans swagger:route PUT /plans plans putPlans

create/update plans from file

*/
type PutPlans struct {
	Context *middleware.Context
	Handler PutPlansHandler
}

func (o *PutPlans) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutPlansParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
