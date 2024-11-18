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
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type MongoDBAtlasProvider struct { //nolint
	terraformutils.Provider
	/*pubKey  string
	privKey string*/
}

func (p *MongoDBAtlasProvider) Init(args []string) error {
	/*pubKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
	if pubKey == "" {
		return errors.New("set MONGODB_ATLAS_PUBLIC_KEY env var")
	}
	p.pubKey = pubKey

	privKey := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")
	if privKey == "" {
		return errors.New("set MONGODB_ATLAS_PRIVATE_KEY env var")
	}
	p.privKey = privKey*/

	return nil
}

func (p *MongoDBAtlasProvider) GetName() string {
	return "mongodbatlas"
}

func (p *MongoDBAtlasProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		/*"pubKey":  cty.StringVal(p.pubKey),
		"privKey": cty.StringVal(p.privKey),*/
	})
}

func (p *MongoDBAtlasProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		/*"pubKey":  p.pubKey,
		"privKey": p.privKey,*/
	})
	return nil
}

func (p *MongoDBAtlasProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"mongodbatlas_project": &ProjectGenerator{},
		//"mongodbatlas_advanced_cluster": &ClusterGenerator{},
	}
}

func (p MongoDBAtlasProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p MongoDBAtlasProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}
