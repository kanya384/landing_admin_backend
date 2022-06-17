// Code generated by go-swagger; DO NOT EDIT.

package project_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetProjectInfoPojectInfoIDHandlerFunc turns a function with the right signature into a get project info poject info ID handler
type GetProjectInfoPojectInfoIDHandlerFunc func(GetProjectInfoPojectInfoIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetProjectInfoPojectInfoIDHandlerFunc) Handle(params GetProjectInfoPojectInfoIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetProjectInfoPojectInfoIDHandler interface for that can handle valid get project info poject info ID params
type GetProjectInfoPojectInfoIDHandler interface {
	Handle(GetProjectInfoPojectInfoIDParams, interface{}) middleware.Responder
}

// NewGetProjectInfoPojectInfoID creates a new http.Handler for the get project info poject info ID operation
func NewGetProjectInfoPojectInfoID(ctx *middleware.Context, handler GetProjectInfoPojectInfoIDHandler) *GetProjectInfoPojectInfoID {
	return &GetProjectInfoPojectInfoID{Context: ctx, Handler: handler}
}

/* GetProjectInfoPojectInfoID swagger:route GET /projectInfo/{pojectInfoID} projectInfo getProjectInfoPojectInfoId

get pojectInfo by id

*/
type GetProjectInfoPojectInfoID struct {
	Context *middleware.Context
	Handler GetProjectInfoPojectInfoIDHandler
}

func (o *GetProjectInfoPojectInfoID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetProjectInfoPojectInfoIDParams()
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
