package message_parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const LINK_PATTERN = `(http|ftp|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`
const MENTION_PATTERN = "@.[a-zA-Z0-9]{3,30}"
const EMOTIONS_PATTERN = "\\(awesome\\)|\\(badass\\)|\\(bicepleft\\)|\\(allthethings\\)|\\(challengeaccepted\\)"

var tmp string

func TestMessageParse(t *testing.T) {
	expected := "{\"mentions\":[\"@Mike\",\"@John\",\"@Peter\"],\"links\":[\"http://google.com\",\"https://golang.org\"],\"emotions\":[\"(bicepleft)\",\"(badass)\"]}"
	messageParser := NewMessageParser(MENTION_PATTERN, LINK_PATTERN, EMOTIONS_PATTERN)
	actual := messageParser.Parse("Hello  @Mike @John @Peter (bicepleft) (badass) hello, http://google.com https://golang.org")
	assert.Equal(t, expected, actual)
}

func TestMessageParseParallel(t *testing.T) {
	expected := "{\"mentions\":[\"@Mike\",\"@John\",\"@Peter\"],\"links\":[\"http://google.com\",\"https://golang.org\"],\"emotions\":[\"(bicepleft)\",\"(badass)\"]}"
	messageParser := NewMessageParser(MENTION_PATTERN, LINK_PATTERN, EMOTIONS_PATTERN)
	actual := messageParser.ParseParallel("Hello  @Mike @John @Peter (bicepleft) (badass) hello, http://google.com https://golang.org")
	assert.Equal(t, expected, actual)
}

func BenchmarkMessageParse(b *testing.B) {
	messageParser := NewMessageParser(MENTION_PATTERN, LINK_PATTERN, EMOTIONS_PATTERN)
	b.ResetTimer()
	var result string

	// Store results of computing in package level variables
	// to avoid compiler optimizations.
	for i := 0; i < b.N; i++ {
		result = messageParser.Parse("Hello  @Mike @John @Peter (bicepleft) (badass) hello, http://google.com https://golang.org")
	}

	tmp = result
}

func BenchmarkMessageParseParallel(b *testing.B) {
	messageParser := NewMessageParser(MENTION_PATTERN, LINK_PATTERN, EMOTIONS_PATTERN)
	b.ResetTimer()
	var result string

	// Store results of computing in package level variables
	// to avoid compiler optimizations.
	for i := 0; i < b.N; i++ {
		result = messageParser.ParseParallel("Hello  @Mike @John @Peter (bicepleft) (badass) hello, http://google.com https://golang.org")
	}

	tmp = result
}
