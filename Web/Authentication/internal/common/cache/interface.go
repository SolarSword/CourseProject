package cache

type Cache interface {
	InitCache() error
	CloseCache()
}
