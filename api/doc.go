// Copyright 2022 The kpt and Nephio Authors
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

// +domain=kpt.dev
//go:generate go run k8s.io/code-generator/cmd/client-gen@v0.29.2  --plural-exceptions PackageRevisionResources:PackageRevisionResources  --input-base "github.com/nephio-project/porch"  --input-dirs ""  --input "api/porch/v1alpha1"  --output-package "github.com/nephio-project/porch/api/generated/clientset"   --clientset-name versioned  --trim-path-prefix "github.com/nephio-project/porch/api"  --go-header-file ../scripts/boilerplate.go.txt
//go:generate go run k8s.io/code-generator/cmd/lister-gen --input-dirs github.com/nephio-project/porch/api/porch/v1alpha1 -p ../../generated/listers --go-header-file ../scripts/boilerplate.go.txt
//go:generate go run k8s.io/code-generator/cmd/informer-gen --input-dirs github.com/nephio-project/porch/api/porch/v1alpha1 --versioned-clientset-package github.com/nephio-project/porch/api/generated/clientset/versioned --listers-package github.com/nephio-project/porch/api/generated/listers --internal-clientset-package github.com/nephio-project/porch/api/generated/informers/internalversion -p github.com/nephio-project/porch/api/generated/informers --trim-path-prefix "github.com/nephio-project/porch/api" --plural-exceptions PackageRevisionResources:PackageRevisionResources --go-header-file ../scripts/boilerplate.go.txt
//go:generate go run k8s.io/code-generator/cmd/conversion-gen --input-dirs ./porch -O zz_generated.conversion --go-header-file ../scripts/boilerplate.go.txt

package apis
