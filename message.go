package message_parser

import "encoding/json"

type Message struct {
	Mentions []string `json:"mentions"`
	Links    []string `json:"links"`
	Emotions []string `json:"emotions"`
}

func NewMessage(mentions, links, emotions [][]byte) Message {
	// Preallocate buffers for data.
	m := make([]string, 0, len(mentions))
	l := make([]string, 0, len(links))
	e := make([]string, 0, len(emotions))

	for index, _ := range mentions {
		m = append(m, BytesToString(mentions[index]))
	}

	for index, _ := range emotions {
		e = append(e, BytesToString(emotions[index]))
	}

	for index, _ := range links {
		l = append(l, BytesToString(links[index]))
	}

	return Message{
		Mentions: m,
		Emotions: e,
		Links:    l,
	}
}

func (message *Message) String() string {
	messageMarshaled, err := json.Marshal(*message)
	Info.Printf("%s", message.Links)

	if err != nil {
		Error.Printf("Error during marshalling message")
	}

	return BytesToString(messageMarshaled)
}
