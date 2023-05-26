/*
Taikun - WebApi

Testing KubeConfigRoleApiService

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

func Test_taikungoclient_KubeConfigRoleApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test KubeConfigRoleApiService KubeconfigroleList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubeConfigRoleApi.KubeconfigroleList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
