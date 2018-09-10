package v1

import (
	v1 "github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// UserProjectBindingLister helps list UserProjectBindings.
type UserProjectBindingLister interface {
	// List lists all UserProjectBindings in the indexer.
	List(selector labels.Selector) (ret []*v1.UserProjectBinding, err error)
	// Get retrieves the UserProjectBinding from the index for a given name.
	Get(name string) (*v1.UserProjectBinding, error)
	UserProjectBindingListerExpansion
}

// userProjectBindingLister implements the UserProjectBindingLister interface.
type userProjectBindingLister struct {
	indexer cache.Indexer
}

// NewUserProjectBindingLister returns a new UserProjectBindingLister.
func NewUserProjectBindingLister(indexer cache.Indexer) UserProjectBindingLister {
	return &userProjectBindingLister{indexer: indexer}
}

// List lists all UserProjectBindings in the indexer.
func (s *userProjectBindingLister) List(selector labels.Selector) (ret []*v1.UserProjectBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.UserProjectBinding))
	})
	return ret, err
}

// Get retrieves the UserProjectBinding from the index for a given name.
func (s *userProjectBindingLister) Get(name string) (*v1.UserProjectBinding, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("userprojectbinding"), name)
	}
	return obj.(*v1.UserProjectBinding), nil
}
