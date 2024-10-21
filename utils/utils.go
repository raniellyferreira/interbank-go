package interutils

import (
	"strings"
	"time"

	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// StructToMap converts a struct to a map
func StructToMap(obj interface{}) map[string]string {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil
	}

	var objMap map[string]string
	err = json.Unmarshal(data, &objMap)
	if err != nil {
		return nil
	}
	return objMap
}

// JsonUnmarshal unmarshals a JSON string to an object
func JsonUnmarshal(data []byte, obj any) error {
	return json.Unmarshal(data, obj)
}

// MustParseTime parses a time string using RFC3339
func MustParseTime(timeStr string) time.Time {
	t, _ := time.Parse(time.RFC3339, timeStr)
	return t
}

// FormatTime formats a time to a string using RFC3339
func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

// GenerateUUID generates a UUID
func UUIDString() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
