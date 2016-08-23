package main

import (
	"regexp"
	"log"
	"io"
	"unsafe"
	"reflect"
)

const LINK_PATTERN = "_^(?:(?:https?|ftp)://)(?:\\S+(?::\\S*)?@)?(?:(?!10(?:\\.\\d{1,3}){3})(?!127(?:\\.\\d{1,3}){3})(?!169\\.254(?:\\.\\d{1,3}){2})(?!192\\.168(?:\\.\\d{1,3}){2})(?!172\\.(?:1[6-9]|2\\d|3[0-1])(?:\\.\\d{1,3}){2})(?:[1-9]\\d?|1\\d\\d|2[01]\\d|22[0-3])(?:\\.(?:1?\\d{1,2}|2[0-4]\\d|25[0-5])){2}(?:\\.(?:[1-9]\\d?|1\\d\\d|2[0-4]\\d|25[0-4]))|(?:(?:[a-z\x{00a1}-\x{ffff}0-9]+-?)*[a-z\x{00a1}-\x{ffff}0-9]+)(?:\\.(?:[a-z\x{00a1}-\x{ffff}0-9]+-?)*[a-z\x{00a1}-\x{ffff}0-9]+)*(?:\\.(?:[a-z\x{00a1}-\x{ffff}]{2,})))(?::\\d{2,5})?(?:/[^\\s]*)?$_iuS"
const MENTION_PATTERN = "@*"

type MessageParser interface {
	parse(string) string
}

type MessageParserImpl struct {
	mentionRegexp regexp.Regexp
	linkRegexp    regexp.Regexp
	mentionsRegexp    regexp.Regexp
}

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Init(traceHandle io.Writer,
		infoHandle io.Writer,
		warningHandle io.Writer,
		errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
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

func NewMessageParser(mentionPattern, linkPattern string) MessageParserImpl {
	mentionRe, err := regexp.Compile(mentionPattern)

	if err != nil {
		Error.Printf("Error %s during compiling regexp %s",
			err.Error(),
			mentionPattern)
	}

	linkRe, err2 := regexp.Compile(linkPattern)

	if err2 != nil {
		Error.Printf("Error %s during compiling regexp %s",
			err2.Error(),
			linkPattern)
	}

	linkRe, err3 := regexp.Compile(mentions)

	if err2 != nil {
		Error.Printf("Error %s during compiling regexp %s",
			err2.Error(),
			linkPattern)
	}

	return MessageParserImpl{
		mentionRegexp: mentionRe,
		linkRegexp: linkRe,
	}
}

func StringToByteSlice(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func findAll(re regexp.Regexp, message []byte) [][]byte {
	var result [][]byte
	result = make([]byte, 1)

	for {
		found := re.Find(message)

		if found == nil {
			break
		} else {
			result = append(result, found)
		}
	}

	return result
}

func (messageParser *MessageParserImpl) parse(message string) {
	messageSlice := StringToByteSlice(message)

	mentions := findAll(messageParser.mentionRegexp, messageSlice)
	links := findAll(messageParser.linkRegexp, messageSlice)
}

func main() {
	messageParser := NewMessageParser(MENTION_PATTERN, LINK_PATTERN)
	messageParser.parse("Hello @Mike")
}
