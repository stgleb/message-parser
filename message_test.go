package message_parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessageToString(t *testing.T) {
	expected := "{\"mentions\":[\"mention\"],\"links\":[\"link\"],\"emotions\":[\"emotion\"]}"
	var mentions [][]byte
	var emotions [][]byte
	var links [][]byte

	mentions = append(mentions, []byte("mention"))
	emotions = append(emotions, []byte("emotion"))
	links = append(links, []byte("link"))

	message := NewMessage(mentions, links, emotions)

	messageString := message.String()
	assert.Equal(t, expected, messageString)
}
