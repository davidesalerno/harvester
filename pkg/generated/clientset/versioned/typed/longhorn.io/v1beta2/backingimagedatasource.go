/*
Copyright 2025 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1beta2

import (
	"context"

	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v1beta2 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// BackingImageDataSourcesGetter has a method to return a BackingImageDataSourceInterface.
// A group's client should implement this interface.
type BackingImageDataSourcesGetter interface {
	BackingImageDataSources(namespace string) BackingImageDataSourceInterface
}

// BackingImageDataSourceInterface has methods to work with BackingImageDataSource resources.
type BackingImageDataSourceInterface interface {
	Create(ctx context.Context, backingImageDataSource *v1beta2.BackingImageDataSource, opts v1.CreateOptions) (*v1beta2.BackingImageDataSource, error)
	Update(ctx context.Context, backingImageDataSource *v1beta2.BackingImageDataSource, opts v1.UpdateOptions) (*v1beta2.BackingImageDataSource, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, backingImageDataSource *v1beta2.BackingImageDataSource, opts v1.UpdateOptions) (*v1beta2.BackingImageDataSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta2.BackingImageDataSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta2.BackingImageDataSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.BackingImageDataSource, err error)
	BackingImageDataSourceExpansion
}

// backingImageDataSources implements BackingImageDataSourceInterface
type backingImageDataSources struct {
	*gentype.ClientWithList[*v1beta2.BackingImageDataSource, *v1beta2.BackingImageDataSourceList]
}

// newBackingImageDataSources returns a BackingImageDataSources
func newBackingImageDataSources(c *LonghornV1beta2Client, namespace string) *backingImageDataSources {
	return &backingImageDataSources{
		gentype.NewClientWithList[*v1beta2.BackingImageDataSource, *v1beta2.BackingImageDataSourceList](
			"backingimagedatasources",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1beta2.BackingImageDataSource { return &v1beta2.BackingImageDataSource{} },
			func() *v1beta2.BackingImageDataSourceList { return &v1beta2.BackingImageDataSourceList{} }),
	}
}
