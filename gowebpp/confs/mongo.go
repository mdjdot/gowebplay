package confs

import "gopkg.in/mgo.v2"

// MongoConn mongo 连接
var MongoConn *mgo.Session

// InitMongo 初始化 mongodb
func InitMongo() {
	conn, err := mgo.Dial("mongodb://dm:dmtest@127.0.0.1:27017/appdb")
	if err != nil {
		panic(err)
	}
	MongoConn = conn
}
