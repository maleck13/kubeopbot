package alert

import (
	"fmt"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"

	"github.com/maleck13/kubeop-bot/pkg/apis/kubeopbot"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
)

func NewStrategy(typer runtime.ObjectTyper) alertStrategy {
	return alertStrategy{typer, names.SimpleNameGenerator}
}

func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	apiserver, ok := obj.(*kubeopbot.Alert)
	if !ok {
		return nil, nil, false, fmt.Errorf("given object is not an Alert.")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), alertToSelectableFields(apiserver), apiserver.Initializers != nil, nil
}

// MatchAlert is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func MatchAlert(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// alertToSelectableFields returns a field set that represents the object.
func alertToSelectableFields(obj *kubeopbot.Alert) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

type alertStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (alertStrategy) NamespaceScoped() bool {
	return false
}

func (alertStrategy) PrepareForCreate(ctx genericapirequest.Context, obj runtime.Object) {
}

func (alertStrategy) PrepareForUpdate(ctx genericapirequest.Context, obj, old runtime.Object) {
}

func (alertStrategy) Validate(ctx genericapirequest.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (alertStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (alertStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (alertStrategy) Canonicalize(obj runtime.Object) {
}

func (alertStrategy) ValidateUpdate(ctx genericapirequest.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
