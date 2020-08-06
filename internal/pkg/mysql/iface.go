package mysql

type Iface interface {
	Insert()
	Update()
	Delete()
	Find()
	FindOne()
}
