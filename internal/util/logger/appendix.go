package logger

import (
	"context"
	"path"
	"runtime"
)

type (
	Appendix interface {
		run(ctx context.Context, level Level, record map[string]interface{}) map[string]interface{}
	}
	runtimeAppendix struct {
		level Level
		skip  int
	}
	contextAppendix struct {
		contextKey interface{}
		extraKey   string
	}
)

const DefaultCallerSkip = 3

// AppendExtra extraに値を追加する
func AppendExtra(record, data map[string]interface{}) map[string]interface{} {
	extra := map[string]interface{}{}
	if _, ok := record["extra"]; ok {
		extra, ok = record["extra"].(map[string]interface{})
		if !ok {
			extra = map[string]interface{}{}
		}
	}
	for key, value := range data {
		extra[key] = value
	}
	record["extra"] = extra
	return record
}

// NewRuntimeAppendix extraに実行ファイル情報を追加するAppendix
func NewRuntimeAppendix(level Level, skip int) Appendix {
	return runtimeAppendix{level: level, skip: skip}
}

func (r runtimeAppendix) run(_ context.Context, level Level, record map[string]interface{}) map[string]interface{} {
	if level < r.level {
		return record
	}
	if _, file, line, ok := runtime.Caller(r.skip); ok {
		record = AppendExtra(record, map[string]interface{}{
			"file": path.Base(file),
			"path": file,
			"line": line,
		})
	}
	return record
}

// NewContextAppendix 指定されたkeyを使ってctxからextraに値を入れるAppendix
func NewContextAppendix(contextKey interface{}, extraKey string) Appendix {
	return contextAppendix{contextKey: contextKey, extraKey: extraKey}
}

func (e contextAppendix) run(ctx context.Context, _ Level, record map[string]interface{}) map[string]interface{} {
	extra := map[string]interface{}{}
	if _, ok := record["extra"]; ok {
		extra, ok = record["extra"].(map[string]interface{})
		if !ok {
			extra = map[string]interface{}{}
		}
	}
	if value, ok := ctx.Value(e.contextKey).(string); ok {
		extra[e.extraKey] = value
		record["extra"] = extra
	}
	return record
}
