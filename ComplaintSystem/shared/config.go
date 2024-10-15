package config

import (
	"time"
)

// explain why we have the "shared" folder, why we have a config here and another config in seperate projects in the lecture?
type DbConfig struct {
	MongoDuration  time.Duration
	MongoClientURI string
}

var cfgs = map[string]DbConfig{
	"prod": {
		MongoDuration: time.Second * 10,
		//MongoClientURI: "mongodb+srv://ebraremine:<19990706>@complaintsystem.wwmzy.mongodb.net/?retryWrites=true&w=majority&appName=ComplaintSystem",
		MongoClientURI: "mongodb+srv://ebraremine:19990706@complaintsystem.wwmzy.mongodb.net/",
	},
	"qa": {
		MongoDuration: time.Second * 10,
		//MongoClientURI: "mongodb+srv://ebraremine:<19990706>@complaintsystem.wwmzy.mongodb.net/?retryWrites=true&w=majority&appName=ComplaintSystem",
		MongoClientURI: "mongodb+srv://ebraremine:19990706@complaintsystem.wwmzy.mongodb.net/",
	},
	"dev": {
		MongoDuration: time.Second * 10,
		//MongoClientURI: "mongodb+srv://ebraremine:<19990706>@complaintsystem.wwmzy.mongodb.net/?retryWrites=true&w=majority&appName=ComplaintSystem",
		MongoClientURI: "mongodb+srv://ebraremine:19990706@complaintsystem.wwmzy.mongodb.net/",
	},
}

func GetDBConfig(env string) *DbConfig {
	config, isExist := cfgs[env]
	if !isExist {
		panic("config does not exist")
	}
	return &config
}
