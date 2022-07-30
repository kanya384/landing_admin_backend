// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// GetActionsOKCode is the HTTP code returned for type GetActionsOK
const GetActionsOKCode int = 200

/*GetActionsOK returns filtered actions

swagger:response getActionsOK
*/
type GetActionsOK struct {

	/*
	  In: Body
	*/
	Payload models.GetActionsResponse `json:"body,omitempty"`
}

// NewGetActionsOK creates GetActionsOK with default headers values
func NewGetActionsOK() *GetActionsOK {

	return &GetActionsOK{}
}

// WithPayload adds the payload to the get actions o k response
func (o *GetActionsOK) WithPayload(payload models.GetActionsResponse) *GetActionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get actions o k response
func (o *GetActionsOK) SetPayload(payload models.GetActionsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetActionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.GetActionsResponse{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetActionsBadRequestCode is the HTTP code returned for type GetActionsBadRequest
const GetActionsBadRequestCode int = 400

/*GetActionsBadRequest Bad request

swagger:response getActionsBadRequest
*/
type GetActionsBadRequest struct {
}

// NewGetActionsBadRequest creates GetActionsBadRequest with default headers values
func NewGetActionsBadRequest() *GetActionsBadRequest {

	return &GetActionsBadRequest{}
}

// WriteResponse to the client
func (o *GetActionsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetActionsInternalServerErrorCode is the HTTP code returned for type GetActionsInternalServerError
const GetActionsInternalServerErrorCode int = 500

/*GetActionsInternalServerError Internal Server Error

swagger:response getActionsInternalServerError
*/
type GetActionsInternalServerError struct {
}

// NewGetActionsInternalServerError creates GetActionsInternalServerError with default headers values
func NewGetActionsInternalServerError() *GetActionsInternalServerError {

	return &GetActionsInternalServerError{}
}

// WriteResponse to the client
func (o *GetActionsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}