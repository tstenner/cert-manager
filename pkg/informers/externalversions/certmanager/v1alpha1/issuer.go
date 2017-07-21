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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	certmanager_v1alpha1 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha1"
	client "github.com/jetstack/cert-manager/pkg/client"
	internalinterfaces "github.com/jetstack/cert-manager/pkg/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/jetstack/cert-manager/pkg/listers/certmanager/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// IssuerInformer provides access to a shared informer and lister for
// Issuers.
type IssuerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.IssuerLister
}

type issuerInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newIssuerInformer(client client.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.CertmanagerV1alpha1().Issuers(v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.CertmanagerV1alpha1().Issuers(v1.NamespaceAll).Watch(options)
			},
		},
		&certmanager_v1alpha1.Issuer{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *issuerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&certmanager_v1alpha1.Issuer{}, newIssuerInformer)
}

func (f *issuerInformer) Lister() v1alpha1.IssuerLister {
	return v1alpha1.NewIssuerLister(f.Informer().GetIndexer())
}
