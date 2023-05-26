/*
Taikun - WebApi

Testing ProjectActionsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package taikungoclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/chnyda/testgoclient"
)

func Test_taikungoclient_ProjectActionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectActionsApiService ProjectactionsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		httpRes, err := apiClient.ProjectActionsApi.ProjectactionsDelete(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectActionsApiService ProjectactionsEdit", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		httpRes, err := apiClient.ProjectActionsApi.ProjectactionsEdit(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
