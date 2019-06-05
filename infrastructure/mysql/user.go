package mysql

import (
	"github.com/gocraft/dbr"
)

type User struct {
	ID       int64  `db:"id"`
	Password string `db:"password"`
	Name     string `db:"name"`
	Email    string `db:"email"`
}

type UserDao struct {
	conn  *dbr.Connection
	table string
}

func NewUserDao() UserDao {
	return UserDao{
		//TODO: connection情報作成
		//conn:  ConnMaster,
		table: "user",
	}
}

func (this UserDao) FindById(userId int64) (*User, error) {
	sess := this.conn.NewSession(nil)
	var record *User
	_, err := sess.Select("*").From(this.table).Where("user_id = ?", userId).Load(&record)
	return record, err
}

