package ml

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestFilgerPunctuationMarks(t *testing.T) {
	Convey("Given a sentence", t, func() {
		s := "a, b (c); d: e."

		Convey("when applying FilterPunctuationMarks to it", func() {
			res := FilterPunctuationMarks(s)

			Convey("it should remove punctuation marks", func() {
				So(res, ShouldEqual, "a  b  c   d  e ")
			})
		})
	})
}

func TestFilterStopWords(t *testing.T) {
	Convey("Given a sequence of words", t, func() {
		s := strings.Split("this is a very good pen v ", " ")

		Convey("when applying FilterStopWords", func() {
			res := FilterStopWords(s)

			Convey("it should remove predefined stop words", func() {
				So(res, ShouldResemble, []string{"very", "good", "pen"})
			})
		})
	})
}
