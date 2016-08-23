package main

import (
	"../"
	"fmt"
)

const LINK_PATTERN = "_^(?:(?:https?|ftp)://)(?:\\S+(?::\\S*)?@)?(?:(?!10(?:\\.\\d{1,3}){3})(?!127(?:\\.\\d{1,3}){3})(?!169\\.254(?:\\.\\d{1,3}){2})(?!192\\.168(?:\\.\\d{1,3}){2})(?!172\\.(?:1[6-9]|2\\d|3[0-1])(?:\\.\\d{1,3}){2})(?:[1-9]\\d?|1\\d\\d|2[01]\\d|22[0-3])(?:\\.(?:1?\\d{1,2}|2[0-4]\\d|25[0-5])){2}(?:\\.(?:[1-9]\\d?|1\\d\\d|2[0-4]\\d|25[0-4]))|(?:(?:[a-z\x{00a1}-\x{ffff}0-9]+-?)*[a-z\x{00a1}-\x{ffff}0-9]+)(?:\\.(?:[a-z\x{00a1}-\x{ffff}0-9]+-?)*[a-z\x{00a1}-\x{ffff}0-9]+)*(?:\\.(?:[a-z\x{00a1}-\x{ffff}]{2,})))(?::\\d{2,5})?(?:/[^\\s]*)?$_iuS"
const MENTION_PATTERN = "@*"
const EMOTIONS_PATTERN = "\\(awesome\\)|\\(badass\\)|\\(bicepleft\\)|\\(allthethings\\)|\\(challengeaccepted\\)"



func main() {
	messageParser := message_parser.NewMessageParser(MENTION_PATTERN, LINK_PATTERN, EMOTIONS_PATTERN)
	message := messageParser.Parse("Hello @Mike (bicepleft) hello, http://google.com")
	fmt.Printf("Message decoded \n %s", message)
}
