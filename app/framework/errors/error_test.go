package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
)

func TestIsBusinessError(t *testing.T) {
	patterns := map[string]struct {
		err  error
		want bool
	}{
		"errがnilの場合はfalse":                                        {err: nil, want: false},
		"errがxerrors.Errorfで作られたerrの場合はfalse":                     {err: xerrors.Errorf("system error"), want: false},
		"errがBusinessErrorfで作られたerrの場合はtrue":                      {err: newBusinessError("business error"), want: true},
		"errがBusinessErrorfをwrapしたxerrors.Errorfで作られたerrの場合はtrue": {err: xerrors.Errorf("error: %w", newBusinessError("business error")), want: true},
	}

	for tn, tc := range patterns {
		t.Run(tn, func(t *testing.T) {
			assert.Equal(t, tc.want, IsBusinessError(tc.err))
		})
	}
}

func TestIsSystemError(t *testing.T) {
	patterns := map[string]struct {
		err  error
		want bool
	}{
		"errがnilの場合はfalse":                                         {err: nil, want: false},
		"errがxerrors.Errorfで作られたerrの場合はtrue":                       {err: xerrors.Errorf("system error"), want: true},
		"errがBusinessErrorfで作られたerrの場合はfalse":                      {err: newBusinessError("business error"), want: false},
		"errがBusinessErrorfをwrapしたxerrors.Errorfで作られたerrの場合はfalse": {err: xerrors.Errorf("error: %w", newBusinessError("business error")), want: false},
	}

	for tn, tc := range patterns {
		t.Run(tn, func(t *testing.T) {
			assert.Equal(t, tc.want, IsSystemError(tc.err))
		})
	}
}
