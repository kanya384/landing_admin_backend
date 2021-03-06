// Code generated by go-swagger; DO NOT EDIT.

package titles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// DeleteTitlesIDOKCode is the HTTP code returned for type DeleteTitlesIDOK
const DeleteTitlesIDOKCode int = 200

/*DeleteTitlesIDOK delete titles success

swagger:response deleteTitlesIdOK
*/
type DeleteTitlesIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewDeleteTitlesIDOK creates DeleteTitlesIDOK with default headers values
func NewDeleteTitlesIDOK() *DeleteTitlesIDOK {

	return &DeleteTitlesIDOK{}
}

// WithPayload adds the payload to the delete titles Id o k response
func (o *DeleteTitlesIDOK) WithPayload(payload *models.ResultResponse) *DeleteTitlesIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete titles Id o k response
func (o *DeleteTitlesIDOK) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTitlesIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteTitlesIDBadRequestCode is the HTTP code returned for type DeleteTitlesIDBadRequest
const DeleteTitlesIDBadRequestCode int = 400

/*DeleteTitlesIDBadRequest Bad request

swagger:response deleteTitlesIdBadRequest
*/
type DeleteTitlesIDBadRequest struct {
}

// NewDeleteTitlesIDBadRequest creates DeleteTitlesIDBadRequest with default headers values
func NewDeleteTitlesIDBadRequest() *DeleteTitlesIDBadRequest {

	return &DeleteTitlesIDBadRequest{}
}

// WriteResponse to the client
func (o *DeleteTitlesIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// DeleteTitlesIDInternalServerErrorCode is the HTTP code returned for type DeleteTitlesIDInternalServerError
const DeleteTitlesIDInternalServerErrorCode int = 500

/*DeleteTitlesIDInternalServerError Internal Server Error

swagger:response deleteTitlesIdInternalServerError
*/
type DeleteTitlesIDInternalServerError struct {
}

// NewDeleteTitlesIDInternalServerError creates DeleteTitlesIDInternalServerError with default headers values
func NewDeleteTitlesIDInternalServerError() *DeleteTitlesIDInternalServerError {

	return &DeleteTitlesIDInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteTitlesIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
