/*
Taikun - WebApi

Testing CatalogAppApiService

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

func Test_taikungoclient_CatalogAppApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CatalogAppApiService CatalogAppCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CatalogAppApi.CatalogAppCreate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogAppApiService CatalogAppDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		httpRes, err := apiClient.CatalogAppApi.CatalogAppDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogAppApiService CatalogAppDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var catalogAppId int32

		resp, httpRes, err := apiClient.CatalogAppApi.CatalogAppDetails(context.Background(), catalogAppId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogAppApiService CatalogAppEditParams", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CatalogAppApi.CatalogAppEditParams(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogAppApiService CatalogAppEditVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CatalogAppApi.CatalogAppEditVersion(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogAppApiService CatalogAppLockManager", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CatalogAppApi.CatalogAppLockManager(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogAppApiService CatalogAppParamDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var catalogAppId int32

		resp, httpRes, err := apiClient.CatalogAppApi.CatalogAppParamDetails(context.Background(), catalogAppId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
