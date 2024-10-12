package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
)

func TestNewBusinessError(t *testing.T) {
	patterns := map[string]struct {
		format    string
		a         []any
		errString string
	}{
		"可変長引数がない場合":  {format: "error", a: nil, errString: "error"},
		"可変長引数が1つの場合": {format: "error %s", a: []any{"arg1"}, errString: "error arg1"},
		"可変長引数が2つの場合": {format: "error %s, %s", a: []any{"arg1", "arg2"}, errString: "error arg1, arg2"},
	}

	for tn, tc := range patterns {
		t.Run(tn, func(t *testing.T) {
			err := NewBusinessError(tc.format, tc.a...)
			assert.Error(t, err)
			assert.True(t, IsBusinessError(err))
			assert.Equal(t, tc.errString, err.Error())
		})
	}
}

func TestIsBusinessError(t *testing.T) {
	patterns := map[string]struct {
		err  error
		want bool
	}{
		"errがnilの場合はfalse":                                        {err: nil, want: false},
		"errがxerrors.Errorfで作られたerrの場合はfalse":                     {err: xerrors.Errorf("system error"), want: false},
		"errがBusinessErrorfで作られたerrの場合はtrue":                      {err: NewBusinessError("business error"), want: true},
		"errがBusinessErrorfをwrapしたxerrors.Errorfで作られたerrの場合はtrue": {err: xerrors.Errorf("error: %w", NewBusinessError("business error")), want: true},
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
		"errがBusinessErrorfで作られたerrの場合はfalse":                      {err: NewBusinessError("business error"), want: false},
		"errがBusinessErrorfをwrapしたxerrors.Errorfで作られたerrの場合はfalse": {err: xerrors.Errorf("error: %w", NewBusinessError("business error")), want: false},
	}

	for tn, tc := range patterns {
		t.Run(tn, func(t *testing.T) {
			assert.Equal(t, tc.want, IsSystemError(tc.err))
		})
	}
}
