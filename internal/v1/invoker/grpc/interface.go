package grpc

type grpc struct {}

func Initialize() *grpc { return &grpc{} }

type Handler interface {}
