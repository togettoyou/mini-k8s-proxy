package main

type ResourceEventHandler struct {
	Ev chan<- interface{}
}

func (reh *ResourceEventHandler) OnAdd(obj interface{}) {
	eventHandlerFunc(reh.Ev, obj)
}

func (reh *ResourceEventHandler) OnUpdate(oldObj, newObj interface{}) {
	eventHandlerFunc(reh.Ev, newObj)
}

func (reh *ResourceEventHandler) OnDelete(obj interface{}) {
	eventHandlerFunc(reh.Ev, obj)
}

func eventHandlerFunc(events chan<- interface{}, obj interface{}) {
	select {
	case events <- obj:
	default:
	}
}
