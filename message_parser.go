package message_parser

import (
	"regexp"
	"log"
	"io"
	"os"
)


type MessageParser interface {
	parse(string) string
}

type MessageParserImpl struct {
	mentionRegexp regexp.Regexp
	linkRegexp    regexp.Regexp
	emotionsRegexp    regexp.Regexp
}

var (
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func InitLoggers(debugHandle io.Writer,
		infoHandle io.Writer,
		warningHandle io.Writer,
		errorHandle io.Writer) {

	Debug = log.New(debugHandle,
		"DEBUG: ",
		log.Ldate | log.Ltime | log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate | log.Ltime | log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate | log.Ltime | log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate | log.Ltime | log.Lshortfile)
}

func init() {
	InitLoggers(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
}

func NewMessageParser(mentionPattern, linkPattern, emotions string) MessageParserImpl {
	Info.Printf("Create new MessageParser")
	mentionRe, err := regexp.Compile(mentionPattern)

	if err != nil {
		Error.Printf("Error %s during compiling metions regexp %s",
			err.Error(),
			mentionPattern)
	}

	linkRe, err2 := regexp.Compile(linkPattern)

	if err2 != nil {
		Error.Printf("Error %s during compiling link regexp %s",
			err2.Error(),
			linkPattern)
	}

	emotionsRe, err3 := regexp.Compile(emotions)

	if err3 != nil {
		Error.Printf("Error %s during compiling emotions regexp %s",
			err2.Error(),
			linkPattern)
	}

	return MessageParserImpl{
		mentionRegexp: *mentionRe,
		linkRegexp: *linkRe,
		emotionsRegexp: *emotionsRe,
	}
}

// TODO: consider possibility to change Find by FindAll
func findAll(re regexp.Regexp, message []byte) [][]byte {
	// Use FindAll from regexp package, argument n int
	// might be upper bound of entries to return.
	result := re.FindAll(message, 1 << 32)

	return result
}

func (messageParser *MessageParserImpl) Parse(messageRaw string) string {
	Info.Printf("Start parsing message %s", messageRaw)
	messageSlice := StringToByteSlice(messageRaw)
	Info.Printf("Convert string to byte slice")

	mentions := findAll(messageParser.mentionRegexp, messageSlice)
	Info.Printf("Found %d mentions", len(mentions))
	links := findAll(messageParser.linkRegexp, messageSlice)
	Info.Printf("Found %d links", len(links))
	emotions := findAll(messageParser.emotionsRegexp, messageSlice)
	Info.Printf("Found %d emotions", len(emotions))

	// Create new Message object
	message := NewMessage(mentions, links, emotions)

	return message.String()
}

