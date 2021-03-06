// Code generated by go-swagger; DO NOT EDIT.

package video

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// DeleteVideoIDOKCode is the HTTP code returned for type DeleteVideoIDOK
const DeleteVideoIDOKCode int = 200

/*DeleteVideoIDOK delete video success

swagger:response deleteVideoIdOK
*/
type DeleteVideoIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewDeleteVideoIDOK creates DeleteVideoIDOK with default headers values
func NewDeleteVideoIDOK() *DeleteVideoIDOK {

	return &DeleteVideoIDOK{}
}

// WithPayload adds the payload to the delete video Id o k response
func (o *DeleteVideoIDOK) WithPayload(payload *models.ResultResponse) *DeleteVideoIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete video Id o k response
func (o *DeleteVideoIDOK) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVideoIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteVideoIDBadRequestCode is the HTTP code returned for type DeleteVideoIDBadRequest
const DeleteVideoIDBadRequestCode int = 400

/*DeleteVideoIDBadRequest Bad request

swagger:response deleteVideoIdBadRequest
*/
type DeleteVideoIDBadRequest struct {
}

// NewDeleteVideoIDBadRequest creates DeleteVideoIDBadRequest with default headers values
func NewDeleteVideoIDBadRequest() *DeleteVideoIDBadRequest {

	return &DeleteVideoIDBadRequest{}
}

// WriteResponse to the client
func (o *DeleteVideoIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// DeleteVideoIDInternalServerErrorCode is the HTTP code returned for type DeleteVideoIDInternalServerError
const DeleteVideoIDInternalServerErrorCode int = 500

/*DeleteVideoIDInternalServerError Internal Server Error

swagger:response deleteVideoIdInternalServerError
*/
type DeleteVideoIDInternalServerError struct {
}

// NewDeleteVideoIDInternalServerError creates DeleteVideoIDInternalServerError with default headers values
func NewDeleteVideoIDInternalServerError() *DeleteVideoIDInternalServerError {

	return &DeleteVideoIDInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteVideoIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
