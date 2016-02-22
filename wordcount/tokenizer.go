package wordcount

import (
	"fmt"
	"gopkg.in/sensorbee/sensorbee.v0/bql/udf"
	"gopkg.in/sensorbee/sensorbee.v0/core"
	"gopkg.in/sensorbee/sensorbee.v0/data"
	"strings"
)

// Tokenizer is a UDSF that tokenizes texts in a specific field of tuples.
type Tokenizer struct {
	field string
}

// Processe implements udf.UDSF.Process. It tokenizes a field of tuples.
func (t *Tokenizer) Process(ctx *core.Context, tuple *core.Tuple, w core.Writer) error {
	var kwd []string
	if v, ok := tuple.Data[t.field]; !ok {
		return fmt.Errorf("the tuple doesn't have the required field: %v", t.field)
	} else if s, err := data.AsString(v); err != nil {
		return fmt.Errorf("'%v' field must be string: %v", t.field, err)
	} else {
		kwd = strings.Split(s, " ")
	}

	for _, k := range kwd {
		out := tuple.Copy()
		out.Data[t.field] = data.String(k)
		if err := w.Write(ctx, out); err != nil {
			return err
		}
	}
	return nil
}

// Terminate implements udf.UDSF.Terminate.
func (t *Tokenizer) Terminate(ctx *core.Context) error {
	return nil
}

// CreateTokenizer creates a new instance of Tokenizer.
func CreateTokenizer(decl udf.UDSFDeclarer, inputStream, field string) (udf.UDSF, error) {
	if err := decl.Input(inputStream, nil); err != nil {
		return nil, err
	}
	return &Tokenizer{
		field: field,
	}, nil
}
