package messages

import (
	"github.com/Kerah/flowmaster/core"
	"github.com/Kerah/flowmaster/messages/internal/flat"
	flatbuffers "github.com/google/flatbuffers/go"
)

func New() core.Message {
	m := &msg{}
	return m
}

type msg struct {
	raw         *flat.Message
	flowId      string
	contentType uint32
	flowType    int32
	body        []byte
}

func (msg *msg) SetFlowId(id string) MessageBuilder {
	msg.flowId = id
	return msg
}

func (msg *msg) SetBody(body []byte) MessageBuilder {
	msg.body = body
	return msg
}

func (msg *msg) Message() core.Message {
	return msg
}

func (msg *msg) Unmarshall(data []byte) error {
	msg.raw = flat.GetRootAsMessage(data, 0)
	return nil
}

func (msg *msg) Marshall() (data []byte, err error) {
	builder := flatbuffers.NewBuilder(1024)
	flowId := builder.CreateString(msg.flowId)
	body := builder.CreateByteVector(msg.body)

	flat.MessageStart(builder)
	flat.MessageAddContentType(builder, msg.contentType)
	flat.MessageAddFlowType(builder, msg.flowType)
	flat.MessageAddFlowId(builder, flowId)
	flat.MessageAddBody(builder, body)
	builder.Finish(flat.MessageEnd(builder))
	data = builder.FinishedBytes()
	return
}
func (msg *msg) Body() []byte {
	if msg.raw != nil {
		return msg.raw.Body()
	}
	return msg.body
}

func (msg *msg) Flow() string {
	if msg.raw != nil {
		return string(msg.raw.FlowId())
	}
	return msg.flowId
}

func (msg *msg) FlowType() int32 {
	if msg.raw != nil {
		return msg.raw.FlowType()
	}
	return msg.flowType
}

func (msg *msg) ContentType() uint32 {
	if msg.raw != nil {
		return msg.raw.ContentType()
	}
	return msg.contentType
}

func (msg *msg) SetFlowType(flowType int32) MessageBuilder {
	msg.flowType = flowType
	return msg
}
func (msg *msg) SetContentType(contentType uint32) MessageBuilder {
	msg.contentType = contentType
	return msg
}

type MessageBuilder interface {
	SetFlowId(string) MessageBuilder
	SetBody([]byte) MessageBuilder
	SetFlowType(int32) MessageBuilder
	SetContentType(uint32) MessageBuilder
	Message() core.Message
}

func Builder() MessageBuilder {
	msg := &msg{}
	return msg
}
