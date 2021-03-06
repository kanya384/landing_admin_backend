// Code generated by go-swagger; DO NOT EDIT.

package project_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetProjectInfoHandlerFunc turns a function with the right signature into a get project info handler
type GetProjectInfoHandlerFunc func(GetProjectInfoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetProjectInfoHandlerFunc) Handle(params GetProjectInfoParams) middleware.Responder {
	return fn(params)
}

// GetProjectInfoHandler interface for that can handle valid get project info params
type GetProjectInfoHandler interface {
	Handle(GetProjectInfoParams) middleware.Responder
}

// NewGetProjectInfo creates a new http.Handler for the get project info operation
func NewGetProjectInfo(ctx *middleware.Context, handler GetProjectInfoHandler) *GetProjectInfo {
	return &GetProjectInfo{Context: ctx, Handler: handler}
}

/* GetProjectInfo swagger:route GET /projectInfo projectInfo getProjectInfo

gets filtered projectInfo list

*/
type GetProjectInfo struct {
	Context *middleware.Context
	Handler GetProjectInfoHandler
}

func (o *GetProjectInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetProjectInfoParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
