// Code generated by go-swagger; DO NOT EDIT.

package hod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetYearsHandlerFunc turns a function with the right signature into a get years handler
type GetYearsHandlerFunc func(GetYearsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetYearsHandlerFunc) Handle(params GetYearsParams) middleware.Responder {
	return fn(params)
}

// GetYearsHandler interface for that can handle valid get years params
type GetYearsHandler interface {
	Handle(GetYearsParams) middleware.Responder
}

// NewGetYears creates a new http.Handler for the get years operation
func NewGetYears(ctx *middleware.Context, handler GetYearsHandler) *GetYears {
	return &GetYears{Context: ctx, Handler: handler}
}

/* GetYears swagger:route GET /years hod getYears

gets years list

*/
type GetYears struct {
	Context *middleware.Context
	Handler GetYearsHandler
}

func (o *GetYears) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetYearsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
