//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2024 The RSI Authors.

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
// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	appskruiseio "github.com/yangsoon/apiserver/pkg/apis/apps_kruise_io"
	pub "github.com/openkruise/kruise-api/apps/pub"
	v1alpha1 "github.com/openkruise/kruise-api/apps/v1alpha1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1alpha1.CloneSet)(nil), (*appskruiseio.CloneSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CloneSet_To_apps_kruise_io_CloneSet(a.(*v1alpha1.CloneSet), b.(*appskruiseio.CloneSet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*appskruiseio.CloneSet)(nil), (*v1alpha1.CloneSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_kruise_io_CloneSet_To_v1alpha1_CloneSet(a.(*appskruiseio.CloneSet), b.(*v1alpha1.CloneSet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.CloneSetCondition)(nil), (*appskruiseio.CloneSetCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CloneSetCondition_To_apps_kruise_io_CloneSetCondition(a.(*v1alpha1.CloneSetCondition), b.(*appskruiseio.CloneSetCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*appskruiseio.CloneSetCondition)(nil), (*v1alpha1.CloneSetCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_kruise_io_CloneSetCondition_To_v1alpha1_CloneSetCondition(a.(*appskruiseio.CloneSetCondition), b.(*v1alpha1.CloneSetCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.CloneSetList)(nil), (*appskruiseio.CloneSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CloneSetList_To_apps_kruise_io_CloneSetList(a.(*v1alpha1.CloneSetList), b.(*appskruiseio.CloneSetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*appskruiseio.CloneSetList)(nil), (*v1alpha1.CloneSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_kruise_io_CloneSetList_To_v1alpha1_CloneSetList(a.(*appskruiseio.CloneSetList), b.(*v1alpha1.CloneSetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.CloneSetScaleStrategy)(nil), (*appskruiseio.CloneSetScaleStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CloneSetScaleStrategy_To_apps_kruise_io_CloneSetScaleStrategy(a.(*v1alpha1.CloneSetScaleStrategy), b.(*appskruiseio.CloneSetScaleStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*appskruiseio.CloneSetScaleStrategy)(nil), (*v1alpha1.CloneSetScaleStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_kruise_io_CloneSetScaleStrategy_To_v1alpha1_CloneSetScaleStrategy(a.(*appskruiseio.CloneSetScaleStrategy), b.(*v1alpha1.CloneSetScaleStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.CloneSetSpec)(nil), (*appskruiseio.CloneSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CloneSetSpec_To_apps_kruise_io_CloneSetSpec(a.(*v1alpha1.CloneSetSpec), b.(*appskruiseio.CloneSetSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*appskruiseio.CloneSetSpec)(nil), (*v1alpha1.CloneSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_kruise_io_CloneSetSpec_To_v1alpha1_CloneSetSpec(a.(*appskruiseio.CloneSetSpec), b.(*v1alpha1.CloneSetSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.CloneSetStatus)(nil), (*appskruiseio.CloneSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CloneSetStatus_To_apps_kruise_io_CloneSetStatus(a.(*v1alpha1.CloneSetStatus), b.(*appskruiseio.CloneSetStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*appskruiseio.CloneSetStatus)(nil), (*v1alpha1.CloneSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_kruise_io_CloneSetStatus_To_v1alpha1_CloneSetStatus(a.(*appskruiseio.CloneSetStatus), b.(*v1alpha1.CloneSetStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.CloneSetUpdateStrategy)(nil), (*appskruiseio.CloneSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CloneSetUpdateStrategy_To_apps_kruise_io_CloneSetUpdateStrategy(a.(*v1alpha1.CloneSetUpdateStrategy), b.(*appskruiseio.CloneSetUpdateStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*appskruiseio.CloneSetUpdateStrategy)(nil), (*v1alpha1.CloneSetUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_kruise_io_CloneSetUpdateStrategy_To_v1alpha1_CloneSetUpdateStrategy(a.(*appskruiseio.CloneSetUpdateStrategy), b.(*v1alpha1.CloneSetUpdateStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.UpdateScatterTerm)(nil), (*appskruiseio.UpdateScatterTerm)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_UpdateScatterTerm_To_apps_kruise_io_UpdateScatterTerm(a.(*v1alpha1.UpdateScatterTerm), b.(*appskruiseio.UpdateScatterTerm), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*appskruiseio.UpdateScatterTerm)(nil), (*v1alpha1.UpdateScatterTerm)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apps_kruise_io_UpdateScatterTerm_To_v1alpha1_UpdateScatterTerm(a.(*appskruiseio.UpdateScatterTerm), b.(*v1alpha1.UpdateScatterTerm), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_CloneSet_To_apps_kruise_io_CloneSet(in *v1alpha1.CloneSet, out *appskruiseio.CloneSet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_CloneSetSpec_To_apps_kruise_io_CloneSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_CloneSetStatus_To_apps_kruise_io_CloneSetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_CloneSet_To_apps_kruise_io_CloneSet is an autogenerated conversion function.
