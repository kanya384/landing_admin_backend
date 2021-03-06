// Code generated by go-swagger; DO NOT EDIT.

package hod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// GetMonthsIDOKCode is the HTTP code returned for type GetMonthsIDOK
const GetMonthsIDOKCode int = 200

/*GetMonthsIDOK month get success

swagger:response getMonthsIdOK
*/
type GetMonthsIDOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Month `json:"body,omitempty"`
}

// NewGetMonthsIDOK creates GetMonthsIDOK with default headers values
func NewGetMonthsIDOK() *GetMonthsIDOK {

	return &GetMonthsIDOK{}
}

// WithPayload adds the payload to the get months Id o k response
func (o *GetMonthsIDOK) WithPayload(payload []*models.Month) *GetMonthsIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get months Id o k response
func (o *GetMonthsIDOK) SetPayload(payload []*models.Month) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMonthsIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Month, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetMonthsIDBadRequestCode is the HTTP code returned for type GetMonthsIDBadRequest
const GetMonthsIDBadRequestCode int = 400

/*GetMonthsIDBadRequest Bad request

swagger:response getMonthsIdBadRequest
*/
type GetMonthsIDBadRequest struct {
}

// NewGetMonthsIDBadRequest creates GetMonthsIDBadRequest with default headers values
func NewGetMonthsIDBadRequest() *GetMonthsIDBadRequest {

	return &GetMonthsIDBadRequest{}
}

// WriteResponse to the client
func (o *GetMonthsIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetMonthsIDInternalServerErrorCode is the HTTP code returned for type GetMonthsIDInternalServerError
const GetMonthsIDInternalServerErrorCode int = 500

/*GetMonthsIDInternalServerError Internal Server Error

swagger:response getMonthsIdInternalServerError
*/
type GetMonthsIDInternalServerError struct {
}

// NewGetMonthsIDInternalServerError creates GetMonthsIDInternalServerError with default headers values
func NewGetMonthsIDInternalServerError() *GetMonthsIDInternalServerError {

	return &GetMonthsIDInternalServerError{}
}

// WriteResponse to the client
func (o *GetMonthsIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
