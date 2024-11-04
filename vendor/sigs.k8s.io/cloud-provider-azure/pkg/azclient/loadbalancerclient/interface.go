/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// +azure:enableclientgen:=true
package loadbalancerclient

import (
	"context"

	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"

	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/utils"
)

// +azure:client:verbs=get;createorupdate;delete;list,resource=LoadBalancer,packageName=github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6,packageAlias=armnetwork,clientName=LoadBalancersClient,expand=true,rateLimitKey=loadBalancerRateLimit,etag=true,azureStackCloudAPIVersion="2018-11-01"
type Interface interface {
	utils.GetWithExpandFunc[armnetwork.LoadBalancer]
	utils.CreateOrUpdateFunc[armnetwork.LoadBalancer]
	utils.DeleteFunc[armnetwork.LoadBalancer]
	utils.ListFunc[armnetwork.LoadBalancer]
	MigrateToIPBased(ctx context.Context, groupName string, loadBalancerName string, options *armnetwork.LoadBalancersClientMigrateToIPBasedOptions) (armnetwork.LoadBalancersClientMigrateToIPBasedResponse, error)
}
