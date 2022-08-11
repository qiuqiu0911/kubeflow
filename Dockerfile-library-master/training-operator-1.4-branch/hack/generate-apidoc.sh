#!/bin/bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This shell is used to auto generate some useful tools for k8s, such as lister,
# informer, deepcopy, defaulter and so on.

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..

cd ${SCRIPT_ROOT}

crd-ref-docs --log-level DEBUG\
    --source-path ./pkg/apis/tensorflow/v1 \
		--config ./docs/api/autogen/config.yaml \
		--templates-dir ./docs/api/autogen/templates \
		--output-path ./docs/api/tensorflow_generated.asciidoc \
		--max-depth 30

crd-ref-docs --log-level DEBUG\
    --source-path ./pkg/apis/pytorch/v1 \
		--config ./docs/api/autogen/config.yaml \
		--templates-dir ./docs/api/autogen/templates \
		--output-path ./docs/api/pytorch_generated.asciidoc \
		--max-depth 30

crd-ref-docs --log-level DEBUG\
    --source-path ./pkg/apis/mxnet/v1 \
		--config ./docs/api/autogen/config.yaml \
		--templates-dir ./docs/api/autogen/templates \
		--output-path ./docs/api/mxnet_generated.asciidoc \
		--max-depth 30

crd-ref-docs --log-level DEBUG\
    --source-path ./pkg/apis/xgboost/v1 \
		--config ./docs/api/autogen/config.yaml \
		--templates-dir ./docs/api/autogen/templates \
		--output-path ./docs/api/xgboost_generated.asciidoc \
		--max-depth 30

crd-ref-docs --log-level DEBUG\
    --source-path ./pkg/apis/mpi/v1 \
		--config ./docs/api/autogen/config.yaml \
		--templates-dir ./docs/api/autogen/templates \
		--output-path ./docs/api/mpi_generated.asciidoc \
		--max-depth 30

cd - > /dev/null
