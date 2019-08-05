package signature

import (
	"reflect"
	"testing"
)

func Test_base64Decode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{"success", args{"ZGF0YQ"}, []byte("data"), false},
		{"error", args{"ZGF0YQ==="}, []byte("data"), true},
		{"error", args{"ZGF0YQ="}, []byte("data"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := base64Decode(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("base64Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("base64Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base64Encode(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"success", args{[]byte("data")}, "ZGF0YQ"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := base64Encode(tt.args.data); got != tt.want {
				t.Errorf("base64Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
