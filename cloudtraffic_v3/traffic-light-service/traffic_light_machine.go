// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files
package trafficlight

import (
	"encoding/json"
)

func NewTrafficLight(manager TrafficLightManager) TrafficLight {
    m := &trafficLightStruct{}
    m._manager_ = manager
    
    // Validate interfaces
    var _ TrafficLight = m
    var _ TrafficLight_actions = m
    
    // Create and intialize start state compartment.
    m._compartment_ = NewTrafficLightCompartment(TrafficLightState_Begin)
    
    // Override domain variables.
    m.flashColor = ""
    
    // Send system start event
    e := FrameEvent{Msg:">"}
    m._mux_(&e)
    return m
}


type TrafficLightState uint

const (
    TrafficLightState_Begin TrafficLightState = iota
    TrafficLightState_Red
    TrafficLightState_Green
    TrafficLightState_Yellow
    TrafficLightState_FlashingRed
    TrafficLightState_End
    TrafficLightState_Working
)

type Marshal interface {
    Marshal() []byte
}

type TrafficLight interface {
    Marshal
    Start() 
    Stop() 
    Tick() 
    SystemError() 
    SystemRestart() 
    ChangeColor(color string) 
    GetColor() string
}

type TrafficLight_actions interface {
    initTrafficLight() 
    enterRed() 
    enterGreen() 
    enterYellow() 
    enterFlashingRed() 
    startWorkingTimer() 
    stopWorkingTimer() 
    startFlashingTimer() 
    stopFlashingTimer() 
    changeFlashingAnimation() 
    destroyTrafficLight() 
}


type trafficLightStruct struct {
    _manager_ TrafficLightManager
    _compartment_ *TrafficLightCompartment
    _nextCompartment_ *TrafficLightCompartment
    flashColor string
}

type marshalStruct struct {
    TrafficLightCompartment
    FlashColor string
}

func LoadTrafficLight(manager TrafficLightManager, data []byte) TrafficLight {
    m := &trafficLightStruct{}
    m._manager_ = manager
    
    // Validate interfaces
    var _ TrafficLight = m
    var _ TrafficLight_actions = m
    
    // Unmarshal
    var marshal marshalStruct
    err := json.Unmarshal(data, &marshal)
    if err != nil {
        return nil
    }
    
    // Initialize machine
    m._compartment_ = &marshal.TrafficLightCompartment
    
    m.flashColor = marshal.FlashColor
    
    return m
    
}

func (m *trafficLightStruct) MarshalJSON() ([]byte, error) {
    data := marshalStruct{
        TrafficLightCompartment: *m._compartment_,
        FlashColor: m.flashColor,
    }
    return json.Marshal(data)
}

func (m *trafficLightStruct) Marshal() []byte {
    data, err := json.Marshal(m)
    if err != nil {
        return nil
    }
    return data
    
}
//===================== Interface Block ===================//

func (m *trafficLightStruct) Start()  {
    e := FrameEvent{Msg:"start"}
    m._mux_(&e)
}

func (m *trafficLightStruct) Stop()  {
    e := FrameEvent{Msg:"stop"}
    m._mux_(&e)
}

func (m *trafficLightStruct) Tick()  {
    e := FrameEvent{Msg:"tick"}
    m._mux_(&e)
}

func (m *trafficLightStruct) SystemError()  {
    e := FrameEvent{Msg:"systemError"}
    m._mux_(&e)
}

func (m *trafficLightStruct) SystemRestart()  {
    e := FrameEvent{Msg:"systemRestart"}
    m._mux_(&e)
}

func (m *trafficLightStruct) ChangeColor(color string)  {
    params := make(map[string]interface{})
    params["color"] = color
    e := FrameEvent{Msg:"changeColor", Params:params}
    m._mux_(&e)
}

func (m *trafficLightStruct) GetColor() string {
    e := FrameEvent{Msg:"getColor"}
    m._mux_(&e)
    return  e.Ret.(string)
}

//====================== Multiplexer ====================//

func (m *trafficLightStruct) _mux_(e *FrameEvent) {
    switch m._compartment_.State {
    case TrafficLightState_Begin:
        m._TrafficLightState_Begin_(e)
    case TrafficLightState_Red:
        m._TrafficLightState_Red_(e)
    case TrafficLightState_Green:
        m._TrafficLightState_Green_(e)
    case TrafficLightState_Yellow:
        m._TrafficLightState_Yellow_(e)
    case TrafficLightState_FlashingRed:
        m._TrafficLightState_FlashingRed_(e)
    case TrafficLightState_End:
        m._TrafficLightState_End_(e)
    case TrafficLightState_Working:
        m._TrafficLightState_Working_(e)
    }
    
    if m._nextCompartment_ != nil {
        nextCompartment := m._nextCompartment_
        m._nextCompartment_ = nil
        if nextCompartment._forwardEvent_ != nil && 
           nextCompartment._forwardEvent_.Msg == ">" {
            m._mux_(&FrameEvent{Msg: "<", Params: m._compartment_.ExitArgs, Ret: nil})
            m._compartment_ = nextCompartment
            m._mux_(nextCompartment._forwardEvent_)
        } else {
            m._do_transition_(nextCompartment)
            if nextCompartment._forwardEvent_ != nil {
                m._mux_(nextCompartment._forwardEvent_)
            }
        }
        nextCompartment._forwardEvent_ = nil
    }
}

