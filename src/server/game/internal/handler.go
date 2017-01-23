package internal

import (
	"reflect"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"server/msg"
)

func init() {
	handler(&msg.Ready{}, handleReady)
	handler(&msg.Room{}, handleRoom)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleReady(args []interface{}) {
	// 收到的 Ready 消息
	m := args[0].(*msg.Ready)
	// Todo 消息的发送者,uid
	user := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("ready %v", m.State)

	//房间对应的user -> 1


	// 给发送者回应一个 Hello 消息
	user.WriteMsg(&msg.Ready{
		State: "1",
	})
}
