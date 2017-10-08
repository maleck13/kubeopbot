// +build !ignore_autogenerated

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	kubeopbot "github.com/maleck13/kubeop-bot/pkg/apis/kubeopbot"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_Alert_To_kubeopbot_Alert,
		Convert_kubeopbot_Alert_To_v1alpha1_Alert,
		Convert_v1alpha1_AlertList_To_kubeopbot_AlertList,
		Convert_kubeopbot_AlertList_To_v1alpha1_AlertList,
		Convert_v1alpha1_AlertSpec_To_kubeopbot_AlertSpec,
		Convert_kubeopbot_AlertSpec_To_v1alpha1_AlertSpec,
		Convert_v1alpha1_AlertStatus_To_kubeopbot_AlertStatus,
		Convert_kubeopbot_AlertStatus_To_v1alpha1_AlertStatus,
	)
}

func autoConvert_v1alpha1_Alert_To_kubeopbot_Alert(in *Alert, out *kubeopbot.Alert, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_AlertSpec_To_kubeopbot_AlertSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_AlertStatus_To_kubeopbot_AlertStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Alert_To_kubeopbot_Alert is an autogenerated conversion function.
func Convert_v1alpha1_Alert_To_kubeopbot_Alert(in *Alert, out *kubeopbot.Alert, s conversion.Scope) error {
	return autoConvert_v1alpha1_Alert_To_kubeopbot_Alert(in, out, s)
}

func autoConvert_kubeopbot_Alert_To_v1alpha1_Alert(in *kubeopbot.Alert, out *Alert, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_kubeopbot_AlertSpec_To_v1alpha1_AlertSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_kubeopbot_AlertStatus_To_v1alpha1_AlertStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_kubeopbot_Alert_To_v1alpha1_Alert is an autogenerated conversion function.
func Convert_kubeopbot_Alert_To_v1alpha1_Alert(in *kubeopbot.Alert, out *Alert, s conversion.Scope) error {
	return autoConvert_kubeopbot_Alert_To_v1alpha1_Alert(in, out, s)
}

func autoConvert_v1alpha1_AlertList_To_kubeopbot_AlertList(in *AlertList, out *kubeopbot.AlertList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]kubeopbot.Alert)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_AlertList_To_kubeopbot_AlertList is an autogenerated conversion function.
func Convert_v1alpha1_AlertList_To_kubeopbot_AlertList(in *AlertList, out *kubeopbot.AlertList, s conversion.Scope) error {
	return autoConvert_v1alpha1_AlertList_To_kubeopbot_AlertList(in, out, s)
}

func autoConvert_kubeopbot_AlertList_To_v1alpha1_AlertList(in *kubeopbot.AlertList, out *AlertList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Alert)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_kubeopbot_AlertList_To_v1alpha1_AlertList is an autogenerated conversion function.
func Convert_kubeopbot_AlertList_To_v1alpha1_AlertList(in *kubeopbot.AlertList, out *AlertList, s conversion.Scope) error {
	return autoConvert_kubeopbot_AlertList_To_v1alpha1_AlertList(in, out, s)
}

func autoConvert_v1alpha1_AlertSpec_To_kubeopbot_AlertSpec(in *AlertSpec, out *kubeopbot.AlertSpec, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_AlertSpec_To_kubeopbot_AlertSpec is an autogenerated conversion function.
func Convert_v1alpha1_AlertSpec_To_kubeopbot_AlertSpec(in *AlertSpec, out *kubeopbot.AlertSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_AlertSpec_To_kubeopbot_AlertSpec(in, out, s)
}

func autoConvert_kubeopbot_AlertSpec_To_v1alpha1_AlertSpec(in *kubeopbot.AlertSpec, out *AlertSpec, s conversion.Scope) error {
	return nil
}

// Convert_kubeopbot_AlertSpec_To_v1alpha1_AlertSpec is an autogenerated conversion function.
func Convert_kubeopbot_AlertSpec_To_v1alpha1_AlertSpec(in *kubeopbot.AlertSpec, out *AlertSpec, s conversion.Scope) error {
	return autoConvert_kubeopbot_AlertSpec_To_v1alpha1_AlertSpec(in, out, s)
}

func autoConvert_v1alpha1_AlertStatus_To_kubeopbot_AlertStatus(in *AlertStatus, out *kubeopbot.AlertStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_AlertStatus_To_kubeopbot_AlertStatus is an autogenerated conversion function.
func Convert_v1alpha1_AlertStatus_To_kubeopbot_AlertStatus(in *AlertStatus, out *kubeopbot.AlertStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_AlertStatus_To_kubeopbot_AlertStatus(in, out, s)
}

func autoConvert_kubeopbot_AlertStatus_To_v1alpha1_AlertStatus(in *kubeopbot.AlertStatus, out *AlertStatus, s conversion.Scope) error {
	return nil
}

// Convert_kubeopbot_AlertStatus_To_v1alpha1_AlertStatus is an autogenerated conversion function.
func Convert_kubeopbot_AlertStatus_To_v1alpha1_AlertStatus(in *kubeopbot.AlertStatus, out *AlertStatus, s conversion.Scope) error {
	return autoConvert_kubeopbot_AlertStatus_To_v1alpha1_AlertStatus(in, out, s)
}