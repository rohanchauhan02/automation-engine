package config

import (
	"reflect"
	"testing"
)

func TestNewImmutableConfig(t *testing.T) {
	tests := []struct {
		name string
		want ImmutableConfigInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewImmutableConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewImmutableConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
