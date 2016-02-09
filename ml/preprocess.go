package ml

// FilterPunctuationMarks removes punctuation marks such as ',' or '.' and
// replace a mark with a space. Although emoticons play an important role in
// expressing sentiment and they would be a great hint for machine learning
// algorithms, this function removes most of emoticons for simplicity.
func FilterPunctuationMarks(s string) string {
	res := []rune(s)
	for i, r := range res {
		if _, ok := punctuationMarks[r]; ok {
			res[i] = ' '
		}
	}
	return string(res)
}

// FilterStopWords filters out some stop words like "a" or "the" from the given
// array of words. This function doesn't support user defined dictionaries.
// It also filters out one-letter words and empty words (i.e. "").
// All words needs to be lower cased before applying this function.
func FilterStopWords(a []string) []string {
	res := make([]string, len(a))
	i := 0
	for _, w := range a {
		if len(w) <= 1 {
			continue
		} else if _, ok := stopWords[w]; ok {
			// concurrent read on a map is safe as long as there's no concurrent write
			continue
		}

		res[i] = w
		i++
	}
	return res[:i]
}

var (
	punctuationMarks = map[rune]struct{}{}
	stopWords        = map[string]struct{}{}
)

func init() {
	ps := []rune("!#$%^&*()=+\\|`~[{]};:'\",<.>/?")
	for _, p := range ps {
		punctuationMarks[p] = struct{}{}
	}

	// There can be more stop words but machine learning algorithms usually
	// consider words appearing too often less valuable. Thus, they don't
	// have to be removed actually. This is provided as a demonstration.
	ws := []string{"a", "the", "be", "is", "am", "are", "it", "this", "that",
		"and", "or", "not", "as", "to"}
	for _, w := range ws {
		stopWords[w] = struct{}{}
	}
}
