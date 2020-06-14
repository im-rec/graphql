package rest

type rest struct {}

func Initialize() *rest { return &rest{} }

type Handler interface {}