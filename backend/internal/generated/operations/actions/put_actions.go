// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutActionsHandlerFunc turns a function with the right signature into a put actions handler
type PutActionsHandlerFunc func(PutActionsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PutActionsHandlerFunc) Handle(params PutActionsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PutActionsHandler interface for that can handle valid put actions params
type PutActionsHandler interface {
	Handle(PutActionsParams, interface{}) middleware.Responder
}

// NewPutActions creates a new http.Handler for the put actions operation
func NewPutActions(ctx *middleware.Context, handler PutActionsHandler) *PutActions {
	return &PutActions{Context: ctx, Handler: handler}
}

/* PutActions swagger:route PUT /actions actions putActions

create action

*/
type PutActions struct {
	Context *middleware.Context
	Handler PutActionsHandler
}

func (o *PutActions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutActionsParams()
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
