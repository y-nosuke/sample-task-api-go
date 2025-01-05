package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newBusinessError(t *testing.T) {
	patterns := map[string]struct {
		errorCode ErrorCode
		format    string
		a         []any
		errString string
	}{
		"可変長引数がない場合":  {errorCode: BadRequest, format: "error", a: nil, errString: "error"},
		"可変長引数が1つの場合": {errorCode: BadRequest, format: "error %s", a: []any{"arg1"}, errString: "error arg1"},
		"可変長引数が2つの場合": {errorCode: BadRequest, format: "error %s, %s", a: []any{"arg1", "arg2"}, errString: "error arg1, arg2"},
	}

	for tn, tc := range patterns {
		t.Run(tn, func(t *testing.T) {
			err := newBusinessError(tc.errorCode, tc.format, tc.a...)
			assert.Error(t, err)
			assert.True(t, IsBusinessError(err))
			assert.Equal(t, tc.errString, err.Error())
		})
	}
}
