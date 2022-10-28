package store

type Store interface {
	Repository() Repository
	Sessions() Sessions
}
