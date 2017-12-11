package ethdb

import "time"

// TableTx orm table define
type TableTx struct {
	ID         int64     `xorm:"pk autoincr"`
	TX         string    `xorm:"notnull"`
	From       string    `xorm:"index(from_to)"`
	To         string    `xorm:"index(from_to)"`
	Asset      string    `xorm:"notnull"`
	Value      string    `xorm:"notnull"`
	Blocks     uint64    `xorm:"notnull index"`
	GasPrice   string    `xorm:"notnull"`
	Gas        string    `xorm:"notnull"`
	CreateTime time.Time `xorm:"TIMESTAMP notnull"`
}

// TableName xorm table name
func (table *TableTx) TableName() string {
	return "eth_tx"
}
