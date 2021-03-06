// Code generated by go-swagger; DO NOT EDIT.

package project_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostProjectInfoOrdersHandlerFunc turns a function with the right signature into a post project info orders handler
type PostProjectInfoOrdersHandlerFunc func(PostProjectInfoOrdersParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostProjectInfoOrdersHandlerFunc) Handle(params PostProjectInfoOrdersParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostProjectInfoOrdersHandler interface for that can handle valid post project info orders params
type PostProjectInfoOrdersHandler interface {
	Handle(PostProjectInfoOrdersParams, interface{}) middleware.Responder
}

// NewPostProjectInfoOrders creates a new http.Handler for the post project info orders operation
func NewPostProjectInfoOrders(ctx *middleware.Context, handler PostProjectInfoOrdersHandler) *PostProjectInfoOrders {
	return &PostProjectInfoOrders{Context: ctx, Handler: handler}
}

/* PostProjectInfoOrders swagger:route POST /projectInfo/orders projectInfo postProjectInfoOrders

updates projectInfo oders

*/
type PostProjectInfoOrders struct {
	Context *middleware.Context
	Handler PostProjectInfoOrdersHandler
}

func (o *PostProjectInfoOrders) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostProjectInfoOrdersParams()
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
