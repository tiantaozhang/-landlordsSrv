package msg

import (
	"encoding/binary"
	"net"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}
	//Hello
	//
	data := []byte(`{
		"Hello": {
			"Name": "leaf"
		}
	}`)

	m := make([]byte, 2+len(data))
	binary.BigEndian.PutUint16(m, uint16((len(data))))
	copy(m[2:], data)
	conn.Write(m)
	time.Sleep(1000)
}
