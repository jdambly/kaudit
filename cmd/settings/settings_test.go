package settings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLabelSelector_HappyPath(t *testing.T) {
	s := &Settings{}
	selector := "key1=value1,key2=value2"
	expected := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	labels, err := s.ParseLabelSelector(selector)

	assert.NoError(t, err, "Unexpected error")
	assert.Equal(t, expected, labels, "The labels do not match the expected labels")
}

func TestParseLabelSelector_EmptyString(t *testing.T) {
	s := &Settings{}
	selector := ""
	expected := map[string]string{}
	labels, err := s.ParseLabelSelector(selector)

	assert.Error(t, err, "label selector is empty")
	assert.Equal(t, expected, labels, "The labels do not match the expected labels")
}

func TestParseLabelSelector_InvalidFormat(t *testing.T) {
	s := &Settings{}
	selector := "key1=value1,key2"

	_, err := s.ParseLabelSelector(selector)

	assert.NoError(t, err, "Expected error")
}
