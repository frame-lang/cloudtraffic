// emitted from framec_v0.8.0
// get include files at https://github.com/frame-lang/frame-ancillary-files
package trafficlight

func NewTrafficLightMom(isInit bool) TrafficLightMom {
    m := &trafficLightMomStruct{}
    
    // Validate interfaces
    var _ TrafficLightMom = m
    var _ TrafficLightMom_actions = m
    m._compartment_ = NewTrafficLightMomCompartment(TrafficLightMomState_Entry)
    
    // Initialize domain
    m.trafficLight = nil
    
    // Send system start event
    params := make(map[string]interface{})
    params["isInit"] = isInit
    e := FrameEvent{Msg:">", Params:params}
    m._mux_(&e)
    return m
}


type TrafficLightMomState uint

const (
    TrafficLightMomState_Entry TrafficLightMomState = iota
    TrafficLightMomState_Save
    TrafficLightMomState_TrafficLightApi
)

type TrafficLightMom interface {
    InitTrafficLight() 
    Tick() 
    Init() 
    End() 
    EnterRed() 
    EnterGreen() 
    EnterYellow() 
    EnterFlashingRed() 
    StartWorkingTimer() 
    StopWorkingTimer() 
    StartFlashingTimer() 
    StopFlashingTimer() 
    StartFlashing() 
    ChangeFlashingAnimation() 
    SystemError() 
    SystemRestart() 
    DestroyTrafficLight() 
}

type TrafficLightMom_actions interface {
    enterRed() 
    enterGreen() 
    enterYellow() 
    enterFlashingRed() 
    startWorkingTimer() 
    stopWorkingTimer() 
    startFlashingTimer() 
    stopFlashingTimer() 
    initTrafficLight() 
    changeFlashingAnimation() 
    destroyTrafficLight() 
    saveInDisk(data []byte) 
    loadFromDisk() []byte
}


type trafficLightMomStruct struct {
    _compartment_ *TrafficLightMomCompartment
    _nextCompartment_ *TrafficLightMomCompartment
    trafficLight TrafficLight
}

//===================== Interface Block ===================//

func (m *trafficLightMomStruct) InitTrafficLight()  {
    e := FrameEvent{Msg:"initTrafficLight"}
    m._mux_(&e)
}

func (m *trafficLightMomStruct) Tick()  {
    e := FrameEvent{Msg:"tick"}
    m._mux_(&e)
}

func (m *trafficLightMomStruct) Init()  {
    e := FrameEvent{Msg:"init"}
    m._mux_(&e)
}

func (m *trafficLightMomStruct) End()  {
    e := FrameEvent{Msg:"end"}
    m._mux_(&e)
}

func (m *trafficLightMomStruct) EnterRed()  {
    e := FrameEvent{Msg:"enterRed"}
    m._mux_(&e)
}

func (m *trafficLightMomStruct) EnterGreen()  {
    e := FrameEvent{Msg:"enterGreen"}
    m._mux_(&e)
}

func (m *trafficLightMomStruct) EnterYellow()  {
    e := FrameEvent{Msg:"enterYellow"}
    m._mux_(&e)
}

func (m *trafficLightMomStruct) EnterFlashingRed()  {
    e := FrameEvent{Msg:"enterFlashingRed"}
    m._mux_(&e)
}

func (m *trafficLightMomStruct) StartWorkingTimer()  {
    e := FrameEvent{Msg:"startWorkingTimer"}
    m._mux_(&e)
}

func (m *trafficLightMomStruct) StopWorkingTimer()  {
    e := FrameEvent{Msg:"stopWorkingTimer"}
    m._mux_(&e)
}

func (m *trafficLightMomStruct) StartFlashingTimer()  {
    e := FrameEvent{Msg:"startFlashingTimer"}
    m._mux_(&e)
}

func (m *trafficLightMomStruct) StopFlashingTimer()  {
    e := FrameEvent{Msg:"stopFlashingTimer"}
    m._mux_(&e)
}

func (m *trafficLightMomStruct) StartFlashing()  {
    e := FrameEvent{Msg:"startFlashing"}
    m._mux_(&e)
}

func (m *trafficLightMomStruct) ChangeFlashingAnimation()  {
    e := FrameEvent{Msg:"changeFlashingAnimation"}
    m._mux_(&e)
}

func (m *trafficLightMomStruct) SystemError()  {
    e := FrameEvent{Msg:"systemError"}
    m._mux_(&e)
}

func (m *trafficLightMomStruct) SystemRestart()  {
    e := FrameEvent{Msg:"systemRestart"}
    m._mux_(&e)
}

func (m *trafficLightMomStruct) DestroyTrafficLight()  {
    e := FrameEvent{Msg:"destroyTrafficLight"}
    m._mux_(&e)
}

//====================== Multiplexer ====================//

