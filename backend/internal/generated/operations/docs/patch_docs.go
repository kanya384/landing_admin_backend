// Code generated by go-swagger; DO NOT EDIT.

package docs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchDocsHandlerFunc turns a function with the right signature into a patch docs handler
type PatchDocsHandlerFunc func(PatchDocsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchDocsHandlerFunc) Handle(params PatchDocsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PatchDocsHandler interface for that can handle valid patch docs params
type PatchDocsHandler interface {
	Handle(PatchDocsParams, interface{}) middleware.Responder
}

// NewPatchDocs creates a new http.Handler for the patch docs operation
func NewPatchDocs(ctx *middleware.Context, handler PatchDocsHandler) *PatchDocs {
	return &PatchDocs{Context: ctx, Handler: handler}
}

/* PatchDocs swagger:route PATCH /docs docs patchDocs

update doc

*/
type PatchDocs struct {
	Context *middleware.Context
	Handler PatchDocsHandler
}

func (o *PatchDocs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPatchDocsParams()
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
