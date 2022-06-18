// Code generated by go-swagger; DO NOT EDIT.

package project_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewDeleteProjectInfoPojectInfoIDParams creates a new DeleteProjectInfoPojectInfoIDParams object
//
// There are no default values defined in the spec.
func NewDeleteProjectInfoPojectInfoIDParams() DeleteProjectInfoPojectInfoIDParams {

	return DeleteProjectInfoPojectInfoIDParams{}
}

// DeleteProjectInfoPojectInfoIDParams contains all the bound params for the delete project info poject info ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteProjectInfoPojectInfoID
type DeleteProjectInfoPojectInfoIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*String ID of the pojectInfo to get
	  Required: true
	  In: path
	*/
	PojectInfoID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteProjectInfoPojectInfoIDParams() beforehand.
func (o *DeleteProjectInfoPojectInfoIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rPojectInfoID, rhkPojectInfoID, _ := route.Params.GetOK("pojectInfoID")
	if err := o.bindPojectInfoID(rPojectInfoID, rhkPojectInfoID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindPojectInfoID binds and validates parameter PojectInfoID from path.
func (o *DeleteProjectInfoPojectInfoIDParams) bindPojectInfoID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.PojectInfoID = raw

	return nil
}