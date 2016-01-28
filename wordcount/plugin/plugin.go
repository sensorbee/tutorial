package plugin

import (
	"github.com/sensorbee/tutorial/wordcount"
	"pfi/sensorbee/sensorbee/bql"
	"pfi/sensorbee/sensorbee/bql/udf"
)

func init() {
	bql.MustRegisterGlobalSourceCreator("wc_sentences",
		bql.SourceCreatorFunc(wordcount.CreateSentences))
	udf.MustRegisterGlobalUDSFCreator("wc_tokenizer",
		udf.MustConvertToUDSFCreator(wordcount.CreateTokenizer))
}
