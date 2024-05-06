# Copyright 2024 The RSI Authors.
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

# ==============================================================================
# Makefile helper functions for code-gen
#

API_DIRS ?= $(filter-out %.go, $(wildcard ${ROOT_DIR}/pkg/apis/*))
APIS ?= $(foreach api,${API_DIRS},$(notdir ${api}))

.PHONY: code-gen.run
code-gen.run: code-gen.deepcopy code-gen.defaulter code-gen.conversion code-gen.openapi

.PHONY: code-gen.deepcopy.%
code-gen.deepcopy.%:
	@echo "===========> Generate deepcopy code for /pkg/apis/$*"
	@deepcopy-gen --go-header-file $(ROOT_DIR)/hack/boilerplate/boilerplate.go.txt \
		--input-dirs $(ROOT_PACKAGE)/pkg/apis/$* \
		--output-base $(ROOT_DIR)/_output \
		-O zz_generated.deepcopy \
		-v $(V)
	@if [ $$? -ne 0 ]; then \
		echo "===========> !!! Fail to generate deepcopy code for /pkg/apis/$*, Please run \"make code-gen V=10\" see more detail !!!!"; \
		exit 1; \
	fi
	@$(if $(wildcard $(ROOT_DIR)/_output/$(ROOT_PACKAGE)/pkg/apis/$*), \
		rsync -avr -W $(ROOT_DIR)/_output/$(ROOT_PACKAGE)/pkg/apis/$*/ $(ROOT_DIR)/pkg/apis/$* )

.PHONY: code-gen.deepcopy
code-gen.deepcopy: tools.verify.deepcopy-gen $(addprefix code-gen.deepcopy., $(APIS))

.PHONY: code-gen.defaulter.%
code-gen.defaulter.%:
	@echo "===========> Generate default code for /pkg/apis/$*"
	$(eval VERSION_FOLDERS := $(shell find $(ROOT_DIR)/pkg/apis/$* -type d -name 'v*'))
	$(foreach folder,$(VERSION_FOLDERS),$(eval VERSION_ARRAY += --input-dirs $(ROOT_PACKAGE)/pkg/apis/$*/$(notdir $(folder))))
	@defaulter-gen --go-header-file $(ROOT_DIR)/hack/boilerplate/boilerplate.go.txt \
		$(VERSION_ARRAY) \
		--output-base $(ROOT_DIR)/_output \
		-O zz_generated.defaults \
		-v $(V)
	@if [ $$? -ne 0 ]; then \
		echo "===========> !!! Fail to generate default code for /pkg/apis/$*, Please run \"make code-gen V=10\" see more detail !!!!"; \
		exit 1; \
	fi
	@$(if $(wildcard $(ROOT_DIR)/_output/$(ROOT_PACKAGE)/pkg/apis/$*/v*), \
		rsync -avr -W $(ROOT_DIR)/_output/$(ROOT_PACKAGE)/pkg/apis/$*/ $(ROOT_DIR)/pkg/apis/$* )

.PHONY: code-gen.defaulter
code-gen.defaulter: tools.verify.defaulter-gen $(addprefix code-gen.defaulter., $(APIS))

.PHONY: code-gen.conversion.%
code-gen.conversion.%:
	@echo "===========> Generate conversion code for /pkg/apis/$*"
	$(eval VERSION_FOLDERS := $(shell find $(ROOT_DIR)/pkg/apis/$* -type d -name 'v*'))
	$(foreach folder,$(VERSION_FOLDERS),$(eval VERSION_ARRAY += --input-dirs $(ROOT_PACKAGE)/pkg/apis/$*/$(notdir $(folder))))
	@conversion-gen --go-header-file $(ROOT_DIR)/hack/boilerplate/boilerplate.go.txt \
		$(VERSION_ARRAY) \
		--output-base $(ROOT_DIR)/_output \
		-O zz_generated.conversion \
		-v $(V)
	@if [ $$? -ne 0 ]; then \
		echo "===========> !!! Fail to generate conversion code for /pkg/apis/$*, Please run \"make code-gen V=10\" see more detail !!!!"; \
		exit 1; \
	fi
	@$(if $(wildcard $(ROOT_DIR)/_output/$(ROOT_PACKAGE)/pkg/apis/$*), \
		rsync -avr -W $(ROOT_DIR)/_output/$(ROOT_PACKAGE)/pkg/apis/$*/ $(ROOT_DIR)/pkg/apis/$* )

.PHONY: code-gen.conversion
code-gen.conversion: tools.verify.conversion-gen $(addprefix code-gen.conversion., $(APIS))

.PHONY: code-gen.openapi
code-gen.openapi: tools.verify.openapi-gen go.vendor
	@echo "===========> Generate openapi for rsi-apiserver"
	@openapi-gen \
		--go-header-file $(ROOT_DIR)/hack/boilerplate/boilerplate.go.txt \
		-O zz_generated.openapi \
		--report-filename "$(ROOT_DIR)/pkg/api/api-rules/violation_exceptions.list"\
		--input-dirs "k8s.io/apimachinery/pkg/api/resource" \
		--input-dirs "k8s.io/apimachinery/pkg/util/intstr" \
		--input-dirs "k8s.io/apimachinery/pkg/apis/meta/v1" \
		--input-dirs "k8s.io/apimachinery/pkg/runtime" \
		--input-dirs "k8s.io/apimachinery/pkg/version" \
		--input-dirs "github.com/openkruise/kruise-api/apps/pub" \
		--input-dirs "$(ROOT_PACKAGE)/vendor/k8s.io/api/core/v1" \
		--input-dirs "$(ROOT_PACKAGE)/vendor/k8s.io/api/apps/v1" \
		--input-dirs "$(ROOT_PACKAGE)/vendor/github.com/openkruise/kruise-api/apps/v1alpha1" \
		--output-package "$(ROOT_PACKAGE)/pkg/generated/openapi" \
		--output-base "$(ROOT_DIR)/_output" \
		-v $(V)
	@if [ $$? -ne 0 ]; then \
		echo "===========> !!! Fail to generate conversion code for /pkg/apis/$*, Please run \"make code-gen V=10\" see more detail !!!!"; \
		exit 1; \
	fi
	@$(if $(wildcard $(ROOT_DIR)/_output/$(ROOT_PACKAGE)/pkg/generated/openapi/$*), \
		rsync -avr -W $(ROOT_DIR)/_output/$(ROOT_PACKAGE)/pkg/generated/openapi/ $(ROOT_DIR)/pkg/generated/openapi )