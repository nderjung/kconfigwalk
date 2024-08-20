# SPDX-License-Identifier: BSD-3-Clause
# Copyright (c) 2024, Alexander Jung.
# Licensed under the BSD-3-Clause License (the "License").
# You may not use this file except in compliance with the License.

WORKDIR ?= $(CURDIR)
GO      ?= go

.PHONY: all
all: kconfigwalk

.PHONY: kconfigwalk
kconfigwalk: tidy
	$(GO) build -v -o $@ ./...

.PHONY: tidy
tidy:
	$(GO) mod tidy
