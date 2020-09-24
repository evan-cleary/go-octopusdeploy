package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewProjectService(t *testing.T) {
	serviceFunction := newProjectService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceProjectService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *projectService
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

func TestProjectServiceGetWithEmptyID(t *testing.T) {
	service := newProjectService(&sling.Sling{}, emptyString)

	resource, err := service.GetByID(emptyString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)
}
