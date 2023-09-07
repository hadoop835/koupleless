package utils

import (
	"github.com/sofastack/sofa-serverless/internal/constants/finalizer"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"testing"
	"time"
)

func TestAddNotExistedFinalizer(t *testing.T) {
	meta := &metav1.ObjectMeta{}
	assert.True(t, AddFinalizer(meta, finalizer.ModuleReplicaSetExistedFinalizer))
}

func TestHasNotExistedFinalizer(t *testing.T) {
	meta := &metav1.ObjectMeta{}
	assert.False(t, HasFinalizer(meta, finalizer.ModuleReplicaSetExistedFinalizer))
}

func TestRemoveNotExistedFinalizer(t *testing.T) {
	meta := &metav1.ObjectMeta{}
	assert.False(t, RemoveFinalizer(meta, finalizer.ModuleReplicaSetExistedFinalizer))
}

func TestAddExistedFinalizer(t *testing.T) {
	meta := &metav1.ObjectMeta{}
	meta.Finalizers = []string{finalizer.ModuleReplicaSetExistedFinalizer}
	assert.False(t, AddFinalizer(meta, finalizer.ModuleReplicaSetExistedFinalizer))
}

func TestHasExistedFinalizer(t *testing.T) {
	meta := &metav1.ObjectMeta{}
	meta.Finalizers = []string{finalizer.ModuleReplicaSetExistedFinalizer}
	assert.True(t, HasFinalizer(meta, finalizer.ModuleReplicaSetExistedFinalizer))
}

func TestRemoveExistedFinalizer(t *testing.T) {
	meta := &metav1.ObjectMeta{}
	meta.Finalizers = []string{finalizer.ModuleReplicaSetExistedFinalizer}
	assert.True(t, RemoveFinalizer(meta, finalizer.ModuleReplicaSetExistedFinalizer))
}

func TestKeyEqual(t *testing.T) {
	request := ctrl.Request{NamespacedName: types.NamespacedName{Namespace: "default", Name: "test"}}
	key := Key(request)
	assert.Equal(t, "default/test", key)
}

func TestGetNextReconcileTime(t *testing.T) {
	now := time.Now()
	reconcileTime := GetNextReconcileTime(now)
	assert.Equal(t, "10s", reconcileTime.String())

	m, _ := time.ParseDuration("-11m")
	reconcileTime2 := GetNextReconcileTime(now.Add(m))
	assert.Equal(t, "1m0s", reconcileTime2.String())

	m31, _ := time.ParseDuration("-31m")
	reconcileTime3 := GetNextReconcileTime(now.Add(m31))
	assert.Equal(t, "5m0s", reconcileTime3.String())

	h1, _ := time.ParseDuration("-61m")
	reconcileTime4 := GetNextReconcileTime(now.Add(h1))
	assert.Equal(t, "10m0s", reconcileTime4.String())
}