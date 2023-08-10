package service

import "github.com/deta/deta-go/service/base"

var SokujiDb *base.Base
var SokujiChannelSokujiIdDb *base.Base
var SokujiUserChannelIdDb *base.Base

func init() {
	var err error
	SokujiDb, err = base.New(Deta, "Sokuji")
	if err != nil {
		panic(err)
	}
	SokujiChannelSokujiIdDb, err = base.New(Deta, "SokujiChannelSokujiId")
	if err != nil {
		panic(err)
	}
	SokujiUserChannelIdDb, err = base.New(Deta, "SokujiUserChannelId")
	if err != nil {
		panic(err)
	}
}

func GetSokuji(id string) (map[string]interface{}, error) {
	var i map[string]interface{}
	err := SokujiDb.Get(id, &i)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func GetSokujiByChannelId(channelId string) (map[string]interface{}, error) {
	var i map[string]interface{}
	err := SokujiChannelSokujiIdDb.Get(channelId, &i)
	if err != nil {
		return nil, err
	}
	return GetSokuji(i["sokujiId"].(string))
}

func GetSokujiByUserId(userId string) (map[string]interface{}, error) {
	var i map[string]interface{}
	err := SokujiUserChannelIdDb.Get(userId, &i)
	if err != nil {
		return nil, err
	}
	return GetSokujiByChannelId(i["channelId"].(string))
}
