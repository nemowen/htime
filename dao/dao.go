package dao

type IDaoManager interface {
	Add(o *interface{}) (state bool, err error)
	Del(id int64) (state bool, err error)
	Edit(o *interface{}) (state bool, err error)
	Get(id int64) (o *interface{}, err error)
}
