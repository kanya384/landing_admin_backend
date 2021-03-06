// Code generated by go-swagger; DO NOT EDIT.

package hod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutYearsHandlerFunc turns a function with the right signature into a put years handler
type PutYearsHandlerFunc func(PutYearsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PutYearsHandlerFunc) Handle(params PutYearsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PutYearsHandler interface for that can handle valid put years params
type PutYearsHandler interface {
	Handle(PutYearsParams, interface{}) middleware.Responder
}

// NewPutYears creates a new http.Handler for the put years operation
func NewPutYears(ctx *middleware.Context, handler PutYearsHandler) *PutYears {
	return &PutYears{Context: ctx, Handler: handler}
}

/* PutYears swagger:route PUT /years hod putYears

create year

*/
type PutYears struct {
	Context *middleware.Context
	Handler PutYearsHandler
}

func (o *PutYears) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutYearsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
