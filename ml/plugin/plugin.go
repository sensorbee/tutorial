package plugin

import (
	"github.com/sensorbee/tutorial/ml"
	"gopkg.in/sensorbee/sensorbee.v0/bql/udf"
)

func init() {
	udf.MustRegisterGlobalUDF("filter_punctuation_marks",
		udf.MustConvertGeneric(ml.FilterPunctuationMarks))
	udf.MustRegisterGlobalUDF("filter_stop_words",
		udf.MustConvertGeneric(ml.FilterStopWords))
}
