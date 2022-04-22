// Code generated by go-swagger; DO NOT EDIT.

package posters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostPostersMaxParseMemory sets the maximum size in bytes for
// the multipart form parser for this operation.
//
// The default value is 32 MB.
// The multipart parser stores up to this + 10MB.
var PostPostersMaxParseMemory int64 = 32 << 20

// NewPostPostersParams creates a new PostPostersParams object
//
// There are no default values defined in the spec.
func NewPostPostersParams() PostPostersParams {

	return PostPostersParams{}
}

// PostPostersParams contains all the bound params for the post posters operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostPosters
type PostPostersParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: formData
	*/
	Description string
	/*The file to upload
	  In: formData
	*/
	File io.ReadCloser
	/*
	  Required: true
	  In: formData
	*/
	ID string
	/*
	  Required: true
	  In: formData
	*/
	Order int64
	/*
	  Required: true
	  In: formData
	*/
	Title string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostPostersParams() beforehand.
func (o *PostPostersParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(PostPostersMaxParseMemory); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}
	fds := runtime.Values(r.Form)

	fdDescription, fdhkDescription, _ := fds.GetOK("description")
	if err := o.bindDescription(fdDescription, fdhkDescription, route.Formats); err != nil {
		res = append(res, err)
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil && err != http.ErrMissingFile {
		res = append(res, errors.New(400, "reading file %q failed: %v", "file", err))
	} else if err == http.ErrMissingFile {
		// no-op for missing but optional file parameter
	} else if err := o.bindFile(file, fileHeader); err != nil {
		res = append(res, err)
	} else {
		o.File = &runtime.File{Data: file, Header: fileHeader}
	}

	fdID, fdhkID, _ := fds.GetOK("id")
	if err := o.bindID(fdID, fdhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	fdOrder, fdhkOrder, _ := fds.GetOK("order")
	if err := o.bindOrder(fdOrder, fdhkOrder, route.Formats); err != nil {
		res = append(res, err)
	}

	fdTitle, fdhkTitle, _ := fds.GetOK("title")
	if err := o.bindTitle(fdTitle, fdhkTitle, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDescription binds and validates parameter Description from formData.
func (o *PostPostersParams) bindDescription(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("description", "formData", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("description", "formData", raw); err != nil {
		return err
	}
	o.Description = raw

	return nil
}

// bindFile binds file parameter File.
//
// The only supported validations on files are MinLength and MaxLength
func (o *PostPostersParams) bindFile(file multipart.File, header *multipart.FileHeader) error {
	return nil
}

// bindID binds and validates parameter ID from formData.
func (o *PostPostersParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("id", "formData", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("id", "formData", raw); err != nil {
		return err
	}
	o.ID = raw

	return nil
}

// bindOrder binds and validates parameter Order from formData.
func (o *PostPostersParams) bindOrder(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("order", "formData", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("order", "formData", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("order", "formData", "int64", raw)
	}
	o.Order = value

	return nil
}

// bindTitle binds and validates parameter Title from formData.
func (o *PostPostersParams) bindTitle(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("title", "formData", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("title", "formData", raw); err != nil {
		return err
	}
	o.Title = raw

	return nil
}
