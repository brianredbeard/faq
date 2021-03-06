package formats

import (
	"bytes"
	"encoding/json"

	"github.com/alecthomas/chroma/quick"
)

type jsonEncoding struct{}

func (jsonEncoding) MarshalJSONBytes(jsonBytes []byte) ([]byte, error) {
	// It's already JSON, silly!
	return jsonBytes, nil
}

func (jsonEncoding) UnmarshalJSONBytes(jsonBytes []byte) ([]byte, error) {
	// It's already JSON, silly!
	return jsonBytes, nil
}

func (jsonEncoding) Raw(jsonBytes []byte) ([]byte, error) {
	// This is a super naive attempt to just strip off quotes.
	quoteByte := byte(0x22)
	if jsonBytes[0] == quoteByte && jsonBytes[len(jsonBytes)-1] == quoteByte {
		return jsonBytes[1 : len(jsonBytes)-1], nil
	}
	return jsonBytes, nil
}

func (jsonEncoding) PrettyPrint(jsonBytes []byte) ([]byte, error) {
	var i interface{}
	err := json.Unmarshal(jsonBytes, &i)
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(i, "", "  ")
}

func (jsonEncoding) Color(jsonBytes []byte) ([]byte, error) {
	var b bytes.Buffer
	if err := quick.Highlight(&b, string(jsonBytes), "json", ChromaFormatter(), ChromaStyle()); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func init() {
	ByName["json"] = jsonEncoding{}
	ByName["js"] = jsonEncoding{}
	ByName["javascript"] = jsonEncoding{}
}
