// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cmd

import (
	mongodbatlas_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/mongodbatlas"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdMongoDBAtlasImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mongodbatlas",
		Short: "Import current state to Terraform configuration from MongoDB Atlas",
		Long:  "Import current state to Terraform configuration from MongoDB Atlas",
		RunE: func(cmd *cobra.Command, args []string) error {
			// pubKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
			// if len(pubKey) == 0 {
			// 	return errors.New("MONGODB_ATLAS_PUBLIC_KEY for MongoDB Atlas must be set through `MONGODB_ATLAS_PUBLIC_KEY` env var")
			// }
			// privKey := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")
			// if len(privKey) == 0 {
			// 	return errors.New("MONGODB_ATLAS_PRIVATE_KEY for MongoDB Atlas must be set through `MONGODB_ATLAS_PRIVATE_KEY` env var")
			// }

			provider := newMongoDBAtlasImporter()
			err := Import(provider, options, []string{ /*pubKey, privKey*/ })
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newMongoDBAtlasImporter()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "", "")
	return cmd
}

func newMongoDBAtlasImporter() terraformutils.ProviderGenerator {
	return &mongodbatlas_terraforming.MongoDBAtlasProvider{}
}
