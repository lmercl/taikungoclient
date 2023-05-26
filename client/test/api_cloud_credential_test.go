/*
Taikun - WebApi

Testing CloudCredentialApiService

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

func Test_taikungoclient_CloudCredentialApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CloudCredentialApiService CloudcredentialsAllFlavors", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.CloudCredentialApi.CloudcredentialsAllFlavors(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudCredentialApiService CloudcredentialsDashboardList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CloudCredentialApi.CloudcredentialsDashboardList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudCredentialApiService CloudcredentialsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var cloudId int32

		httpRes, err := apiClient.CloudCredentialApi.CloudcredentialsDelete(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudCredentialApiService CloudcredentialsExceeded", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CloudCredentialApi.CloudcredentialsExceeded(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudCredentialApiService CloudcredentialsForCli", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CloudCredentialApi.CloudcredentialsForCli(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudCredentialApiService CloudcredentialsForProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CloudCredentialApi.CloudcredentialsForProject(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudCredentialApiService CloudcredentialsLockManager", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CloudCredentialApi.CloudcredentialsLockManager(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudCredentialApiService CloudcredentialsMakeDefault", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CloudCredentialApi.CloudcredentialsMakeDefault(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudCredentialApiService CloudcredentialsOrgList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CloudCredentialApi.CloudcredentialsOrgList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
