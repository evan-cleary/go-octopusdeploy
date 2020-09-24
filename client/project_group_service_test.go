package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewProjectGroupService(t *testing.T) {
	serviceFunction := newProjectGroupService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceProjectGroupService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *projectGroupService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate},
		{"EmptyURITemplate", serviceFunction, client, emptyString},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}

func TestProjectGroupServiceGetWithEmptyID(t *testing.T) {
	service := newProjectGroupService(&sling.Sling{}, emptyString)

	resource, err := service.GetByID(emptyString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)
}
