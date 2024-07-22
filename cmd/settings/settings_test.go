package settings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListOpts_HappyPath(t *testing.T) {
	s := &Settings{
		NodeName:      "test-node",
		Namespace:     "test-namespace",
		LabelSelector: "test-label=test-value,extra-label=extra-value",
	}
	listOpts := s.ListOpts()
	assert.Equal(t, "test-label=test-value,extra-label=extra-value", listOpts.LabelSelector)
	assert.Equal(t, "spec.nodeName=test-node", listOpts.FieldSelector)
}

func TestListOpts_Empty(t *testing.T) {
	s := &Settings{
		NodeName:      "",
		Namespace:     "",
		LabelSelector: "",
	}
	listOpts := s.ListOpts()
	assert.Equal(t, "", listOpts.LabelSelector)
	assert.Equal(t, "spec.nodeName=", listOpts.FieldSelector)
}

func TestListOpts_BadLabelSelector(t *testing.T) {
	s := &Settings{
		NodeName:      "test-node",
		Namespace:     "test-namespace",
		LabelSelector: "test-label=key,bad-label",
	}
	listOpts := s.ListOpts()
	assert.Equal(t, "test-label=key,bad-label", listOpts.LabelSelector)
	assert.Equal(t, "spec.nodeName=test-node", listOpts.FieldSelector)
}
