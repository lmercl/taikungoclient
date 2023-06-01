/*
Taikun - WebApi

Testing StandaloneVMDisksApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package taikuncore

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/chnyda/taikungoclient"
)

func Test_taikuncore_StandaloneVMDisksApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StandaloneVMDisksApiService StandalonevmdisksCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StandaloneVMDisksApi.StandalonevmdisksCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandaloneVMDisksApiService StandalonevmdisksDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.StandaloneVMDisksApi.StandalonevmdisksDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandaloneVMDisksApiService StandalonevmdisksPurge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.StandaloneVMDisksApi.StandalonevmdisksPurge(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandaloneVMDisksApiService StandalonevmdisksReset", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.StandaloneVMDisksApi.StandalonevmdisksReset(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandaloneVMDisksApiService StandalonevmdisksUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.StandaloneVMDisksApi.StandalonevmdisksUpdate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandaloneVMDisksApiService StandalonevmdisksUpdateSize", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.StandaloneVMDisksApi.StandalonevmdisksUpdateSize(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
