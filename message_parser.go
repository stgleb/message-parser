package message_parser

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

type MessageParser interface {
	Parse(messageRaw string) string
	ParseParallel(messageRaw string) string
}

type MessageParserImpl struct {
	mentionRegexp  regexp.Regexp
	linkRegexp     regexp.Regexp
	emotionsRegexp regexp.Regexp
}

var (
	Debug *log.Logger
	Info  *log.Logger
	Error *log.Logger
)

func InitLoggers(debugHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Debug = log.New(debugHandle,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func init() {
	InitLoggers(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
}

func NewMessageParser(mentionPattern, linkPattern, emotions string) MessageParserImpl {
	Debug.Printf("Create new MessageParser")
	mentionRe, err := regexp.CompilePOSIX(mentionPattern)

	if err != nil {
		Error.Fatalf("Error %s during compiling metions regexp %s",
			err.Error(),
			mentionPattern)
	}

	linkRe, err2 := regexp.Compile(linkPattern)

	if err2 != nil {
		Error.Fatalf("Error %s during compiling link regexp %s",
			err2.Error(),
			linkPattern)
	}

	emotionsRe, err3 := regexp.Compile(emotions)

	if err3 != nil {
		Error.Fatalf("Error %s during compiling emotions regexp %s",
			err2.Error(),
			linkPattern)
	}

	return MessageParserImpl{
		mentionRegexp:  *mentionRe,
		linkRegexp:     *linkRe,
		emotionsRegexp: *emotionsRe,
	}
}

func findAll(re regexp.Regexp, message []byte) [][]byte {
	// Use FindAll from regexp package, argument n int
	// might be upper bound of entries to return.
	result := re.FindAll(message, 1<<32)

	return result
}

func findAllParallel(re regexp.Regexp, message []byte, resultChan chan<- [][]byte) {
	// Use FindAll from regexp package, argument n int
	// might be upper bound of entries to return.
	result := re.FindAll(message, 1<<32)
	resultChan <- result
}

func (messageParser *MessageParserImpl) Parse(messageRaw string) string {
	Debug.Printf("Start parsing message %s", messageRaw)
	messageSlice := StringToByteSlice(messageRaw)
	Debug.Printf("Convert string to byte slice")

	mentions := findAll(messageParser.mentionRegexp, messageSlice)
	Debug.Printf("Found %d mentions", len(mentions))
	links := findAll(messageParser.linkRegexp, messageSlice)
	Debug.Printf("Found %d links", len(links))
	emotions := findAll(messageParser.emotionsRegexp, messageSlice)
	Debug.Printf("Found %d emotions", len(emotions))

	// Create new Message object
	message := NewMessage(mentions, links, emotions)

	return message.String()
}

func (messageParser *MessageParserImpl) ParseParallel(messageRaw string) string {
	var mentions [][]byte
	var links [][]byte
	var emotions [][]byte

	Debug.Printf("Start parsing message %s", messageRaw)
	messageSlice := StringToByteSlice(messageRaw)
	Debug.Printf("Convert string to byte slice")

	mentionsChan := make(chan [][]byte)
	emotionsChan := make(chan [][]byte)
	linksChan := make(chan [][]byte)

	// Pass channels for results to parsing goroutinges
	go findAllParallel(messageParser.mentionRegexp,
		messageSlice,
		mentionsChan)
	go findAllParallel(messageParser.linkRegexp,
		messageSlice,
		linksChan)
	go findAllParallel(messageParser.emotionsRegexp,
		messageSlice,
		emotionsChan)

	// Collect results from all parsing goroutines
	for i := 0; i < 3; i += 1 {
		select {
		case mentions = <-mentionsChan:
			Debug.Printf("Found %d mentions", len(mentions))
		case links = <-linksChan:
			Debug.Printf("Found %d links", len(links))
		case emotions = <-emotionsChan:
			Debug.Printf("Found %d emotions", len(emotions))
		}
	}

	// Create new Message object
	message := NewMessage(mentions, links, emotions)

	return message.String()
}
