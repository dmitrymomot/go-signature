package signature

import (
	"encoding/json"
	"testing"
)

func Test_calculateHmac(t *testing.T) {
	var testSigningKey = []byte("secret")
	type args struct {
		securedInput []byte
		key          []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"calculateHmac", args{[]byte("data"), testSigningKey}, "9818e3306ba5ac267b5f2679fe4abd37e6cd7b54"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateHmac(tt.args.securedInput, tt.args.key); got != tt.want {
				t.Errorf("calculateHmac() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateHmac256(t *testing.T) {
	var testSigningKey = []byte("secret")
	type args struct {
		securedInput []byte
		key          []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"calculateHmac256", args{[]byte("data"), testSigningKey}, "1b2c16b75bd2a870c114153ccda5bcfca63314bc722fa160d690de133ccbb9db"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateHmac256(tt.args.securedInput, tt.args.key); got != tt.want {
				t.Errorf("calculateHmac256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateHmac(t *testing.T) {
	var testSigningKey = []byte("secret")
	var testInvalidSigningKey = []byte("invalid secret")
	claims := &tokenClaims{Payload: []byte("data")}
	b, _ := json.Marshal(claims)
	s1 := []byte(calculateHmac(b, testSigningKey))
	s256 := []byte(calculateHmac256(b, testSigningKey))

	type args struct {
		securedInput []byte
		sign         []byte
		key          []byte
		fn           hmacFunc
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"calculateHmac", args{b, s1, testSigningKey, calculateHmac}, false},
		{"calculateHmac256", args{b, s256, testSigningKey, calculateHmac256}, false},
		{"invalid signature", args{b, s256, testInvalidSigningKey, calculateHmac256}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateHmac(tt.args.securedInput, tt.args.sign, tt.args.key, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("validateHmac() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
