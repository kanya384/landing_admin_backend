// Code generated by go-swagger; DO NOT EDIT.

package hod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// DeleteYearsIDOKCode is the HTTP code returned for type DeleteYearsIDOK
const DeleteYearsIDOKCode int = 200

/*DeleteYearsIDOK year delete success

swagger:response deleteYearsIdOK
*/
type DeleteYearsIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewDeleteYearsIDOK creates DeleteYearsIDOK with default headers values
func NewDeleteYearsIDOK() *DeleteYearsIDOK {

	return &DeleteYearsIDOK{}
}

// WithPayload adds the payload to the delete years Id o k response
func (o *DeleteYearsIDOK) WithPayload(payload *models.ResultResponse) *DeleteYearsIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete years Id o k response
func (o *DeleteYearsIDOK) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteYearsIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteYearsIDBadRequestCode is the HTTP code returned for type DeleteYearsIDBadRequest
const DeleteYearsIDBadRequestCode int = 400

/*DeleteYearsIDBadRequest Bad request

swagger:response deleteYearsIdBadRequest
*/
type DeleteYearsIDBadRequest struct {
}

// NewDeleteYearsIDBadRequest creates DeleteYearsIDBadRequest with default headers values
func NewDeleteYearsIDBadRequest() *DeleteYearsIDBadRequest {

	return &DeleteYearsIDBadRequest{}
}

// WriteResponse to the client
func (o *DeleteYearsIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// DeleteYearsIDInternalServerErrorCode is the HTTP code returned for type DeleteYearsIDInternalServerError
const DeleteYearsIDInternalServerErrorCode int = 500

/*DeleteYearsIDInternalServerError Internal Server Error

swagger:response deleteYearsIdInternalServerError
*/
type DeleteYearsIDInternalServerError struct {
}

// NewDeleteYearsIDInternalServerError creates DeleteYearsIDInternalServerError with default headers values
func NewDeleteYearsIDInternalServerError() *DeleteYearsIDInternalServerError {

	return &DeleteYearsIDInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteYearsIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
