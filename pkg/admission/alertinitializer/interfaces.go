/*
Copyright 2017 The Kubernetes Authors.

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

package alertinitializer

import (
	informers "github.com/maleck13/kubeop-bot/pkg/client/informers_generated/internalversion"
	"k8s.io/apiserver/pkg/admission"
)

// WantsInternalAlertInformerFactory defines a function which sets InformerFactory for admission plugins that need it
type WantsInternalAlertInformerFactory interface {
	SetInternalAlertInformerFactory(informers.SharedInformerFactory)
	admission.Validator
}
