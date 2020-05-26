// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package webhook

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// daemonSetProcessorFactory implements the processFactory interface to mutate pod template spec in daemonsets.
type daemonSetProcessorFactory struct{}

func (pf *daemonSetProcessorFactory) kind() metav1.GroupVersionKind {
	return metav1.GroupVersionKind{
		Group:   appsv1.SchemeGroupVersion.Group,
		Version: appsv1.SchemeGroupVersion.Version,
		Kind:    "DaemonSet",
	}
}

func (pf *daemonSetProcessorFactory) newProcessor() processor {
	return &podSpecProcessorImpl{
		podSpecCallbacks: &daemonSetCallbacks{},
	}
}

func (pf *daemonSetProcessorFactory) isMutating() bool {
	return true
}

// daemonSetCallbacks implements podSpecCallbacks
type daemonSetCallbacks struct {
	daemonSet appsv1.DaemonSet
}

func (d *daemonSetCallbacks) getObject() runtime.Object {
	return &(d.daemonSet)
}

func (d *daemonSetCallbacks) getPodSpec() *corev1.PodSpec {
	return &(d.daemonSet.Spec.Template.Spec)
}

func (d *daemonSetCallbacks) getPodLabels() map[string]string {
	return d.daemonSet.Spec.Template.Labels
}
