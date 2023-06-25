package json

import (
	"bytes"
	"compress/gzip"
	"io"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var (
	// Marshal is a shortcut for json.Marshal
	Marshal = json.Marshal
	// Unmarshal is a shortcut for json.Unmarshal
	Unmarshal = json.Unmarshal
	// NewEncoder is a shortcut for json.NewEncoder
	NewEncoder = json.NewEncoder
	// NewDecoder is a shortcut for json.NewDecoder
	NewDecoder = json.NewDecoder
	Get        = json.Get
)

// MarshalIndent is a shortcut for json.MarshalIndent
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

// UnmarshalFromString is a shortcut for json.Unmarshal
func UnmarshalFromString(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}

// UnmarshalFromString is a shortcut for json.Unmarshal
func UnmarshalFromBytes(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}

// UnmarshalFromString is a shortcut for json.Unmarshal
func UnmarshalFromReader(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// MarshalToString is a shortcut for json.Marshal
func MarshalToString(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	return string(b), err
}

// MarshalToBuffer is a shortcut for json.Marshal
func MarshalToBuffer(v interface{}, buf *bytes.Buffer) error {
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	return enc.Encode(v)
}

// MarshalToWriter is a shortcut for json.Marshal
func MarshalToWriter(v interface{}, w io.Writer) error {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	return enc.Encode(v)
}

// MarshalToGzip is a shortcut for json.Marshal
func MarshalToGzip(v interface{}, w io.Writer) error {
	gz := gzip.NewWriter(w)
	defer gz.Close()
	return MarshalToWriter(v, gz)
}

// MarshalToGzip is a shortcut for json.Marshal
func MarshalToGzipBuffer(v interface{}, buf *bytes.Buffer) error {
	gz := gzip.NewWriter(buf)
	defer gz.Close()
	return MarshalToWriter(v, gz)
}

// MarshalToGzip is a shortcut for json.Marshal
func MarshalToGzipBytes(v interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := MarshalToGzipBuffer(v, buf)
	return buf.Bytes(), err
}

// MarshalToGzip is a shortcut for json.Marshal
func MarshalToGzipString(v interface{}) (string, error) {
	buf := bytes.NewBuffer(nil)
	err := MarshalToGzipBuffer(v, buf)
	return buf.String(), err
}
