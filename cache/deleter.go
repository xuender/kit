package cache

type deleter interface {
	DeleteExpired()
}