func (m *trafficLightMomStruct) _mux_(e *FrameEvent) {
    switch m._compartment_.State {
    case TrafficLightMomState_Entry:
        m._TrafficLightMomState_Entry_(e)
    case TrafficLightMomState_Save:
        m._TrafficLightMomState_Save_(e)
    case TrafficLightMomState_TrafficLightApi:
        m._TrafficLightMomState_TrafficLightApi_(e)
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

func (m *trafficLightMomStruct) _TrafficLightMomState_Entry_(e *FrameEvent) {
    switch e.Msg {
    case ">":
        if e.Params["isInit"].(bool) {
            return
        }
        var savedData  = m.loadFromDisk()
        m.trafficLight = LoadTrafficLight(m,savedData)
        return
    case "init":
        m.trafficLight = NewTrafficLight(m)
        m.trafficLight.Start()
        // Saving
        compartment := NewTrafficLightMomCompartment(TrafficLightMomState_Save)
        m._transition_(compartment)
        return
    case "tick":
        m.trafficLight.Tick()
        // Done
        compartment := NewTrafficLightMomCompartment(TrafficLightMomState_Save)
        m._transition_(compartment)
        return
    case "systemError":
        m.trafficLight.SystemError()
        // Done
        compartment := NewTrafficLightMomCompartment(TrafficLightMomState_Save)
        m._transition_(compartment)
        return
    case "systemRestart":
        m.trafficLight.SystemRestart()
        // Done
        compartment := NewTrafficLightMomCompartment(TrafficLightMomState_Save)
        m._transition_(compartment)
        return
    case "end":
        m.trafficLight.Stop()
        // Done
        compartment := NewTrafficLightMomCompartment(TrafficLightMomState_Save)
        m._transition_(compartment)
        return
    }
    m._TrafficLightMomState_TrafficLightApi_(e)
    
}

func (m *trafficLightMomStruct) _TrafficLightMomState_Save_(e *FrameEvent) {
    switch e.Msg {
    case ">":
        var jsonData  = m.trafficLight.Marshal()
        m.saveInDisk(jsonData)
        m.trafficLight = nil
        return
    }
}

func (m *trafficLightMomStruct) _TrafficLightMomState_TrafficLightApi_(e *FrameEvent) {
    switch e.Msg {
    case "initTrafficLight":
        m.initTrafficLight()
        return
    case "enterRed":
        m.enterRed()
        return
    case "enterGreen":
        m.enterGreen()
        return
    case "enterYellow":
        m.enterYellow()
        return
    case "enterFlashingRed":
        m.enterFlashingRed()
        return
    case "startWorkingTimer":
        m.startWorkingTimer()
        return
    case "stopWorkingTimer":
        m.stopWorkingTimer()
        return
    case "startFlashingTimer":
        m.startFlashingTimer()
        return
    case "stopFlashingTimer":
        m.stopFlashingTimer()
        return
    case "changeFlashingAnimation":
        m.changeFlashingAnimation()
        return
    case "destroyTrafficLight":
        m.destroyTrafficLight()
        return
    }
}

//=============== Machinery and Mechanisms ==============//

func (m *trafficLightMomStruct) _transition_(compartment *TrafficLightMomCompartment) {
    m._nextCompartment_ = compartment
}

func (m *trafficLightMomStruct) _do_transition_(nextCompartment *TrafficLightMomCompartment) {
    m._mux_(&FrameEvent{Msg: "<", Params: m._compartment_.ExitArgs, Ret: nil})
    m._compartment_ = nextCompartment
    m._mux_(&FrameEvent{Msg: ">", Params: m._compartment_.EnterArgs, Ret: nil})
}

//===================== Actions Block ===================//


/********************************************************

// Unimplemented Actions

func (m *trafficLightMomStruct) enterRed()  {}
func (m *trafficLightMomStruct) enterGreen()  {}
func (m *trafficLightMomStruct) enterYellow()  {}
func (m *trafficLightMomStruct) enterFlashingRed()  {}
func (m *trafficLightMomStruct) startWorkingTimer()  {}
func (m *trafficLightMomStruct) stopWorkingTimer()  {}
func (m *trafficLightMomStruct) startFlashingTimer()  {}
func (m *trafficLightMomStruct) stopFlashingTimer()  {}
func (m *trafficLightMomStruct) initTrafficLight()  {}
func (m *trafficLightMomStruct) changeFlashingAnimation()  {}
func (m *trafficLightMomStruct) destroyTrafficLight()  {}
func (m *trafficLightMomStruct) saveInDisk(data []byte)  {}
func (m *trafficLightMomStruct) loadFromDisk() []byte {}

********************************************************/

//=============== Compartment ==============//

type TrafficLightMomCompartment struct {
    State TrafficLightMomState
    StateArgs map[string]interface{}
    StateVars map[string]interface{}
    EnterArgs map[string]interface{}
    ExitArgs map[string]interface{}
    _forwardEvent_ *FrameEvent
}

func NewTrafficLightMomCompartment(state TrafficLightMomState) *TrafficLightMomCompartment {
    c := &TrafficLightMomCompartment{State: state}
    c.StateArgs = make(map[string]interface{})
    c.StateVars = make(map[string]interface{})
    c.EnterArgs = make(map[string]interface{})
    c.ExitArgs = make(map[string]interface{})
    return c
}
