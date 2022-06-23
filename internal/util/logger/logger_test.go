package logger

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type format struct {
	Name     string `json:"name"`
	Message  string `json:"message"`
	Severity string `json:"severity"`
	File     string `json:"file"`
}

// ログ出力が正しく扱えているかをテスト
func TestLogger_LogToOutput(t *testing.T) {
	bufferStdout := &bytes.Buffer{}
	bufferStderr := &bytes.Buffer{}
	name := "test"
	message := "debug log"
	logger := New(name, DEBUG, Output(bufferStdout), ErrorOutput(bufferStderr))
	logger.Debug(message, map[string]interface{}{
		"params": map[string]interface{}{
			"page": 100,
		},
	})
	v := struct {
		format
		Context struct {
			Params struct {
				Page int `json:"page"`
			} `json:"params"`
		} `json:"context"`
	}{}
	err := json.Unmarshal(bufferStdout.Bytes(), &v)
	assert.Nil(t, err)
	assert.Equal(t, name, v.Name)
	assert.Equal(t, "DEBUG", v.Severity)
	assert.Equal(t, message, v.Message)
	assert.Equal(t, 100, v.Context.Params.Page)
}

// ログ出力が正しく扱えているかをテスト
func TestLogger_LogToErrorOutput(t *testing.T) {
	bufferStdout := &bytes.Buffer{}
	bufferStderr := &bytes.Buffer{}
	message := "waring log"

	logger := New("name", DEBUG, Output(bufferStdout), ErrorOutput(bufferStderr))
	logger.Warning(message, map[string]interface{}{
		"params": map[string]interface{}{
			"page": 100,
		},
	})
	v := struct {
		format
		Context struct {
			Params struct {
				Page int `json:"page"`
			} `json:"params"`
		} `json:"context"`
	}{}
	err := json.Unmarshal(bufferStderr.Bytes(), &v)
	assert.Nil(t, err)
	assert.Equal(t, "WARNING", v.Severity)
	assert.Equal(t, message, v.Message)
	assert.Equal(t, 100, v.Context.Params.Page)
}

func TestLogger_Info(t *testing.T) {
	bufferStdout := &bytes.Buffer{}
	bufferStderr := &bytes.Buffer{}
	message := "info log"
	logger := New("name", INFO, Output(bufferStdout), ErrorOutput(bufferStderr))
	logger.Debug(message, map[string]interface{}{})
	assert.Equal(t, "", bufferStdout.String())

	logger.Info(message, map[string]interface{}{})
	v := format{}
	err := json.Unmarshal(bufferStdout.Bytes(), &v)
	assert.Nil(t, err)
	assert.Equal(t, "INFO", v.Severity)
	assert.Equal(t, message, v.Message)
}

func TestLogger_Warning(t *testing.T) {
	message := "warning log"
	bufferStdout := &bytes.Buffer{}
	bufferStderr := &bytes.Buffer{}
	logger := New("name", INFO, Output(bufferStdout), ErrorOutput(bufferStderr))
	logger.Warning(message, map[string]interface{}{})
	v := format{}

	err := json.Unmarshal(bufferStdout.Bytes(), &v)
	assert.NotNil(t, err)
	assert.Equal(t, "", bufferStdout.String())

	err = json.Unmarshal(bufferStderr.Bytes(), &v)
	assert.Nil(t, err)
	assert.Equal(t, "WARNING", v.Severity)
	assert.Equal(t, message, v.Message)
}
