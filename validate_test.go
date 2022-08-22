package commonsvalidation

import (
	"reflect"
	"testing"

	commonserrors "github.com/finacore/commons-errors"
)

type User struct {
	Name    string `validate:"required,min=3,max=256"`
	Surname string `validate:"required,min=3,max=256"`
	Email   string `validate:"omitempty,email"`
}

func TestValidateModel(t *testing.T) {
	type args struct {
		model interface{}
	}
	tests := []struct {
		name string
		args args
		want []commonserrors.ValidationError
	}{
		{
			name: "nil-model",
			args: args{model: nil},
			want: nil,
		},
		{
			name: "all-in",
			args: args{model: User{Name: "João", Surname: "da Silva", Email: "dasilva@gmail.com"}},
			want: nil,
		},
		{
			name: "no-name",
			args: args{model: User{Surname: "da Silva", Email: "dasilva@gmail.com"}},
			want: []commonserrors.ValidationError{{Field: "User.Name", Message: "Name is a required field"}},
		},
		{
			name: "no-surname",
			args: args{model: User{Name: "João", Email: "dasilva@gmail.com"}},
			want: []commonserrors.ValidationError{{Field: "User.Surname", Message: "Surname is a required field"}},
		},
		{
			name: "no-email",
			args: args{model: User{Name: "João", Surname: "da Silva"}},
			want: nil,
		},
		{
			name: "invalid-email",
			args: args{model: User{Name: "João", Surname: "da Silva", Email: "not-email"}},
			want: []commonserrors.ValidationError{{Field: "User.Email", Message: "Email must be a valid email address"}},
		},
		{
			name: "invalid-email-&-short-name",
			args: args{model: User{Name: "Jo", Surname: "da Silva", Email: "not-email"}},
			want: []commonserrors.ValidationError{
				{
					Field:   "User.Name",
					Message: "Name must be at least 3 characters in length",
				},
				{
					Field:   "User.Email",
					Message: "Email must be a valid email address",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateModel(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
