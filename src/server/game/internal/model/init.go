package model

import (
	"github.com/name5566/leaf/db/mongodb"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	LandLord = "landlord" //DB

)

var mgoc mongodb.DialContext

func init() {
	var err error
	mgoc, err = mongodb.Dial("localhost", 10)
	if err != nil {
		panic(err)
		return
	}
}

func Seq(collection string) (error, int) {
	s := mgoc.Ref()
	defer mgoc.UnRef(s)
	var res struct {
		Seq int `bson:"seq"`
	}
	_, err := s.DB(LandLord).C("seq").FindId(collection).Apply(mgo.Change{
		Update:    bson.M{"$inc": bson.M{"seq": 1}},
		ReturnNew: true,
		Upsert:    true,
	}, &res)
	return err, res.Seq
}

