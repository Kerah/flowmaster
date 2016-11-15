package messages

import "github.com/Kerah/flowmaster/core"

type message struct {
	flowId core.FlowID
	body   []byte
}

func (msg *message) SetFlowId(id core.FlowID) MessageBuilder {
	msg.flowId = id
	return msg
}

func (msg *message) SetBody(body []byte) MessageBuilder {
	msg.body = body
	return msg
}

func (msg *message) Message() core.Message {
	return msg
}

func (msg *message)  Body() []byte {
	return msg.body
}

func (msg *message) Flow() core.FlowID {
	return msg.flowId
}

type MessageBuilder interface {
	SetFlowId(core.FlowID) MessageBuilder
	SetBody([]byte) MessageBuilder
	Message() core.Message

}

func New() core.Message {
	msg := &message{}
	return msg
}

func Builder() MessageBuilder {
	msg := &message{}
	return msg
}
