package db

import (
	"errors"
	"log"

	pb "golang/example5/proto"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"context"
)

type Db struct {
	engine *xorm.Engine
}

var (
	tables []interface{}
)

func (db *Db) ConnectDb() error {
	var err error
	db.engine, err = xorm.NewEngine("mysql", "truong:root@/test2?charset=utf8")
	if err != nil {
		return errors.New("Connect database faild")
	}
	log.Println("Connect database success")
	db.engine.ShowSQL(true)
	return nil
}

func (db *Db) InitDatabase() error {
	initTables()
	err := db.engine.CreateTables(tables...)
	if err != nil {
		return err
	}
	return nil
}

func initTables() {
	tables = append(tables, new(pb.UserPartner))
}


// ----------------- User Partner ------------------

func (db *Db) GetUserPartner(ctx context.Context, rq *pb.UserPartnerRequest) ([]*pb.UserPartner, error) {
	ss := db.engine.Table(pb.UserPartner{})
	if rq.GetUserId() != "" {
		ss.And("user_id = ?", rq.GetUserId())
	}
	if rq.GetPhone() != "" {
		ss.And("phone = ?" , rq.GetPhone())
	}

	if rq.GetLimit() != 0 {
		ss.Limit(int(rq.GetLimit()))
	}
	userPartner := make([]*pb.UserPartner, 0)
	err := ss.Find(&userPartner)
	if err != nil {
		return nil , err
	}
	return userPartner , nil
}


