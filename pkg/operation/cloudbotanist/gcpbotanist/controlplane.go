// Copyright 2018 The Gardener Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gcpbotanist

import (
	"fmt"
)

// GenerateCloudProviderConfig generates the GCE cloud provider config.
func (b *GCPBotanist) GenerateCloudProviderConfig() (string, error) {
	networkName := b.VPCName
	if networkName == "" {
		networkName = b.Shoot.SeedNamespace
	}

	return `[Global]
project-id = ` + b.Project + `
network-name = ` + networkName + `
multizone = false
local-zone = ` + string(b.Shoot.Info.Spec.Cloud.GCP.Zones[0]) + `
node-tags = ` + b.Shoot.SeedNamespace, nil
}

// GenerateKubeAPIServerConfig generates the cloud provider specific values which are required to render the
// Deployment manifest of the kube-apiserver properly.
func (b *GCPBotanist) GenerateKubeAPIServerConfig() (map[string]interface{}, error) {
	return map[string]interface{}{
		"AdditionalParameters": []string{
			fmt.Sprintf("--google-json-key=/srv/cloudprovider/%s", ServiceAccountJSON),
		},
	}, nil
}

// GenerateKubeControllerManagerConfig generates the cloud provider specific values which are required to
// render the Deployment manifest of the kube-controller-manager properly.
func (b *GCPBotanist) GenerateKubeControllerManagerConfig() (map[string]interface{}, error) {
	return map[string]interface{}{
		"AdditionalParameters": []string{
			fmt.Sprintf("--google-json-key=/srv/cloudprovider/%s", ServiceAccountJSON),
		},
	}, nil
}

// GenerateKubeSchedulerConfig generates the cloud provider specific values which are required to render the
// Deployment manifest of the kube-scheduler properly.
func (b *GCPBotanist) GenerateKubeSchedulerConfig() (map[string]interface{}, error) {
	return nil, nil
}

// DeployAutoNodeRepair deploys the auto-node-repair into the Seed cluster. It primary job is to repair
// unHealthy Nodes by replacing them by newer ones - Not needed on GCP yet. To be implemented later
func (b *GCPBotanist) DeployAutoNodeRepair() error {
	return nil
}

// GenerateEtcdBackupSecretData generates the secret data for etcd-operator having credentials and config for
// backup infrastructure.
// TODO: implement backup functionality for GCP
func (b *GCPBotanist) GenerateEtcdBackupSecretData() (map[string][]byte, error) {
	return nil, nil
}

// GenerateEtcdConfig returns the map of Cloud specific backup configuration required by etcd-operator.
// TODO: implement backup functionality for GCP
func (b *GCPBotanist) GenerateEtcdConfig(string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"kind": "StatefulSet",
	}, nil
}