func Convert_v1alpha1_CloneSet_To_apps_kruise_io_CloneSet(in *v1alpha1.CloneSet, out *appskruiseio.CloneSet, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloneSet_To_apps_kruise_io_CloneSet(in, out, s)
}

func autoConvert_apps_kruise_io_CloneSet_To_v1alpha1_CloneSet(in *appskruiseio.CloneSet, out *v1alpha1.CloneSet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_apps_kruise_io_CloneSetSpec_To_v1alpha1_CloneSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_apps_kruise_io_CloneSetStatus_To_v1alpha1_CloneSetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_apps_kruise_io_CloneSet_To_v1alpha1_CloneSet is an autogenerated conversion function.
func Convert_apps_kruise_io_CloneSet_To_v1alpha1_CloneSet(in *appskruiseio.CloneSet, out *v1alpha1.CloneSet, s conversion.Scope) error {
	return autoConvert_apps_kruise_io_CloneSet_To_v1alpha1_CloneSet(in, out, s)
}

func autoConvert_v1alpha1_CloneSetCondition_To_apps_kruise_io_CloneSetCondition(in *v1alpha1.CloneSetCondition, out *appskruiseio.CloneSetCondition, s conversion.Scope) error {
	out.Type = appskruiseio.CloneSetConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1alpha1_CloneSetCondition_To_apps_kruise_io_CloneSetCondition is an autogenerated conversion function.
func Convert_v1alpha1_CloneSetCondition_To_apps_kruise_io_CloneSetCondition(in *v1alpha1.CloneSetCondition, out *appskruiseio.CloneSetCondition, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloneSetCondition_To_apps_kruise_io_CloneSetCondition(in, out, s)
}

func autoConvert_apps_kruise_io_CloneSetCondition_To_v1alpha1_CloneSetCondition(in *appskruiseio.CloneSetCondition, out *v1alpha1.CloneSetCondition, s conversion.Scope) error {
	out.Type = v1alpha1.CloneSetConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_apps_kruise_io_CloneSetCondition_To_v1alpha1_CloneSetCondition is an autogenerated conversion function.
func Convert_apps_kruise_io_CloneSetCondition_To_v1alpha1_CloneSetCondition(in *appskruiseio.CloneSetCondition, out *v1alpha1.CloneSetCondition, s conversion.Scope) error {
	return autoConvert_apps_kruise_io_CloneSetCondition_To_v1alpha1_CloneSetCondition(in, out, s)
}

func autoConvert_v1alpha1_CloneSetList_To_apps_kruise_io_CloneSetList(in *v1alpha1.CloneSetList, out *appskruiseio.CloneSetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]appskruiseio.CloneSet)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_CloneSetList_To_apps_kruise_io_CloneSetList is an autogenerated conversion function.
func Convert_v1alpha1_CloneSetList_To_apps_kruise_io_CloneSetList(in *v1alpha1.CloneSetList, out *appskruiseio.CloneSetList, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloneSetList_To_apps_kruise_io_CloneSetList(in, out, s)
}

func autoConvert_apps_kruise_io_CloneSetList_To_v1alpha1_CloneSetList(in *appskruiseio.CloneSetList, out *v1alpha1.CloneSetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1alpha1.CloneSet)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_apps_kruise_io_CloneSetList_To_v1alpha1_CloneSetList is an autogenerated conversion function.
func Convert_apps_kruise_io_CloneSetList_To_v1alpha1_CloneSetList(in *appskruiseio.CloneSetList, out *v1alpha1.CloneSetList, s conversion.Scope) error {
	return autoConvert_apps_kruise_io_CloneSetList_To_v1alpha1_CloneSetList(in, out, s)
}

func autoConvert_v1alpha1_CloneSetScaleStrategy_To_apps_kruise_io_CloneSetScaleStrategy(in *v1alpha1.CloneSetScaleStrategy, out *appskruiseio.CloneSetScaleStrategy, s conversion.Scope) error {
	out.PodsToDelete = *(*[]string)(unsafe.Pointer(&in.PodsToDelete))
	out.MaxUnavailable = (*intstr.IntOrString)(unsafe.Pointer(in.MaxUnavailable))
	out.DisablePVCReuse = in.DisablePVCReuse
	return nil
}

// Convert_v1alpha1_CloneSetScaleStrategy_To_apps_kruise_io_CloneSetScaleStrategy is an autogenerated conversion function.
func Convert_v1alpha1_CloneSetScaleStrategy_To_apps_kruise_io_CloneSetScaleStrategy(in *v1alpha1.CloneSetScaleStrategy, out *appskruiseio.CloneSetScaleStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloneSetScaleStrategy_To_apps_kruise_io_CloneSetScaleStrategy(in, out, s)
}

func autoConvert_apps_kruise_io_CloneSetScaleStrategy_To_v1alpha1_CloneSetScaleStrategy(in *appskruiseio.CloneSetScaleStrategy, out *v1alpha1.CloneSetScaleStrategy, s conversion.Scope) error {
	out.PodsToDelete = *(*[]string)(unsafe.Pointer(&in.PodsToDelete))
	out.MaxUnavailable = (*intstr.IntOrString)(unsafe.Pointer(in.MaxUnavailable))
	out.DisablePVCReuse = in.DisablePVCReuse
	return nil
}

// Convert_apps_kruise_io_CloneSetScaleStrategy_To_v1alpha1_CloneSetScaleStrategy is an autogenerated conversion function.
func Convert_apps_kruise_io_CloneSetScaleStrategy_To_v1alpha1_CloneSetScaleStrategy(in *appskruiseio.CloneSetScaleStrategy, out *v1alpha1.CloneSetScaleStrategy, s conversion.Scope) error {
	return autoConvert_apps_kruise_io_CloneSetScaleStrategy_To_v1alpha1_CloneSetScaleStrategy(in, out, s)
}

func autoConvert_v1alpha1_CloneSetSpec_To_apps_kruise_io_CloneSetSpec(in *v1alpha1.CloneSetSpec, out *appskruiseio.CloneSetSpec, s conversion.Scope) error {
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
	out.Template = in.Template
	out.VolumeClaimTemplates = *(*[]v1.PersistentVolumeClaim)(unsafe.Pointer(&in.VolumeClaimTemplates))
	if err := Convert_v1alpha1_CloneSetScaleStrategy_To_apps_kruise_io_CloneSetScaleStrategy(&in.ScaleStrategy, &out.ScaleStrategy, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_CloneSetUpdateStrategy_To_apps_kruise_io_CloneSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
		return err
	}
	out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
	out.MinReadySeconds = in.MinReadySeconds
	out.Lifecycle = (*pub.Lifecycle)(unsafe.Pointer(in.Lifecycle))
	return nil
}

// Convert_v1alpha1_CloneSetSpec_To_apps_kruise_io_CloneSetSpec is an autogenerated conversion function.
func Convert_v1alpha1_CloneSetSpec_To_apps_kruise_io_CloneSetSpec(in *v1alpha1.CloneSetSpec, out *appskruiseio.CloneSetSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloneSetSpec_To_apps_kruise_io_CloneSetSpec(in, out, s)
}

func autoConvert_apps_kruise_io_CloneSetSpec_To_v1alpha1_CloneSetSpec(in *appskruiseio.CloneSetSpec, out *v1alpha1.CloneSetSpec, s conversion.Scope) error {
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	out.Selector = (*metav1.LabelSelector)(unsafe.Pointer(in.Selector))
	out.Template = in.Template
	out.VolumeClaimTemplates = *(*[]v1.PersistentVolumeClaim)(unsafe.Pointer(&in.VolumeClaimTemplates))
	if err := Convert_apps_kruise_io_CloneSetScaleStrategy_To_v1alpha1_CloneSetScaleStrategy(&in.ScaleStrategy, &out.ScaleStrategy, s); err != nil {
		return err
	}
	if err := Convert_apps_kruise_io_CloneSetUpdateStrategy_To_v1alpha1_CloneSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, s); err != nil {
		return err
	}
	out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
	out.MinReadySeconds = in.MinReadySeconds
	out.Lifecycle = (*pub.Lifecycle)(unsafe.Pointer(in.Lifecycle))
	return nil
}

// Convert_apps_kruise_io_CloneSetSpec_To_v1alpha1_CloneSetSpec is an autogenerated conversion function.
func Convert_apps_kruise_io_CloneSetSpec_To_v1alpha1_CloneSetSpec(in *appskruiseio.CloneSetSpec, out *v1alpha1.CloneSetSpec, s conversion.Scope) error {
	return autoConvert_apps_kruise_io_CloneSetSpec_To_v1alpha1_CloneSetSpec(in, out, s)
}

func autoConvert_v1alpha1_CloneSetStatus_To_apps_kruise_io_CloneSetStatus(in *v1alpha1.CloneSetStatus, out *appskruiseio.CloneSetStatus, s conversion.Scope) error {
	out.ObservedGeneration = in.ObservedGeneration
	out.Replicas = in.Replicas
	out.ReadyReplicas = in.ReadyReplicas
	out.AvailableReplicas = in.AvailableReplicas
	out.UpdatedReplicas = in.UpdatedReplicas
	out.UpdatedReadyReplicas = in.UpdatedReadyReplicas
	out.UpdatedAvailableReplicas = in.UpdatedAvailableReplicas
	out.ExpectedUpdatedReplicas = in.ExpectedUpdatedReplicas
	out.UpdateRevision = in.UpdateRevision
	out.CurrentRevision = in.CurrentRevision
	out.CollisionCount = (*int32)(unsafe.Pointer(in.CollisionCount))
	out.Conditions = *(*[]appskruiseio.CloneSetCondition)(unsafe.Pointer(&in.Conditions))
	out.LabelSelector = in.LabelSelector
	return nil
}

// Convert_v1alpha1_CloneSetStatus_To_apps_kruise_io_CloneSetStatus is an autogenerated conversion function.
func Convert_v1alpha1_CloneSetStatus_To_apps_kruise_io_CloneSetStatus(in *v1alpha1.CloneSetStatus, out *appskruiseio.CloneSetStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloneSetStatus_To_apps_kruise_io_CloneSetStatus(in, out, s)
}

func autoConvert_apps_kruise_io_CloneSetStatus_To_v1alpha1_CloneSetStatus(in *appskruiseio.CloneSetStatus, out *v1alpha1.CloneSetStatus, s conversion.Scope) error {
	out.ObservedGeneration = in.ObservedGeneration
	out.Replicas = in.Replicas
	out.ReadyReplicas = in.ReadyReplicas
	out.AvailableReplicas = in.AvailableReplicas
	out.UpdatedReplicas = in.UpdatedReplicas
	out.UpdatedReadyReplicas = in.UpdatedReadyReplicas
	out.UpdatedAvailableReplicas = in.UpdatedAvailableReplicas
	out.ExpectedUpdatedReplicas = in.ExpectedUpdatedReplicas
	out.UpdateRevision = in.UpdateRevision
	out.CurrentRevision = in.CurrentRevision
	out.CollisionCount = (*int32)(unsafe.Pointer(in.CollisionCount))
	out.Conditions = *(*[]v1alpha1.CloneSetCondition)(unsafe.Pointer(&in.Conditions))
	out.LabelSelector = in.LabelSelector
	return nil
}

// Convert_apps_kruise_io_CloneSetStatus_To_v1alpha1_CloneSetStatus is an autogenerated conversion function.
func Convert_apps_kruise_io_CloneSetStatus_To_v1alpha1_CloneSetStatus(in *appskruiseio.CloneSetStatus, out *v1alpha1.CloneSetStatus, s conversion.Scope) error {
	return autoConvert_apps_kruise_io_CloneSetStatus_To_v1alpha1_CloneSetStatus(in, out, s)
}

func autoConvert_v1alpha1_CloneSetUpdateStrategy_To_apps_kruise_io_CloneSetUpdateStrategy(in *v1alpha1.CloneSetUpdateStrategy, out *appskruiseio.CloneSetUpdateStrategy, s conversion.Scope) error {
	out.Type = appskruiseio.CloneSetUpdateStrategyType(in.Type)
	out.Partition = (*intstr.IntOrString)(unsafe.Pointer(in.Partition))
	out.MaxUnavailable = (*intstr.IntOrString)(unsafe.Pointer(in.MaxUnavailable))
	out.MaxSurge = (*intstr.IntOrString)(unsafe.Pointer(in.MaxSurge))
	out.Paused = in.Paused
	out.PriorityStrategy = (*pub.UpdatePriorityStrategy)(unsafe.Pointer(in.PriorityStrategy))
	out.ScatterStrategy = *(*appskruiseio.UpdateScatterStrategy)(unsafe.Pointer(&in.ScatterStrategy))
	out.InPlaceUpdateStrategy = (*pub.InPlaceUpdateStrategy)(unsafe.Pointer(in.InPlaceUpdateStrategy))
	return nil
}

// Convert_v1alpha1_CloneSetUpdateStrategy_To_apps_kruise_io_CloneSetUpdateStrategy is an autogenerated conversion function.
func Convert_v1alpha1_CloneSetUpdateStrategy_To_apps_kruise_io_CloneSetUpdateStrategy(in *v1alpha1.CloneSetUpdateStrategy, out *appskruiseio.CloneSetUpdateStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloneSetUpdateStrategy_To_apps_kruise_io_CloneSetUpdateStrategy(in, out, s)
}

func autoConvert_apps_kruise_io_CloneSetUpdateStrategy_To_v1alpha1_CloneSetUpdateStrategy(in *appskruiseio.CloneSetUpdateStrategy, out *v1alpha1.CloneSetUpdateStrategy, s conversion.Scope) error {
	out.Type = v1alpha1.CloneSetUpdateStrategyType(in.Type)
	out.Partition = (*intstr.IntOrString)(unsafe.Pointer(in.Partition))
	out.MaxUnavailable = (*intstr.IntOrString)(unsafe.Pointer(in.MaxUnavailable))
	out.MaxSurge = (*intstr.IntOrString)(unsafe.Pointer(in.MaxSurge))
	out.Paused = in.Paused
	out.PriorityStrategy = (*pub.UpdatePriorityStrategy)(unsafe.Pointer(in.PriorityStrategy))
	out.ScatterStrategy = *(*v1alpha1.UpdateScatterStrategy)(unsafe.Pointer(&in.ScatterStrategy))
	out.InPlaceUpdateStrategy = (*pub.InPlaceUpdateStrategy)(unsafe.Pointer(in.InPlaceUpdateStrategy))
	return nil
}

// Convert_apps_kruise_io_CloneSetUpdateStrategy_To_v1alpha1_CloneSetUpdateStrategy is an autogenerated conversion function.
func Convert_apps_kruise_io_CloneSetUpdateStrategy_To_v1alpha1_CloneSetUpdateStrategy(in *appskruiseio.CloneSetUpdateStrategy, out *v1alpha1.CloneSetUpdateStrategy, s conversion.Scope) error {
	return autoConvert_apps_kruise_io_CloneSetUpdateStrategy_To_v1alpha1_CloneSetUpdateStrategy(in, out, s)
}

func autoConvert_v1alpha1_UpdateScatterTerm_To_apps_kruise_io_UpdateScatterTerm(in *v1alpha1.UpdateScatterTerm, out *appskruiseio.UpdateScatterTerm, s conversion.Scope) error {
	out.Key = in.Key
	out.Value = in.Value
	return nil
}

// Convert_v1alpha1_UpdateScatterTerm_To_apps_kruise_io_UpdateScatterTerm is an autogenerated conversion function.
func Convert_v1alpha1_UpdateScatterTerm_To_apps_kruise_io_UpdateScatterTerm(in *v1alpha1.UpdateScatterTerm, out *appskruiseio.UpdateScatterTerm, s conversion.Scope) error {
	return autoConvert_v1alpha1_UpdateScatterTerm_To_apps_kruise_io_UpdateScatterTerm(in, out, s)
}

func autoConvert_apps_kruise_io_UpdateScatterTerm_To_v1alpha1_UpdateScatterTerm(in *appskruiseio.UpdateScatterTerm, out *v1alpha1.UpdateScatterTerm, s conversion.Scope) error {
	out.Key = in.Key
	out.Value = in.Value
	return nil
}

// Convert_apps_kruise_io_UpdateScatterTerm_To_v1alpha1_UpdateScatterTerm is an autogenerated conversion function.
func Convert_apps_kruise_io_UpdateScatterTerm_To_v1alpha1_UpdateScatterTerm(in *appskruiseio.UpdateScatterTerm, out *v1alpha1.UpdateScatterTerm, s conversion.Scope) error {
	return autoConvert_apps_kruise_io_UpdateScatterTerm_To_v1alpha1_UpdateScatterTerm(in, out, s)
}
