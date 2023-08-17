package cfg_test

import (
	"testing"

	"github.com/xuender/kit/cfg"
)

// nolint: paralleltest
func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		arg     string
		want    cfg.Cipher
		wantErr bool
	}{
		{"err1", "aa(a)", cfg.AES, true},
		{"err2", "AES{a}", cfg.AES, true},
		{"err3", "AES()", cfg.AES, true},
		{"paa1", "AES(A/43wTj2AVQboZZ0lNMqbw==)", cfg.AES, false},
		{"paa2", "AES[A/43wTj2AVQboZZ0lNMqbw==]", cfg.AES, false},
		{"paa3", "DES(A/43wTj2AVQboZZ0lNMqbw==)", cfg.DES, false},
		{"paa4", "DES[A/43wTj2AVQboZZ0lNMqbw==]", cfg.DES, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, got, err := cfg.Parse(test.arg)

			if (err != nil) != test.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, test.wantErr)

				return
			}

			if got != test.want {
				t.Errorf("Parse() got1 = %v, want %v", got, test.want)
			}
		})
	}
}
