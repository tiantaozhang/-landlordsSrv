package msg

// Ready to start game
type Ready struct {
	State int `json:"state"`
}

// Landlord for call landlord
type Landlord struct {
	State int `json:"state"`
}

// A PokeU for a user
type PokeU struct {
	State int    `json:"state"`
	Uid   string `json:"Uid"`
	Pokes []Poke `json:"pokes"`
}

// A Poke is just a p
type Poke struct {
	Color  string `json:"color"`
	Number string `json:"number"`
}

// Rest pokes
type Rest struct {
	Pokes []string `json:"pokes"`
}

// Room create
type Room struct {
	Id     string  `json:"id"`
	Pwd    string  `json:"pwd"`
	Owner  string  `json:"owner"`
	Users  [3]User `json:"users"`
	Status int     `json:"status"`
	Time   int     `json:"time,omitempty"`
}

// RoomS search
type RoomS struct {
	Id string `json:"id"`
}

// RoomP pwd
type RoomP struct {
	Id  string `json:"id"`
	Pwd string `json:"pwd"`
}

// user
type User struct {
	Uid      string `json:"uid"`
	Pwd      string `json:"pwd"`
	NickName string `json:"nickName"`
	Status   int    `json:"status"`
}
