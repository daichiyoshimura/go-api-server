package value

import (
	"reflect"
	"testing"
)

func TestNewAccountName(t *testing.T) {
	name := "JohnSmith"

	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *AccountName
		wantErr bool
	}{
		{
			name: "empty",
			args: args{
				name: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "too long",
			args: args{
				name: "1234567890123456789012345678901234567890123456789012345678901234567890",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "forbidden char",
			args: args{
				name: "*",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid",
			args: args{
				name: name,
			},
			want: &AccountName{
				name: name,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAccountName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAccountName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccountName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountName_Value(t *testing.T) {
	name := "JohnSmith"

	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "define",
			fields: fields{
				name: name,
			},
			want: name,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AccountName{
				name: tt.fields.name,
			}
			if got := a.Value(); got != tt.want {
				t.Errorf("AccountName.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
