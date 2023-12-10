package env

import (
	"os"
	"reflect"
	"testing"

	"github.com/joho/godotenv"
)

func TestNewReader(t *testing.T) {
	tests := []struct {
		name string
		want *Reader
	}{
		{
			name: "define",
			want: &Reader{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReader(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_Read(t *testing.T) {
	_ = godotenv.Load("../../.env")
	tests := []struct {
		stage   string
		name    string
		want    *Server
		want1   *DB
		want2   *JWT
		wantErr bool
	}{
		{
			stage:   "PROD",
			name:    "invalid STAGE",
			want:    nil,
			want1:   nil,
			want2:   nil,
			wantErr: true,
		},
		{
			stage: "DEV",
			name:  "valid STAGE",
			want: &Server{
				host: os.Getenv("SERVER_HOST"),
			},
			want1: &DB{
				host:     os.Getenv("DB_HOST"),
				user:     os.Getenv("DB_USER"),
				password: os.Getenv("DB_PASSWORD"),
				name:     os.Getenv("DB_NAME"),
			},
			want2: &JWT{
				secret: os.Getenv("JWT_SECRET"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("STAGE", tt.stage)
			r := NewReader()
			got, got1, got2, err := r.Read()
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reader.Read() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Reader.Read() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Reader.Read() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
