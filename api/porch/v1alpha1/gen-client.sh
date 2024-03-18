#!/usr/bin/env bash
go run k8s.io/code-generator/cmd/client-gen@v0.29.2 \
 --plural-exceptions PackageRevisionResources:PackageRevisionResources \
 --input-base "github.com/nephio-project/porch" \
 --input-dirs "../.." \
 --input "api/porch/v1alpha1" \
 --output-package "github.com/nephio-project/porch/api/generated/clientset"  \
 --clientset-name versioned \
 --trim-path-prefix "github.com/nephio-project/porch/api" \
 --go-header-file ../../../scripts/boilerplate.go.txt