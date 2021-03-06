// Code generated by go-swagger; DO NOT EDIT.

package editable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"landing_admin_backend/models"
)

// NewPostEditableParams creates a new PostEditableParams object
//
// There are no default values defined in the spec.
func NewPostEditableParams() PostEditableParams {

	return PostEditableParams{}
}

// PostEditableParams contains all the bound params for the post editable operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostEditable
type PostEditableParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Params []*models.Editable
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostEditableParams() beforehand.
func (o *PostEditableParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body []*models.Editable
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("params", "body", "", err))
		} else {

			// validate array of body objects
			for i := range body {
				if body[i] == nil {
					continue
				}
				if err := body[i].Validate(route.Formats); err != nil {
					res = append(res, err)
					break
				}
			}

			if len(res) == 0 {
				o.Params = body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
