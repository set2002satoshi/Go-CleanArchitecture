package controllers

type Context interface {
	Param(key string) string
	BindJSON(obj interface{}) error
	JSON(code int, obj interface{})
}