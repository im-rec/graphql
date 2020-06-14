package event

type event struct {}

func Initialize() *event { return &event{} }

type Handler interface {}