package env

import (
	"reflect"
	"testing"
)

func TestNewStage(t *testing.T) {
	stg := "DEV"
	invalidStg := "PRODUCTION" // is invalid, defined as PROD in this project
	type args struct {
		stg string
	}
	tests := []struct {
		name    string
		args    args
		want    *Stage
		wantErr bool
	}{
		{
			name: "invalid",
			args: args{
				stg: invalidStg,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid",
			args: args{
				stg: stg,
			},
			want: &Stage{
				stg: stg,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStage(tt.args.stg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStage_isDev(t *testing.T) {
	type fields struct {
		stg string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "true",
			fields: fields{
				stg: "DEV",
			},
			want: true,
		},
		{
			name: "false",
			fields: fields{
				stg: "PROD",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stage{
				stg: tt.fields.stg,
			}
			if got := s.isDev(); got != tt.want {
				t.Errorf("Stage.isDev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStage_isTest(t *testing.T) {
	type fields struct {
		stg string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "true",
			fields: fields{
				stg: "TEST",
			},
			want: true,
		},
		{
			name: "false",
			fields: fields{
				stg: "PROD",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stage{
				stg: tt.fields.stg,
			}
			if got := s.isTest(); got != tt.want {
				t.Errorf("Stage.isTest() = %v, want %v", got, tt.want)
			}
		})
	}
}
