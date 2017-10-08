package kubeopbot

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AlertList is a list of Alert objects.
type AlertList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []Alert
}

type AlertSpec struct {
}

type AlertStatus struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Alert struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   AlertSpec
	Status AlertStatus
}
