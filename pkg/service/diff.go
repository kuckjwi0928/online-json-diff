package service

import (
	"encoding/json"
	"github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
)

type DifferService interface {
	Diff([]byte, []byte) (string, error)
}

func NewGoJsonDiffer() DifferService {
	return &goJsonDiffer{
		differ: gojsondiff.New(),
		formatConfig: formatter.AsciiFormatterConfig{
			ShowArrayIndex: true,
			Coloring:       true,
		},
	}
}

type goJsonDiffer struct {
	differ       *gojsondiff.Differ
	formatConfig formatter.AsciiFormatterConfig
}

func (g *goJsonDiffer) unmarshalObjectJson(data []byte) (map[string]interface{}, error) {
	var v map[string]interface{}
	err := json.Unmarshal(data, &v)
	return v, err
}

func (g *goJsonDiffer) unmarshalArrayJson(data []byte) ([]interface{}, error) {
	var v []interface{}
	err := json.Unmarshal(data, &v)
	return v, err
}

func (g *goJsonDiffer) Diff(left, right []byte) (string, error) {
	var d gojsondiff.Diff
	var (
		l   interface{}
		r   interface{}
		err error
	)

	// first byte is '[' means array
	if left[0] == 91 {
		l, err = g.unmarshalArrayJson(left)
		if err != nil {
			return "", err
		}
		r, err = g.unmarshalArrayJson(right)
		if err != nil {
			return "", err
		}
		d = g.differ.CompareArrays(l.([]interface{}), r.([]interface{}))
	} else {
		l, err = g.unmarshalObjectJson(left)
		if err != nil {
			return "", err
		}
		d, _ = g.differ.Compare(left, right)
	}

	if !d.Modified() {
		return "", nil
	}

	f := formatter.NewAsciiFormatter(l, g.formatConfig)
	diffString, err := f.Format(d)

	if err != nil {
		return "", err
	}

	return diffString, nil
}
