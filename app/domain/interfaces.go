package domain

type QueryGetter interface {
	Get(key string) string
}
