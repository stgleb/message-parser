package main

import (
	"../"
	"fmt"
)

const LINK_PATTERN = `(http|ftp|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`
const MENTION_PATTERN = "@.[a-zA-Z0-9]{3,30}"
const EMOTIONS_PATTERN = "\\(awesome\\)|\\(badass\\)|\\(bicepleft\\)|\\(allthethings\\)|\\(challengeaccepted\\)"



func main() {
	messageParser := message_parser.NewMessageParser(MENTION_PATTERN, LINK_PATTERN, EMOTIONS_PATTERN)
	message := messageParser.Parse("Hello  @Mike @John @Peter (bicepleft) (badass) hello, http://google.com https://golang.org")
	fmt.Printf("Message decoded \n %s", message)
}
