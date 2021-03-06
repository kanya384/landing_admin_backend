// Code generated by go-swagger; DO NOT EDIT.

package docs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// PatchDocOKCode is the HTTP code returned for type PatchDocOK
const PatchDocOKCode int = 200

/*PatchDocOK Request success

swagger:response patchDocOK
*/
type PatchDocOK struct {

	/*
	  In: Body
	*/
	Payload *models.Doc `json:"body,omitempty"`
}

// NewPatchDocOK creates PatchDocOK with default headers values
func NewPatchDocOK() *PatchDocOK {

	return &PatchDocOK{}
}

// WithPayload adds the payload to the patch doc o k response
func (o *PatchDocOK) WithPayload(payload *models.Doc) *PatchDocOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch doc o k response
func (o *PatchDocOK) SetPayload(payload *models.Doc) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchDocOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchDocBadRequestCode is the HTTP code returned for type PatchDocBadRequest
const PatchDocBadRequestCode int = 400

/*PatchDocBadRequest Bad request

swagger:response patchDocBadRequest
*/
type PatchDocBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewPatchDocBadRequest creates PatchDocBadRequest with default headers values
func NewPatchDocBadRequest() *PatchDocBadRequest {

	return &PatchDocBadRequest{}
}

// WithPayload adds the payload to the patch doc bad request response
func (o *PatchDocBadRequest) WithPayload(payload *models.ResultResponse) *PatchDocBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch doc bad request response
func (o *PatchDocBadRequest) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchDocBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchDocInternalServerErrorCode is the HTTP code returned for type PatchDocInternalServerError
const PatchDocInternalServerErrorCode int = 500

/*PatchDocInternalServerError Internal Server Error

swagger:response patchDocInternalServerError
*/
type PatchDocInternalServerError struct {
}

// NewPatchDocInternalServerError creates PatchDocInternalServerError with default headers values
func NewPatchDocInternalServerError() *PatchDocInternalServerError {

	return &PatchDocInternalServerError{}
}

// WriteResponse to the client
func (o *PatchDocInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
