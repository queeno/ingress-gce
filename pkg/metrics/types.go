/*
Copyright 2020 The Kubernetes Authors.

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

package metrics

import (
	"k8s.io/api/networking/v1beta1"
	"k8s.io/ingress-gce/pkg/utils"
)

// IngressState defines an ingress and its associated service ports.
type IngressState struct {
	ingress      *v1beta1.Ingress
	servicePorts []utils.ServicePort
}

// NegServiceState contains all the neg usage associated with one service
type NegServiceState struct {
	// standaloneNeg is the count of standalone NEG
	StandaloneNeg int
	// ingressNeg is the count of NEGs created for ingress
	IngressNeg int
	// asmNeg is the count of NEGs created for ASM
	AsmNeg int
}

// IngressMetricsCollector is an interface to update/delete ingress states in the cache
// that is used for computing ingress usage metrics.
type IngressMetricsCollector interface {
	// SetIngress adds/updates ingress state for given ingress key.
	SetIngress(ingKey string, ing IngressState)
	// DeleteIngress removes the given ingress key.
	DeleteIngress(ingKey string)
}

// NegMetricsCollector is an interface to update/delete Neg states in the cache
// that is used for computing neg usage metrics.
type NegMetricsCollector interface {
	// SetNegService adds/updates neg state for given service key.
	SetNegService(svcKey string, negState NegServiceState)
	// DeleteNegService removes the given service key.
	DeleteNegService(svcKey string)
}
