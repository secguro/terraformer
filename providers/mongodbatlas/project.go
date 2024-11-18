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

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	mongodbatlas "go.mongodb.org/atlas-sdk/v20241023001/admin"
)

type ProjectGenerator struct {
	MongoDBAtlasService
}

func (g *ProjectGenerator) createProjectResources(client *mongodbatlas.APIClient) error {
	ctx := context.Background()

	projects, _, err := client.ProjectsApi.ListProjects(ctx).Execute()
	if err != nil {
		return err
	}

	for _, group := range projects.GetResults() {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			*group.Id,
			*group.Id,
			"mongodbatlas_project",
			g.ProviderName,
			[]string{}))
	}

	return nil
}

func (g *ProjectGenerator) InitResources() error {
	client, err := g.generateClient()
	if err != nil {
		return err
	}

	err = g.createProjectResources(client)
	if err != nil {
		return err
	}

	return nil
}
