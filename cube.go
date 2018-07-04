package cube

import (
	"time"
	"encoding/json"
	"errors"
)

var (
	ErrorTimeout = errors.New("cube: request timeout")
)

type Message struct {
	Version string          `json:"version"`
	Id      string          `json:"id"`
	From    string          `json:"from"`
	To      string          `json:"to"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
}

type LogMessageParams struct {
	Time       int64  `json:"time"`
	Id         string `json:"id"`
	Class      string `json:"class"`
	InstanceId string `json:"instanceId"`
	Level      string `json:"level"`
	Text       string `json:"text"`
}

type Cube interface {
	GetParam(param string) string
	GetClass() string
	GetInstanceId() string

	PublishMessage(channel string, message Message) error
	MakeRequest(channel string, message Message, timeout time.Duration) (Message, error)

	LogDebug(text string) error
	LogError(text string) error
	LogFatal(text string) error
	LogInfo(text string) error
	LogWarning(text string) error
	LogTrace(text string) error
}

type HandlerInterface interface {
	OnStart(instance Cube)
	OnStop(instance Cube)
	OnReceiveMessage(instance Cube, channel string, message Message)
	OnReceiveRequest(instance Cube, channel string, message Message, replyToRequest func(Message) error)
}
