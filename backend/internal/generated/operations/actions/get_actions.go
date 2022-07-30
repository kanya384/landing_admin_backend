// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetActionsHandlerFunc turns a function with the right signature into a get actions handler
type GetActionsHandlerFunc func(GetActionsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetActionsHandlerFunc) Handle(params GetActionsParams) middleware.Responder {
	return fn(params)
}

// GetActionsHandler interface for that can handle valid get actions params
type GetActionsHandler interface {
	Handle(GetActionsParams) middleware.Responder
}

// NewGetActions creates a new http.Handler for the get actions operation
func NewGetActions(ctx *middleware.Context, handler GetActionsHandler) *GetActions {
	return &GetActions{Context: ctx, Handler: handler}
}

/* GetActions swagger:route GET /actions actions getActions

gets filtered actions list

*/
type GetActions struct {
	Context *middleware.Context
	Handler GetActionsHandler
}

func (o *GetActions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetActionsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
