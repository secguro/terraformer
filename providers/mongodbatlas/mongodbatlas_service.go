// Copyright 2018 The Terraformer Authors.
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

package mongodbatlas

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	mongodbatlas "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

type MongoDBAtlasService struct { //nolint
	terraformutils.Service
}

func (s *MongoDBAtlasService) generateClient() (*mongodbatlas.APIClient, error) {
	return mongodbatlas.NewClient(mongodbatlas.UseDigestAuth("dehqlrhv", "95fe70b1-4e9f-4b20-a904-d75b3e53e039" /*s.Args["pubKey"].(string), s.Args["privKey"].(string)*/))
}
