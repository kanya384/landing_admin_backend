// Code generated by go-swagger; DO NOT EDIT.

package advantages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// DeleteAdvantagesIDOKCode is the HTTP code returned for type DeleteAdvantagesIDOK
const DeleteAdvantagesIDOKCode int = 200

/*DeleteAdvantagesIDOK delete advantage success

swagger:response deleteAdvantagesIdOK
*/
type DeleteAdvantagesIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewDeleteAdvantagesIDOK creates DeleteAdvantagesIDOK with default headers values
func NewDeleteAdvantagesIDOK() *DeleteAdvantagesIDOK {

	return &DeleteAdvantagesIDOK{}
}

// WithPayload adds the payload to the delete advantages Id o k response
func (o *DeleteAdvantagesIDOK) WithPayload(payload *models.ResultResponse) *DeleteAdvantagesIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete advantages Id o k response
func (o *DeleteAdvantagesIDOK) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAdvantagesIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAdvantagesIDBadRequestCode is the HTTP code returned for type DeleteAdvantagesIDBadRequest
const DeleteAdvantagesIDBadRequestCode int = 400

/*DeleteAdvantagesIDBadRequest Bad request

swagger:response deleteAdvantagesIdBadRequest
*/
type DeleteAdvantagesIDBadRequest struct {
}

// NewDeleteAdvantagesIDBadRequest creates DeleteAdvantagesIDBadRequest with default headers values
func NewDeleteAdvantagesIDBadRequest() *DeleteAdvantagesIDBadRequest {

	return &DeleteAdvantagesIDBadRequest{}
}

// WriteResponse to the client
func (o *DeleteAdvantagesIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// DeleteAdvantagesIDInternalServerErrorCode is the HTTP code returned for type DeleteAdvantagesIDInternalServerError
const DeleteAdvantagesIDInternalServerErrorCode int = 500

/*DeleteAdvantagesIDInternalServerError Internal Server Error

swagger:response deleteAdvantagesIdInternalServerError
*/
type DeleteAdvantagesIDInternalServerError struct {
}

// NewDeleteAdvantagesIDInternalServerError creates DeleteAdvantagesIDInternalServerError with default headers values
func NewDeleteAdvantagesIDInternalServerError() *DeleteAdvantagesIDInternalServerError {

	return &DeleteAdvantagesIDInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteAdvantagesIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}