//===================== Machine Block ===================//

func (m *trafficLightStruct) _TrafficLightState_Begin_(e *FrameEvent) {
    switch e.Msg {
    case "start":
        m.initTrafficLight()
        m.startWorkingTimer()
        compartment := NewTrafficLightCompartment(TrafficLightState_Red)
        m._transition_(compartment)
        return
    }
}

func (m *trafficLightStruct) _TrafficLightState_Red_(e *FrameEvent) {
    switch e.Msg {
    case ">":
        m.enterRed()
        return
    case "tick":
        compartment := NewTrafficLightCompartment(TrafficLightState_Green)
        m._transition_(compartment)
        return
    }
    m._TrafficLightState_Working_(e)
    
}

func (m *trafficLightStruct) _TrafficLightState_Green_(e *FrameEvent) {
    switch e.Msg {
    case ">":
        m.enterGreen()
        return
    case "tick":
        compartment := NewTrafficLightCompartment(TrafficLightState_Yellow)
        m._transition_(compartment)
        return
    }
    m._TrafficLightState_Working_(e)
    
}

func (m *trafficLightStruct) _TrafficLightState_Yellow_(e *FrameEvent) {
    switch e.Msg {
    case ">":
        m.enterYellow()
        return
    case "tick":
        compartment := NewTrafficLightCompartment(TrafficLightState_Red)
        m._transition_(compartment)
        return
    }
    m._TrafficLightState_Working_(e)
    
}

func (m *trafficLightStruct) _TrafficLightState_FlashingRed_(e *FrameEvent) {
    switch e.Msg {
    case ">":
        m.enterFlashingRed()
        m.stopWorkingTimer()
        m.startFlashingTimer()
        return
    case "<":
        m.stopFlashingTimer()
        m.startWorkingTimer()
        return
    case "tick":
        m.changeFlashingAnimation()
        return
    case "changeColor":
        m.flashColor = e.Params["color"].(string)
        return
    case "systemRestart":
        compartment := NewTrafficLightCompartment(TrafficLightState_Red)
        m._transition_(compartment)
        return
    case "stop":
        compartment := NewTrafficLightCompartment(TrafficLightState_End)
        m._transition_(compartment)
        return
    case "getColor":
        e.Ret = m.flashColor
        return
        
    }
}

func (m *trafficLightStruct) _TrafficLightState_End_(e *FrameEvent) {
    switch e.Msg {
    case ">":
        m.flashColor = ""
        m.stopWorkingTimer()
        m.destroyTrafficLight()
        return
    }
}

func (m *trafficLightStruct) _TrafficLightState_Working_(e *FrameEvent) {
    switch e.Msg {
    case "stop":
        compartment := NewTrafficLightCompartment(TrafficLightState_End)
        m._transition_(compartment)
        return
    case "systemError":
        compartment := NewTrafficLightCompartment(TrafficLightState_FlashingRed)
        m._transition_(compartment)
        return
    case "changeColor":
        m.flashColor = e.Params["color"].(string)
        return
    case "getColor":
        e.Ret = m.flashColor
        return
        
    }
}

//=============== Machinery and Mechanisms ==============//

func (m *trafficLightStruct) _transition_(compartment *TrafficLightCompartment) {
    m._nextCompartment_ = compartment
}

func (m *trafficLightStruct) _do_transition_(nextCompartment *TrafficLightCompartment) {
    m._mux_(&FrameEvent{Msg: "<", Params: m._compartment_.ExitArgs, Ret: nil})
    m._compartment_ = nextCompartment
    m._mux_(&FrameEvent{Msg: ">", Params: m._compartment_.EnterArgs, Ret: nil})
}

//===================== Actions Block ===================//


/********************************************************

// Unimplemented Actions

func (m *trafficLightStruct) initTrafficLight()  {}
func (m *trafficLightStruct) enterRed()  {}
func (m *trafficLightStruct) enterGreen()  {}
func (m *trafficLightStruct) enterYellow()  {}
func (m *trafficLightStruct) enterFlashingRed()  {}
func (m *trafficLightStruct) startWorkingTimer()  {}
func (m *trafficLightStruct) stopWorkingTimer()  {}
func (m *trafficLightStruct) startFlashingTimer()  {}
func (m *trafficLightStruct) stopFlashingTimer()  {}
func (m *trafficLightStruct) changeFlashingAnimation()  {}
func (m *trafficLightStruct) destroyTrafficLight()  {}

********************************************************/

//=============== Compartment ==============//

type TrafficLightCompartment struct {
    State TrafficLightState
    StateArgs map[string]interface{}
    StateVars map[string]interface{}
    EnterArgs map[string]interface{}
    ExitArgs map[string]interface{}
    _forwardEvent_ *FrameEvent
}

func NewTrafficLightCompartment(state TrafficLightState) *TrafficLightCompartment {
    c := &TrafficLightCompartment{State: state}
    c.StateArgs = make(map[string]interface{})
    c.StateVars = make(map[string]interface{})
    c.EnterArgs = make(map[string]interface{})
    c.ExitArgs = make(map[string]interface{})
    return c
}
