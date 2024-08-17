package support

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
)

type Util interface {
	LogJson()
}

type Observer struct {
	LogLevel int8 // replace with Enum
	Logger   slog.Logger
}

func (o *Observer) LogJson(v any) {
	if o.LogLevel < 3 {
		return
	}
	_json, err := json.Marshal(v)
	if err != nil {
		o.Logger.Warn("Failed to marshal templates to JSON", "error", err)
	}
	logJson := bytes.Buffer{}
	_ = json.Indent(&logJson, _json, "", "  ")
	fmt.Printf("%s", logJson)
}
