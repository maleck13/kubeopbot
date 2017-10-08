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

package banalert

import (
	"fmt"
	"io"

	"github.com/maleck13/kubeop-bot/pkg/admission/alertinitializer"
	"github.com/maleck13/kubeop-bot/pkg/apis/kubeopbot"
	informers "github.com/maleck13/kubeop-bot/pkg/client/informers_generated/internalversion"
	listers "github.com/maleck13/kubeop-bot/pkg/client/listers_generated/kubeopbot/internalversion"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apiserver/pkg/admission"
)

// Register registers a plugin
func Register(plugins *admission.Plugins) {
	plugins.Register("BanAlert", func(config io.Reader) (admission.Interface, error) {
		return New()
	})
}

type disallowAlert struct {
	*admission.Handler
	lister listers.AlertLister
}

var _ = alertinitializer.WantsInternalAlertInformerFactory(&disallowAlert{})

// Admit ensures that the object in-flight is of kind Alert.
// In addition checks that the Name is not on the banned list.
// The list is stored in Alert API objects.
func (d *disallowAlert) Admit(a admission.Attributes) error {
	// we are only interested in alerts
	if a.GetKind().GroupKind() != kubeopbot.Kind("Alert") {
		return nil
	}

	metaAccessor, err := meta.Accessor(a.GetObject())
	if err != nil {
		return err
	}
	alertName := metaAccessor.GetName()

	if alertName == "" {
		return errors.NewForbidden(
			a.GetResource().GroupResource(),
			a.GetName(),
			fmt.Errorf("this name may not be used, please change the resource name"),
		)
	}

	return nil
}

// SetInternalAlertInformerFactory gets Lister from SharedInformerFactory.
// The lister knows how to lists Alerts.
func (d *disallowAlert) SetInternalAlertInformerFactory(f informers.SharedInformerFactory) {
	d.lister = f.Kubeopbot().InternalVersion().Alerts().Lister()
}

// Validate checks whether the plugin was correctly initialized.
func (d *disallowAlert) Validate() error {
	if d.lister == nil {
		return fmt.Errorf("missing alert lister")
	}
	return nil
}

// New creates a new ban alerts admission plugin
func New() (admission.Interface, error) {
	return &disallowAlert{
		Handler: admission.NewHandler(admission.Create),
	}, nil
}
