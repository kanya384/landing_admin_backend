// Code generated by go-swagger; DO NOT EDIT.

package docs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// PatchDocsOKCode is the HTTP code returned for type PatchDocsOK
const PatchDocsOKCode int = 200

/*PatchDocsOK Request success

swagger:response patchDocsOK
*/
type PatchDocsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Doc `json:"body,omitempty"`
}

// NewPatchDocsOK creates PatchDocsOK with default headers values
func NewPatchDocsOK() *PatchDocsOK {

	return &PatchDocsOK{}
}

// WithPayload adds the payload to the patch docs o k response
func (o *PatchDocsOK) WithPayload(payload *models.Doc) *PatchDocsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch docs o k response
func (o *PatchDocsOK) SetPayload(payload *models.Doc) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchDocsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchDocsBadRequestCode is the HTTP code returned for type PatchDocsBadRequest
const PatchDocsBadRequestCode int = 400

/*PatchDocsBadRequest Bad request

swagger:response patchDocsBadRequest
*/
type PatchDocsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewPatchDocsBadRequest creates PatchDocsBadRequest with default headers values
func NewPatchDocsBadRequest() *PatchDocsBadRequest {

	return &PatchDocsBadRequest{}
}

// WithPayload adds the payload to the patch docs bad request response
func (o *PatchDocsBadRequest) WithPayload(payload *models.ResultResponse) *PatchDocsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch docs bad request response
func (o *PatchDocsBadRequest) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchDocsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchDocsInternalServerErrorCode is the HTTP code returned for type PatchDocsInternalServerError
const PatchDocsInternalServerErrorCode int = 500

/*PatchDocsInternalServerError Internal Server Error

swagger:response patchDocsInternalServerError
*/
type PatchDocsInternalServerError struct {
}

// NewPatchDocsInternalServerError creates PatchDocsInternalServerError with default headers values
func NewPatchDocsInternalServerError() *PatchDocsInternalServerError {

	return &PatchDocsInternalServerError{}
}

// WriteResponse to the client
func (o *PatchDocsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}