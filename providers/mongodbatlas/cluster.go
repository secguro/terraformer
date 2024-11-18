// Copyright 2019 The Terraformer Authors.
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
	"context"
	"fmt"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	mongodbatlas "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

type ClusterGenerator struct {
	MongoDBAtlasService
}

func (g *ClusterGenerator) createClusterResources(client *mongodbatlas.APIClient) error {
	ctx := context.Background()

	log.Println("bumm")

	orgGroups, _, err := client.ClustersApi.ListClustersForAllProjects(ctx).Execute()
	if err != nil {
		return err
	}

	for _, orgGroup := range orgGroups.GetResults() {
		log.Println(*orgGroup.GroupId)
		for _, cluster := range orgGroup.GetClusters() {
			log.Println(*cluster.Name)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				fmt.Sprintf("%s-%s", *orgGroup.GroupId, *cluster.Name),
				fmt.Sprintf("%s-%s", *orgGroup.GroupId, *cluster.Name),
				//*orgGroup.GroupId,
				//*cluster.Name,
				"mongodbatlas_advanced_cluster",
				g.ProviderName,
				[]string{}))
		}
	}

	return nil
}

func (g *ClusterGenerator) InitResources() error {
	client, err := g.generateClient()
	if err != nil {
		return err
	}

	err = g.createClusterResources(client)
	if err != nil {
		return err
	}

	return nil
}
