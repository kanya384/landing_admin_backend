// Code generated by go-swagger; DO NOT EDIT.

package docs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutDocHandlerFunc turns a function with the right signature into a put doc handler
type PutDocHandlerFunc func(PutDocParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PutDocHandlerFunc) Handle(params PutDocParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PutDocHandler interface for that can handle valid put doc params
type PutDocHandler interface {
	Handle(PutDocParams, interface{}) middleware.Responder
}

// NewPutDoc creates a new http.Handler for the put doc operation
func NewPutDoc(ctx *middleware.Context, handler PutDocHandler) *PutDoc {
	return &PutDoc{Context: ctx, Handler: handler}
}

/* PutDoc swagger:route PUT /doc docs putDoc

create doc file

*/
type PutDoc struct {
	Context *middleware.Context
	Handler PutDocHandler
}

func (o *PutDoc) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutDocParams()
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
