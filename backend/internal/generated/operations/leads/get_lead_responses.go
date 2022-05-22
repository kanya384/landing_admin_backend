// Code generated by go-swagger; DO NOT EDIT.

package leads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// GetLeadOKCode is the HTTP code returned for type GetLeadOK
const GetLeadOKCode int = 200

/*GetLeadOK returns leads list

swagger:response getLeadOK
*/
type GetLeadOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Lead `json:"body,omitempty"`
}

// NewGetLeadOK creates GetLeadOK with default headers values
func NewGetLeadOK() *GetLeadOK {

	return &GetLeadOK{}
}

// WithPayload adds the payload to the get lead o k response
func (o *GetLeadOK) WithPayload(payload []*models.Lead) *GetLeadOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get lead o k response
func (o *GetLeadOK) SetPayload(payload []*models.Lead) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLeadOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Lead, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetLeadBadRequestCode is the HTTP code returned for type GetLeadBadRequest
const GetLeadBadRequestCode int = 400

/*GetLeadBadRequest Bad request

swagger:response getLeadBadRequest
*/
type GetLeadBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewGetLeadBadRequest creates GetLeadBadRequest with default headers values
func NewGetLeadBadRequest() *GetLeadBadRequest {

	return &GetLeadBadRequest{}
}

// WithPayload adds the payload to the get lead bad request response
func (o *GetLeadBadRequest) WithPayload(payload *models.ResultResponse) *GetLeadBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get lead bad request response
func (o *GetLeadBadRequest) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLeadBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetLeadInternalServerErrorCode is the HTTP code returned for type GetLeadInternalServerError
const GetLeadInternalServerErrorCode int = 500

/*GetLeadInternalServerError Internal Server Error

swagger:response getLeadInternalServerError
*/
type GetLeadInternalServerError struct {
}

// NewGetLeadInternalServerError creates GetLeadInternalServerError with default headers values
func NewGetLeadInternalServerError() *GetLeadInternalServerError {

	return &GetLeadInternalServerError{}
}

// WriteResponse to the client
func (o *GetLeadInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}