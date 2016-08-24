# message-parser
Simple utility for parsing raw messages to structured JSON

### Basic usage

Create MessageParser object that returns MessageParser interface

```golang
type MessageParser interface {
    Parse(messageRaw string) string
    ParseParallel(messageRaw string) string
}
```
Then call method Parse or ParseParallel.

```golang
messageParser := message_parser.NewMessageParser(MENTION_PATTERN, LINK_PATTERN, EMOTIONS_PATTERN)
message := messageParser.Parse("Hello  @Mike, (bicepleft) hello, https://golang.org")
```

That returns string with json object in following format.

```json
{
    "mentions":["@Mike","@John","@Peter"],
    "links":["http://google.com","https://golang.org"],
    "emotions":["(bicepleft)","(badass)"]
}
```

### Running unit tests

`go test -v`

### Running benchmarks

`go test -bench=.`
