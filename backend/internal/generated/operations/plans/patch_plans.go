// Code generated by go-swagger; DO NOT EDIT.

package plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchPlansHandlerFunc turns a function with the right signature into a patch plans handler
type PatchPlansHandlerFunc func(PatchPlansParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchPlansHandlerFunc) Handle(params PatchPlansParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PatchPlansHandler interface for that can handle valid patch plans params
type PatchPlansHandler interface {
	Handle(PatchPlansParams, interface{}) middleware.Responder
}

// NewPatchPlans creates a new http.Handler for the patch plans operation
func NewPatchPlans(ctx *middleware.Context, handler PatchPlansHandler) *PatchPlans {
	return &PatchPlans{Context: ctx, Handler: handler}
}

/* PatchPlans swagger:route PATCH /plans plans patchPlans

update plan

*/
type PatchPlans struct {
	Context *middleware.Context
	Handler PatchPlansHandler
}

func (o *PatchPlans) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPatchPlansParams()
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
