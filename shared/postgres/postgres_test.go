package postgres

import (
	"reflect"
	"testing"

	_ "github.com/lib/pq"
	"github.com/rohanchauhan02/automation-engine/shared/config"
)

func TestNewPostgres(t *testing.T) {
	type args struct {
		sharedConfig config.ImmutableConfigInterface
	}
	tests := []struct {
		name string
		args args
		want PostgresInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPostgres(tt.args.sharedConfig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostgres() = %v, want %v", got, tt.want)
			}
		})
	}
}
