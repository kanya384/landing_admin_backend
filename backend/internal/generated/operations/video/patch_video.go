// Code generated by go-swagger; DO NOT EDIT.

package video

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchVideoHandlerFunc turns a function with the right signature into a patch video handler
type PatchVideoHandlerFunc func(PatchVideoParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchVideoHandlerFunc) Handle(params PatchVideoParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PatchVideoHandler interface for that can handle valid patch video params
type PatchVideoHandler interface {
	Handle(PatchVideoParams, interface{}) middleware.Responder
}

// NewPatchVideo creates a new http.Handler for the patch video operation
func NewPatchVideo(ctx *middleware.Context, handler PatchVideoHandler) *PatchVideo {
	return &PatchVideo{Context: ctx, Handler: handler}
}

/* PatchVideo swagger:route PATCH /video video patchVideo

update video file

*/
type PatchVideo struct {
	Context *middleware.Context
	Handler PatchVideoHandler
}

func (o *PatchVideo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPatchVideoParams()
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
