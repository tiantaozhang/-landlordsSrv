package model

import (
	"time"

	"github.com/name5566/leaf/log"
	"github.com/tiantaozhang/landlordsSrv/src/server/msg"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func CreateRoom(pwd string) error {
	s := mgoc.Ref()
	defer mgoc.UnRef(s)
	err, seq := Seq("room")
	if err != nil {
		log.Error("inc room seq err:%v", err)
		return err
	}
	uid := getuid()
	_, err = s.DB(LandLord).C("room").Upsert(bson.M{"_id": seq}, bson.M{
		"_id":    seq,
		"pwd":    pwd,
		"owner":  "",
		"users":  []string{uid},
		"status": msg.Normal,
		"time":   time.Now(),
	})
	if err != nil {
		log.Error("upsert room id:%v ,err:%v", seq, err)
		return err
	}
	return nil
}

func JoinRoom(rid, pwd string) error {
	uid := getuid()
	s := mgoc.Ref()
	defer mgoc.UnRef(s)
	_, err := s.DB(LandLord).C("room").Find(bson.M{"_id": rid, "pwd": pwd}).Apply(mgo.Change{
		Update: bson.M{"addToSet": bson.M{"users": uid}},
	}, nil)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func ExitRoom(rid string) error {
	uid:=getuid()
	s:=mgoc.Ref()
	defer mgoc.UnRef(s)
	_, err := s.DB(LandLord).C("room").Find(bson.M{"_id": rid}).Apply(mgo.Change{
		Update: bson.M{"pull": bson.M{"users": uid}},
	}, nil)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
