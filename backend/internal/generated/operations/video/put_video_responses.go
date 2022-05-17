// Code generated by go-swagger; DO NOT EDIT.

package video

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// PutVideoOKCode is the HTTP code returned for type PutVideoOK
const PutVideoOKCode int = 200

/*PutVideoOK Request success

swagger:response putVideoOK
*/
type PutVideoOK struct {

	/*
	  In: Body
	*/
	Payload *models.Video `json:"body,omitempty"`
}

// NewPutVideoOK creates PutVideoOK with default headers values
func NewPutVideoOK() *PutVideoOK {

	return &PutVideoOK{}
}

// WithPayload adds the payload to the put video o k response
func (o *PutVideoOK) WithPayload(payload *models.Video) *PutVideoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put video o k response
func (o *PutVideoOK) SetPayload(payload *models.Video) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutVideoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutVideoBadRequestCode is the HTTP code returned for type PutVideoBadRequest
const PutVideoBadRequestCode int = 400

/*PutVideoBadRequest Bad request

swagger:response putVideoBadRequest
*/
type PutVideoBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewPutVideoBadRequest creates PutVideoBadRequest with default headers values
func NewPutVideoBadRequest() *PutVideoBadRequest {

	return &PutVideoBadRequest{}
}

// WithPayload adds the payload to the put video bad request response
func (o *PutVideoBadRequest) WithPayload(payload *models.ResultResponse) *PutVideoBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put video bad request response
func (o *PutVideoBadRequest) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutVideoBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutVideoInternalServerErrorCode is the HTTP code returned for type PutVideoInternalServerError
const PutVideoInternalServerErrorCode int = 500

/*PutVideoInternalServerError Internal Server Error

swagger:response putVideoInternalServerError
*/
type PutVideoInternalServerError struct {
}

// NewPutVideoInternalServerError creates PutVideoInternalServerError with default headers values
func NewPutVideoInternalServerError() *PutVideoInternalServerError {

	return &PutVideoInternalServerError{}
}

// WriteResponse to the client
func (o *PutVideoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
