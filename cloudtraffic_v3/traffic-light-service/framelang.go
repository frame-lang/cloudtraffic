package trafficlight

type FrameEvent struct {
	Msg    string
	Params map[string]interface{}
	Ret    interface{}
}