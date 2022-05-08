// Code generated by go-swagger; DO NOT EDIT.

package hod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"landing_admin_backend/models"
)

// NewPatchYearsParams creates a new PatchYearsParams object
//
// There are no default values defined in the spec.
func NewPatchYearsParams() PatchYearsParams {

	return PatchYearsParams{}
}

// PatchYearsParams contains all the bound params for the patch years operation
// typically these are obtained from a http.Request
//
// swagger:parameters PatchYears
type PatchYearsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Params *models.Year
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPatchYearsParams() beforehand.
func (o *PatchYearsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Year
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("params", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Params = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}