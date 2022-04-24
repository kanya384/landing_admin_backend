// Code generated by go-swagger; DO NOT EDIT.

package hod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeletePhotosIDHandlerFunc turns a function with the right signature into a delete photos ID handler
type DeletePhotosIDHandlerFunc func(DeletePhotosIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeletePhotosIDHandlerFunc) Handle(params DeletePhotosIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeletePhotosIDHandler interface for that can handle valid delete photos ID params
type DeletePhotosIDHandler interface {
	Handle(DeletePhotosIDParams, interface{}) middleware.Responder
}

// NewDeletePhotosID creates a new http.Handler for the delete photos ID operation
func NewDeletePhotosID(ctx *middleware.Context, handler DeletePhotosIDHandler) *DeletePhotosID {
	return &DeletePhotosID{Context: ctx, Handler: handler}
}

/* DeletePhotosID swagger:route DELETE /photos/{id} hod deletePhotosId

delete photo by id

*/
type DeletePhotosID struct {
	Context *middleware.Context
	Handler DeletePhotosIDHandler
}

func (o *DeletePhotosID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeletePhotosIDParams()
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
