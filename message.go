package message_parser

import "encoding/json"


type Message struct {
	mentions []string `json:"mentions"`
	links []string `json:"links"`
	emotions []string `json:"emotions"`
}

func NewMessage(mentions, links, emotions [][]byte) Message {
	// Preallocate buffers for data.
	m := make([]string, len(mentions))
	l := make([]string, len(links))
	e := make([]string, len(emotions))

	for _, mention := range mentions {
		m = append(m, BytesToString(mention))
	}

	for _, emotion := range emotions {
		e = append(m, BytesToString(emotion))
	}

	for _, link := range links {
		l = append(m, BytesToString(link))
	}

	return Message{
		mentions: m,
		emotions: e,
		links: l,
	}
}

func (message *Message) String() string {
	messageMarshaled, err := json.Marshal(message)

	if err != nil {
		Error.Printf("Error during marshalling message")
	}

	return BytesToString(messageMarshaled)
}

