#!/usr/bin/env bash
 go run k8s.io/code-generator/cmd/client-gen \
  --clientset-name versioned \
  --input-base "" \
  --input-dirs "." \
  --input porch/v1alpha1 \
  -p ../../generated/clientset \
  --output-base "github.com/nephio-project/porch" \
  --plural-exceptions PackageRevisionResources:PackageRevisionResources \
  --go-header-file ../../../scripts/boilerplate.go.txt