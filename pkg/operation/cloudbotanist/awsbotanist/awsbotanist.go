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

package awsbotanist

import (
	"errors"

	"github.com/gardener/gardener/pkg/client/aws"
	"github.com/gardener/gardener/pkg/operation"
)

// New takes an operation object <o> and creates a new AWSBotanist object.
func New(o *operation.Operation) (*AWSBotanist, error) {
	region := o.Shoot.Info.Spec.Cloud.Region

	if o.Shoot.Info.Spec.Cloud.AWS == nil {
		return nil, errors.New("cannot instantiate an AWS botanist if `.spec.cloud.aws` is nil")
	}

	return &AWSBotanist{
		Operation:         o,
		CloudProviderName: "aws",
		SeedAWSClient: aws.NewClient(
			string(o.Seed.Secret.Data[AccessKeyID]),
			string(o.Seed.Secret.Data[SecretAccessKey]),
			region,
		),
		ShootAWSClient: aws.NewClient(
			string(o.Shoot.Secret.Data[AccessKeyID]),
			string(o.Shoot.Secret.Data[SecretAccessKey]),
			region,
		),
	}, nil
}

// GetCloudProviderName returns the Kubernetes cloud provider name for this cloud.
func (b *AWSBotanist) GetCloudProviderName() string {
	return b.CloudProviderName
}
