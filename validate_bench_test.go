package commonsvalidation

import (
	"reflect"
	"testing"

	commonserrors "github.com/finacore/commons-errors"
)

func BenchmarkValidateModel(b *testing.B) {
	type args struct {
		model interface{}
	}
	tests := []struct {
		name string
		args args
		want []commonserrors.ValidationError
	}{
		{
			name: "Errorless",
			args: args{model: User{Name: "João", Surname: "da Silva", Email: "dasilva@gmail.com"}},
			want: nil,
		},
		{
			name: "Error",
			args: args{model: User{Surname: "da Silva", Email: "dasilva@gmail.com"}},
			want: []commonserrors.ValidationError{{Field: "User.Name", Message: "Name is a required field"}},
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			if got := ValidateModel(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				b.Errorf("ValidateModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
