// Code generated by go-swagger; DO NOT EDIT.

package hod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// PatchMonthsOKCode is the HTTP code returned for type PatchMonthsOK
const PatchMonthsOKCode int = 200

/*PatchMonthsOK Request success

swagger:response patchMonthsOK
*/
type PatchMonthsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Month `json:"body,omitempty"`
}

// NewPatchMonthsOK creates PatchMonthsOK with default headers values
func NewPatchMonthsOK() *PatchMonthsOK {

	return &PatchMonthsOK{}
}

// WithPayload adds the payload to the patch months o k response
func (o *PatchMonthsOK) WithPayload(payload *models.Month) *PatchMonthsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch months o k response
func (o *PatchMonthsOK) SetPayload(payload *models.Month) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchMonthsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchMonthsBadRequestCode is the HTTP code returned for type PatchMonthsBadRequest
const PatchMonthsBadRequestCode int = 400

/*PatchMonthsBadRequest Bad request

swagger:response patchMonthsBadRequest
*/
type PatchMonthsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewPatchMonthsBadRequest creates PatchMonthsBadRequest with default headers values
func NewPatchMonthsBadRequest() *PatchMonthsBadRequest {

	return &PatchMonthsBadRequest{}
}

// WithPayload adds the payload to the patch months bad request response
func (o *PatchMonthsBadRequest) WithPayload(payload *models.ResultResponse) *PatchMonthsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch months bad request response
func (o *PatchMonthsBadRequest) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchMonthsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchMonthsInternalServerErrorCode is the HTTP code returned for type PatchMonthsInternalServerError
const PatchMonthsInternalServerErrorCode int = 500

/*PatchMonthsInternalServerError Internal Server Error

swagger:response patchMonthsInternalServerError
*/
type PatchMonthsInternalServerError struct {
}

// NewPatchMonthsInternalServerError creates PatchMonthsInternalServerError with default headers values
func NewPatchMonthsInternalServerError() *PatchMonthsInternalServerError {

	return &PatchMonthsInternalServerError{}
}

// WriteResponse to the client
func (o *PatchMonthsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
