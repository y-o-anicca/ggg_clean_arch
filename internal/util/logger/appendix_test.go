package logger

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendExtra(t *testing.T) {
	record := map[string]interface{}{}
	record = AppendExtra(record, map[string]interface{}{"key": "value"})
	assert.Equal(t, map[string]interface{}{"key": "value"}, record["extra"])

	record = map[string]interface{}{"extra": map[string]interface{}{"key": "value"}}
	record = AppendExtra(record, map[string]interface{}{"foo": "bar"})
	assert.Equal(t, map[string]interface{}{"key": "value", "foo": "bar"}, record["extra"])

	record = map[string]interface{}{"extra": 1}
	record = AppendExtra(record, map[string]interface{}{"key": "value"})
	assert.Equal(t, map[string]interface{}{"key": "value"}, record["extra"])
}

func TestNewRuntimeAppendix(t *testing.T) {
	appendix := NewRuntimeAppendix(DEBUG, DefaultCallerSkip)
	record := map[string]interface{}{}
	record = appendix.run(context.Background(), INFO, record)
	if _, ok := record["extra"]; ok {
		extra := record["extra"].(map[string]interface{})
		line, hasLine := extra["line"]
		assert.True(t, hasLine, "line is undefined")
		assert.IsType(t, line, 0, "line is not int")
		file, hasFile := extra["file"]
		assert.True(t, hasFile, "file is undefined")
		assert.IsType(t, file, "", "file is not string")
		path, hasPath := extra["path"]
		assert.True(t, hasPath, "path is undefined")
		assert.IsType(t, path, "", "path is not string")
	} else {
		t.Error("extra is undefined")
	}
}
