package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddAppendix(t *testing.T) {
	bufferStdout := &bytes.Buffer{}
	contextKey := struct{}{}
	key := "key"
	value := "value"
	logger := New("name", DEBUG, Output(bufferStdout), AddAppendix(NewContextAppendix(contextKey, key)))
	ctx := context.WithValue(context.Background(), contextKey, value)
	logger.InfoContext(ctx, "message", map[string]interface{}{})
	var v struct {
		Extra map[string]string `json:"extra"`
	}
	err := json.Unmarshal(bufferStdout.Bytes(), &v)
	assert.NoError(t, err)
	assert.Equal(t, value, v.Extra[key])
}
