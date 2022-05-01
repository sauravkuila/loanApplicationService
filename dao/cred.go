package dao

type UserDao interface {
	CheckAccess() error
}

func (d *DAO) CheckAccess() error {
	return nil
}
