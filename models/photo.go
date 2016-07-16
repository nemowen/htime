package models

import (
	"time"
)

//相册表
type Photo struct {
	Id       int64     `xorm:int`
	Albumid  int64     `xorm:int`
	Des      string    `xorm:varchar(100)`
	Posttime time.Time `xorm:"DateTime created"`
	Url      string    `xorm:"varchar(70)"`
	Small    string    `xorm:"-"`
}

func (m *Photo) Insert() error {
	if _, err := orm.InsertOne(m); err != nil {
		return err
	}
	return nil
}

func (m *Photo) Read(fields ...string) error {
	if b, err := orm.Cols(fields...).Get(m); !b {
		return err
	}
	return nil
}

func (m *Photo) Update(fields ...string) error {
	if _, err := orm.Cols(fields...).Update(m); err != nil {
		return err
	}
	return nil
}

func (m *Photo) Delete() error {
	if _, err := orm.Delete(m); err != nil {
		return err
	}
	return nil
}
