package msg

import (
	//"github.com/name5566/leaf/network"
	"github.com/name5566/leaf/network/json"
)

//var Processor network.Processor
var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Ready{})
	Processor.Register(&Landlord{})
	Processor.Register(&PokeU{})
	Processor.Register(&Poke{})
	Processor.Register(&Rest{})
	Processor.Register(&Room{})
	Processor.Register(&RoomS{})
	Processor.Register(&RoomP{})
	Processor.Register(&User{})
}
