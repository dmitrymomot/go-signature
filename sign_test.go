package signature

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	var bdata = []byte("data")
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"success", args{bdata}, "eyJwIjoiWkdGMFlRPT0ifS45YzIxMDlmY2I1NDQ2NTE2MTQ0NTUyNjczMGRlNmE0YjI2YTkwOTgz", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew256(t *testing.T) {
	var bdata = []byte("data")
	var token = "eyJwIjoiWkdGMFlRPT0ifS42MTJmYzE4MGFlNGNlYWViNTdhYjc4NTBkNDVkY2FiMzE2NGRjMmJhMzYzZTQ3MWRmOTc5MmQ4NzdhMTFiMzEz"
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"success", args{bdata}, token, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New256(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("New256() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("New256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTemporary(t *testing.T) {
	var bdata = []byte("data")
	var token = "eyJwIjoiWkdGMFlRPT0ifS45YzIxMDlmY2I1NDQ2NTE2MTQ0NTUyNjczMGRlNmE0YjI2YTkwOTgz"
	type args struct {
		data interface{}
		ttl  int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"success", args{bdata, 0}, token, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTemporary(tt.args.data, tt.args.ttl)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTemporary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewTemporary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew256Temporary(t *testing.T) {
	var bdata = []byte("data")
	var token = "eyJwIjoiWkdGMFlRPT0ifS42MTJmYzE4MGFlNGNlYWViNTdhYjc4NTBkNDVkY2FiMzE2NGRjMmJhMzYzZTQ3MWRmOTc5MmQ4NzdhMTFiMzEz"
	type args struct {
		data interface{}
		ttl  int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"success", args{bdata, 0}, token, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New256Temporary(tt.args.data, tt.args.ttl)
			if (err != nil) != tt.wantErr {
				t.Errorf("New256Temporary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("New256Temporary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	var token = "eyJwIjoiWkdGMFlRPT0ifS45YzIxMDlmY2I1NDQ2NTE2MTQ0NTUyNjczMGRlNmE0YjI2YTkwOTgz"
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"success", args{token}, "ZGF0YQ==", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse256(t *testing.T) {
	var token = "eyJwIjoiWkdGMFlRPT0ifS42MTJmYzE4MGFlNGNlYWViNTdhYjc4NTBkNDVkY2FiMzE2NGRjMmJhMzYzZTQ3MWRmOTc5MmQ4NzdhMTFiMzEz"
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"success", args{token}, "ZGF0YQ==", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse256(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse256() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newToken(t *testing.T) {
	var bdata = []byte("data")
	var token = "eyJwIjoiWkdGMFlRPT0ifS45YzIxMDlmY2I1NDQ2NTE2MTQ0NTUyNjczMGRlNmE0YjI2YTkwOTgz"
	type args struct {
		data interface{}
		ttl  int64
		fn   hmacFunc
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"success", args{bdata, 0, calculateHmac}, token, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newToken(tt.args.data, tt.args.ttl, tt.args.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("newToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("newToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseToken(t *testing.T) {
	testSigningKey := []byte("secret")
	token := "eyJwIjoiWkdGMFlRPT0ifS45YzIxMDlmY2I1NDQ2NTE2MTQ0NTUyNjczMGRlNmE0YjI2YTkwOTgz"
	tokenTTL, _ := NewTemporary("data", 300)
	tokenTTLExp, _ := NewTemporary("data", 1)
	invalidToken := "yJwIjoiWkdGMFlRPT0ifS4xNzBlNGU5Zjk3MWE5NzM3YzYwOWJmNmIwMmFiODdlMGIwMTIyZTcz"
	b := []byte("invalid json string")
	ib := base64Encode([]byte(fmt.Sprintf("%s1.%s", string(b), calculateHmac(b, testSigningKey))))
	type args struct {
		token string
		fn    hmacFunc
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"success", args{token, calculateHmac}, "ZGF0YQ==", false},
		{"success with ttl", args{tokenTTL, calculateHmac}, "data", false},
		{"invalid token", args{invalidToken, calculateHmac256}, nil, true},
		{"invalid base64 string", args{"121212===", calculateHmac256}, nil, true},
		{"validation error", args{ib, calculateHmac}, nil, true},
		// TODO: {"unmarshaling error", args{ ib, calculateHmac}, nil, false},
		{"expired", args{tokenTTLExp, calculateHmac}, nil, true},
	}
	for _, tt := range tests {
		time.Sleep(200 * time.Millisecond)
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseToken(tt.args.token, tt.args.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
