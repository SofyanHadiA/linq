package services

type ICacheService interface {
	KeyNil(err error) bool
	Set(key string, value interface{}) error
	Get(key string) (string, error)
	
}
