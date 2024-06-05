package kubeapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestGetPodUIDList_HappyPath(t *testing.T) {
	clientset := fake.NewSimpleClientset()
	namespace := "test-namespace"
	_, _ = clientset.CoreV1().Pods(namespace).Create(context.Background(), &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-pod-1",
			UID:  "test-uid-1",
			Labels: map[string]string{
				"test-label":  "test-value",
				"extra-label": "extra-value",
			},
		},
		Spec: v1.PodSpec{
			NodeName: "test-node",
		},
	}, metav1.CreateOptions{})
	_, _ = clientset.CoreV1().Pods(namespace).Create(context.Background(), &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-pod-2",
			UID:  "test-uid-2",
			Labels: map[string]string{
				"test-label":  "test-value",
				"extra-label": "extra-value",
			},
		},
		Spec: v1.PodSpec{
			NodeName: "test-node",
		},
	}, metav1.CreateOptions{})
	listOpts := metav1.ListOptions{
		LabelSelector: "test-label=test-value,extra-label=extra-value",
		FieldSelector: "spec.nodeName=test-node",
	}
	uids, err := GetPodUIDList(clientset, namespace, listOpts)

	assert.NoError(t, err)
	assert.ElementsMatch(t, []string{"test-uid-1", "test-uid-2"}, uids)
}

func TestGetPodUIDList_NoMatchingPods(t *testing.T) {
	clientset := fake.NewSimpleClientset()
	namespace := "test-namespace"
	listOpts := metav1.ListOptions{
		LabelSelector: "test-label=test-value",
		FieldSelector: "spec.nodeName=test-node",
	}

	uids, err := GetPodUIDList(clientset, namespace, listOpts)

	assert.NoError(t, err)
	assert.Empty(t, uids)
}
