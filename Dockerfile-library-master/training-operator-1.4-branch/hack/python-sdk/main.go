/*
Copyright 2021 kubeflow.org.

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

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/go-openapi/spec"
	mxnet "github.com/kubeflow/training-operator/pkg/apis/mxnet/v1"
	pytorch "github.com/kubeflow/training-operator/pkg/apis/pytorch/v1"
	tensorflow "github.com/kubeflow/training-operator/pkg/apis/tensorflow/v1"
	xgboost "github.com/kubeflow/training-operator/pkg/apis/xgboost/v1"
	"k8s.io/klog"
	"k8s.io/kube-openapi/pkg/common"
)

// Generate OpenAPI spec definitions for API resources
func main() {
	if len(os.Args) <= 1 {
		klog.Fatal("Supply a version")
	}
	version := os.Args[1]
	if !strings.HasPrefix(version, "v") {
		version = "v" + version
	}
	var oAPIDefs = map[string]common.OpenAPIDefinition{}
	defs := spec.Definitions{}

	refCallback := func(name string) spec.Ref {
		return spec.MustCreateRef("#/definitions/" + common.EscapeJsonPointer(swaggify(name)))
	}

	for k, v := range tensorflow.GetOpenAPIDefinitions(refCallback) {
		oAPIDefs[k] = v
	}

	for k, v := range pytorch.GetOpenAPIDefinitions(refCallback) {
		oAPIDefs[k] = v
	}

	for k, v := range mxnet.GetOpenAPIDefinitions(refCallback) {
		oAPIDefs[k] = v
	}

	for k, v := range xgboost.GetOpenAPIDefinitions(refCallback) {
		oAPIDefs[k] = v
	}

	for defName, val := range oAPIDefs {
		defs[swaggify(defName)] = val.Schema
	}
	swagger := spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Swagger:     "2.0",
			Definitions: defs,
			Paths:       &spec.Paths{Paths: map[string]spec.PathItem{}},
			Info: &spec.Info{
				InfoProps: spec.InfoProps{
					Title:       "Kubeflow Training SDK",
					Description: "Python SDK for Kubeflow Training",
					Version:     version,
				},
			},
		},
	}
	jsonBytes, err := json.MarshalIndent(swagger, "", "  ")
	if err != nil {
		klog.Fatal(err.Error())
	}
	fmt.Println(string(jsonBytes))
}

func swaggify(name string) string {
	name = strings.Replace(name, "github.com/kubeflow/training-operator/pkg/apis/tensorflow/", "", -1)
	name = strings.Replace(name, "github.com/kubeflow/training-operator/pkg/apis/pytorch/", "", -1)
	name = strings.Replace(name, "github.com/kubeflow/training-operator/pkg/apis/mxnet/", "", -1)
	name = strings.Replace(name, "github.com/kubeflow/training-operator/pkg/apis/xgboost/", "", -1)
	name = strings.Replace(name, "github.com/kubeflow/common/pkg/apis/common/", "", -1)
	name = strings.Replace(name, "k8s.io/api/core/", "", -1)
	name = strings.Replace(name, "k8s.io/apimachinery/pkg/apis/meta/", "", -1)
	name = strings.Replace(name, "k8s.io/apimachinery/pkg/api/resource", "", -1)
	name = strings.Replace(name, "/", ".", -1)
	return name
}
