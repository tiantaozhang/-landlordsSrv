package model

import (
	"github.com/name5566/leaf/db/mongodb"
)

func ()  {
	c, err := mongodb.Dial("localhost", 10)
	if err != nil {
		panic(err)
		return
	}
	defer c.Close()
}

