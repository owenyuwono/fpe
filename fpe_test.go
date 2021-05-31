package fpe

import (
	"math/big"
	"reflect"
	"testing"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		modulus int64
		subject int64
		key     string
		tweak   string
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{"success - default", args{10001, 1, "my-secret-key", "my-non-secret-tweak"}, big.NewInt(5011), false},
		{"success - different key", args{10001, 1, "different-key", "should-yield-different-result"}, big.NewInt(8779), false},
		{"success - default with subject 2", args{10001, 2, "my-secret-key", "my-non-secret-tweak"}, big.NewInt(9102), false},
		{"success - high modulus", args{99999999, 1, "my-secret-key", "my-non-secret-tweak"}, big.NewInt(88925566), false},
		{"success - high modulus and high subject", args{99999999, 99999998, "my-secret-key", "my-non-secret-tweak"}, big.NewInt(9895125), false},
		{"success - modulus value equal to subject", args{1000, 1000, "my-secret-key", "my-non-secret-tweak"}, big.NewInt(861), false},
		{"success - subject exceeds modulus", args{9999, 10000, "my-secret-key", "my-non-secret-tweak"}, big.NewInt(3655), false},
		{"success - prime modulus", args{9199, 20, "my-secret-key", "my-non-secret-tweak"}, big.NewInt(1097), false},
		{"success - negative modulus", args{-1, 1, "my-secret-key", "my-non-secret-tweak"}, nil, true},
		{"success - negative subject", args{10001, -1, "my-secret-key", "my-non-secret-tweak"}, big.NewInt(5011), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrypt(tt.args.modulus, tt.args.subject, tt.args.key, tt.args.tweak, 3)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
			if got.Cmp(big.NewInt(tt.args.modulus)) == 1 {
				t.Error("Encrypt() exceeds max range")
			}
			if got.Cmp(big.NewInt(0)) == -1 {
				t.Error("Encrypt() is negative")
			}
		})
	}
}
