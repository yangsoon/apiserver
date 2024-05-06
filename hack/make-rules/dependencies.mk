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

# This Makefile is a copy with modifications from https://github.com/marmotedu/iam/blob/master/scripts/make-rules/dependencies.mk

# ==============================================================================
# Makefile helper functions for dependencies
#

.PHONY: dependencies.run
dependencies.run: dependencies.packages dependencies.tools

.PHONY: dependencies.packages
dependencies.packages:
	@$(GO) mod tidy

.PHONY: dependencies.tools
dependencies.tools: dependencies.tools.blocker dependencies.tools.critical

.PHONY: dependencies.tools.blocker
dependencies.tools.blocker: go.build.verify $(addprefix tools.verify., $(BLOCKER_TOOLS))

.PHONY: dependencies.tools.critical
dependencies.tools.critical: $(addprefix tools.verify., $(CRITICAL_TOOLS))

.PHONY: dependencies.tools.trivial
dependencies.tools.trivial: $(addprefix tools.verify., $(TRIVIAL_TOOLS))