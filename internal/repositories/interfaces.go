package repositories

type Deletable interface {
	Delete(id int) error
}
