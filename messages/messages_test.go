package messages

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/Kerah/flowmaster/core"
)

func TestBuilder(t *testing.T) {
	Convey("messages building", t,  func(){
		body := "hello world!"
		serviceName := "serviceName"
		msg := Builder().SetBody([]byte(body)).SetFlowId(serviceName).Message()
		Convey("test getters", func(){
			So(msg.Flow(), ShouldEqual, serviceName)
			So(string(msg.Body()), ShouldEqual, body)
		})
		Convey("test setters", func(){
			serviceNewName := "newSrevice"
			msg.(MessageBuilder).SetFlowId(serviceNewName)
			So(msg.Flow(), ShouldEqual, serviceNewName)
			So(string(msg.Body()), ShouldEqual, body)
		})
	})

	Convey("messages marshalling", t, func(){
		body := "hello world"
		serviceName := "serviceName"
		msgTo := Builder().SetBody([]byte(body)).SetFlowId(serviceName,
		).SetFlowType(13).SetContentType(200).Message()
		data, err := msgTo.(core.Frame).Marshall()
		So(err, ShouldBeNil)
		So(data, ShouldNotBeEmpty)

		msgFrom := New()
		So(msgFrom.(core.Frame).Unmarshall(data), ShouldBeNil)
		So(msgFrom.ContentType(), ShouldEqual, msgTo.ContentType())
		So(msgFrom.FlowType(), ShouldEqual, msgTo.FlowType())
		So(msgFrom.Flow(), ShouldEqual, msgTo.Flow())
		So(string(msgFrom.Body()), ShouldEqual, string(msgTo.Body()))
	})
}
