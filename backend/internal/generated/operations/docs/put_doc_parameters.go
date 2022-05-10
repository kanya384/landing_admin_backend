// Code generated by go-swagger; DO NOT EDIT.

package docs

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
	"github.com/go-openapi/validate"
)

// PutDocMaxParseMemory sets the maximum size in bytes for
// the multipart form parser for this operation.
//
// The default value is 32 MB.
// The multipart parser stores up to this + 10MB.
var PutDocMaxParseMemory int64 = 32 << 20

// NewPutDocParams creates a new PutDocParams object
//
// There are no default values defined in the spec.
func NewPutDocParams() PutDocParams {

	return PutDocParams{}
}

// PutDocParams contains all the bound params for the put doc operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutDoc
type PutDocParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The file to upload
	  Required: true
	  In: formData
	*/
	File io.ReadCloser
	/*docs title
	  Required: true
	  In: formData
	*/
	Title string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutDocParams() beforehand.
func (o *PutDocParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(PutDocMaxParseMemory); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}
	fds := runtime.Values(r.Form)

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "file", err))
	} else if err := o.bindFile(file, fileHeader); err != nil {
		// Required: true
		res = append(res, err)
	} else {
		o.File = &runtime.File{Data: file, Header: fileHeader}
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

// bindFile binds file parameter File.
//
// The only supported validations on files are MinLength and MaxLength
func (o *PutDocParams) bindFile(file multipart.File, header *multipart.FileHeader) error {
	return nil
}

// bindTitle binds and validates parameter Title from formData.
func (o *PutDocParams) bindTitle(rawData []string, hasKey bool, formats strfmt.Registry) error {
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