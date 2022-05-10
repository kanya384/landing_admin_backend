// Code generated by go-swagger; DO NOT EDIT.

package docs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// PutDocOKCode is the HTTP code returned for type PutDocOK
const PutDocOKCode int = 200

/*PutDocOK Request success

swagger:response putDocOK
*/
type PutDocOK struct {

	/*
	  In: Body
	*/
	Payload *models.Doc `json:"body,omitempty"`
}

// NewPutDocOK creates PutDocOK with default headers values
func NewPutDocOK() *PutDocOK {

	return &PutDocOK{}
}

// WithPayload adds the payload to the put doc o k response
func (o *PutDocOK) WithPayload(payload *models.Doc) *PutDocOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put doc o k response
func (o *PutDocOK) SetPayload(payload *models.Doc) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutDocOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutDocBadRequestCode is the HTTP code returned for type PutDocBadRequest
const PutDocBadRequestCode int = 400

/*PutDocBadRequest Bad request

swagger:response putDocBadRequest
*/
type PutDocBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewPutDocBadRequest creates PutDocBadRequest with default headers values
func NewPutDocBadRequest() *PutDocBadRequest {

	return &PutDocBadRequest{}
}

// WithPayload adds the payload to the put doc bad request response
func (o *PutDocBadRequest) WithPayload(payload *models.ResultResponse) *PutDocBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put doc bad request response
func (o *PutDocBadRequest) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutDocBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutDocInternalServerErrorCode is the HTTP code returned for type PutDocInternalServerError
const PutDocInternalServerErrorCode int = 500

/*PutDocInternalServerError Internal Server Error

swagger:response putDocInternalServerError
*/
type PutDocInternalServerError struct {
}

// NewPutDocInternalServerError creates PutDocInternalServerError with default headers values
func NewPutDocInternalServerError() *PutDocInternalServerError {

	return &PutDocInternalServerError{}
}

// WriteResponse to the client
func (o *PutDocInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
