// Code generated by go-swagger; DO NOT EDIT.

package leads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// PutLeadOKCode is the HTTP code returned for type PutLeadOK
const PutLeadOKCode int = 200

/*PutLeadOK Request success

swagger:response putLeadOK
*/
type PutLeadOK struct {

	/*
	  In: Body
	*/
	Payload *models.Lead `json:"body,omitempty"`
}

// NewPutLeadOK creates PutLeadOK with default headers values
func NewPutLeadOK() *PutLeadOK {

	return &PutLeadOK{}
}

// WithPayload adds the payload to the put lead o k response
func (o *PutLeadOK) WithPayload(payload *models.Lead) *PutLeadOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put lead o k response
func (o *PutLeadOK) SetPayload(payload *models.Lead) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutLeadOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutLeadBadRequestCode is the HTTP code returned for type PutLeadBadRequest
const PutLeadBadRequestCode int = 400

/*PutLeadBadRequest Bad request

swagger:response putLeadBadRequest
*/
type PutLeadBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewPutLeadBadRequest creates PutLeadBadRequest with default headers values
func NewPutLeadBadRequest() *PutLeadBadRequest {

	return &PutLeadBadRequest{}
}

// WithPayload adds the payload to the put lead bad request response
func (o *PutLeadBadRequest) WithPayload(payload *models.ResultResponse) *PutLeadBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put lead bad request response
func (o *PutLeadBadRequest) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutLeadBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutLeadInternalServerErrorCode is the HTTP code returned for type PutLeadInternalServerError
const PutLeadInternalServerErrorCode int = 500

/*PutLeadInternalServerError Internal Server Error

swagger:response putLeadInternalServerError
*/
type PutLeadInternalServerError struct {
}

// NewPutLeadInternalServerError creates PutLeadInternalServerError with default headers values
func NewPutLeadInternalServerError() *PutLeadInternalServerError {

	return &PutLeadInternalServerError{}
}

// WriteResponse to the client
func (o *PutLeadInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
