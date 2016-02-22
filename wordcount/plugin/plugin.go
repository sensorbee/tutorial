package plugin

import (
	"github.com/sensorbee/tutorial/wordcount"
	"gopkg.in/sensorbee/sensorbee.v0/bql"
	"gopkg.in/sensorbee/sensorbee.v0/bql/udf"
)

func init() {
	bql.MustRegisterGlobalSourceCreator("wc_sentences",
		bql.SourceCreatorFunc(wordcount.CreateSentences))
	udf.MustRegisterGlobalUDSFCreator("wc_tokenizer",
		udf.MustConvertToUDSFCreator(wordcount.CreateTokenizer))
}
