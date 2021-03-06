// Code generated by go-swagger; DO NOT EDIT.

package docs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// DeleteDocIDOKCode is the HTTP code returned for type DeleteDocIDOK
const DeleteDocIDOKCode int = 200

/*DeleteDocIDOK delete doc success

swagger:response deleteDocIdOK
*/
type DeleteDocIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewDeleteDocIDOK creates DeleteDocIDOK with default headers values
func NewDeleteDocIDOK() *DeleteDocIDOK {

	return &DeleteDocIDOK{}
}

// WithPayload adds the payload to the delete doc Id o k response
func (o *DeleteDocIDOK) WithPayload(payload *models.ResultResponse) *DeleteDocIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete doc Id o k response
func (o *DeleteDocIDOK) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDocIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteDocIDBadRequestCode is the HTTP code returned for type DeleteDocIDBadRequest
const DeleteDocIDBadRequestCode int = 400

/*DeleteDocIDBadRequest Bad request

swagger:response deleteDocIdBadRequest
*/
type DeleteDocIDBadRequest struct {
}

// NewDeleteDocIDBadRequest creates DeleteDocIDBadRequest with default headers values
func NewDeleteDocIDBadRequest() *DeleteDocIDBadRequest {

	return &DeleteDocIDBadRequest{}
}

// WriteResponse to the client
func (o *DeleteDocIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// DeleteDocIDInternalServerErrorCode is the HTTP code returned for type DeleteDocIDInternalServerError
const DeleteDocIDInternalServerErrorCode int = 500

/*DeleteDocIDInternalServerError Internal Server Error

swagger:response deleteDocIdInternalServerError
*/
type DeleteDocIDInternalServerError struct {
}

// NewDeleteDocIDInternalServerError creates DeleteDocIDInternalServerError with default headers values
func NewDeleteDocIDInternalServerError() *DeleteDocIDInternalServerError {

	return &DeleteDocIDInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteDocIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
