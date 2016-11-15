package messages

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/Kerah/flowmaster/core"
)

func TestBuilder(t *testing.T) {
	Convey("messages", t,  func(){
		body := "hello world!"
		serviceName := core.FlowID("serviceName")
		msg := Builder().SetBody([]byte(body)).SetFlowId(serviceName).Message()
		Convey("test getters", func(){
			So(msg.Flow(), ShouldEqual, serviceName)
			So(string(msg.Body()), ShouldEqual, body)
		})
		Convey("test setters", func(){
			serviceNewName := core.FlowID("newSrevice")
			msg.(MessageBuilder).SetFlowId(serviceNewName)
			So(msg.Flow(), ShouldEqual, serviceNewName)
			So(string(msg.Body()), ShouldEqual, body)
		})
	})
}
