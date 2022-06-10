// emitted from framec_v0.10.0
// get include files at https://github.com/frame-lang/frame-ancillary-files
package trafficlight

func NewTrafficLightManager(createWorkflow bool) TrafficLightManager {
    m := &trafficLightManagerStruct{}
    
    // Validate interfaces
    var _ TrafficLightManager = m
    var _ TrafficLightManager_actions = m
    
    // Create and intialize start state compartment.
    m._compartment_ = NewTrafficLightManagerCompartment(TrafficLightManagerState_Start)
    m._compartment_.EnterArgs["createWorkflow"] = createWorkflow
    
    // Override domain variables.
    m.trafficLight = nil
    
    // Send system start event
    e := FrameEvent{Msg:">", Params:m._compartment_.EnterArgs}
    m._mux_(&e)
    return m
}


type TrafficLightManagerState uint

const (
    TrafficLightManagerState_Start TrafficLightManagerState = iota
    TrafficLightManagerState_Create
    TrafficLightManagerState_Load
    TrafficLightManagerState_Working
    TrafficLightManagerState_Save
    TrafficLightManagerState_Stop
    TrafficLightManagerState_HandleExternalEvents
    TrafficLightManagerState_HandleControllerEvents
)

type TrafficLightManager interface {
    InitTrafficLight() 
    Tick() 
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
    ConnectionClosed() 
}

type TrafficLightManager_actions interface {
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
    removeWorkflowFromRedis() 
    setWorkflowInRedis(data []byte) 
    getWorkflowFromRedis() []byte
}


type trafficLightManagerStruct struct {
    _compartment_ *TrafficLightManagerCompartment
    _nextCompartment_ *TrafficLightManagerCompartment
    trafficLight TrafficLight
}

//===================== Interface Block ===================//

func (m *trafficLightManagerStruct) InitTrafficLight()  {
    e := FrameEvent{Msg:"initTrafficLight"}
    m._mux_(&e)
}

func (m *trafficLightManagerStruct) Tick()  {
    e := FrameEvent{Msg:"tick"}
    m._mux_(&e)
}

func (m *trafficLightManagerStruct) End()  {
    e := FrameEvent{Msg:"end"}
    m._mux_(&e)
}

func (m *trafficLightManagerStruct) EnterRed()  {
    e := FrameEvent{Msg:"enterRed"}
    m._mux_(&e)
}

func (m *trafficLightManagerStruct) EnterGreen()  {
    e := FrameEvent{Msg:"enterGreen"}
    m._mux_(&e)
}

func (m *trafficLightManagerStruct) EnterYellow()  {
    e := FrameEvent{Msg:"enterYellow"}
    m._mux_(&e)
}

func (m *trafficLightManagerStruct) EnterFlashingRed()  {
    e := FrameEvent{Msg:"enterFlashingRed"}
    m._mux_(&e)
}

func (m *trafficLightManagerStruct) StartWorkingTimer()  {
    e := FrameEvent{Msg:"startWorkingTimer"}
    m._mux_(&e)
}

func (m *trafficLightManagerStruct) StopWorkingTimer()  {
    e := FrameEvent{Msg:"stopWorkingTimer"}
    m._mux_(&e)
}

func (m *trafficLightManagerStruct) StartFlashingTimer()  {
    e := FrameEvent{Msg:"startFlashingTimer"}
    m._mux_(&e)
}

func (m *trafficLightManagerStruct) StopFlashingTimer()  {
    e := FrameEvent{Msg:"stopFlashingTimer"}
    m._mux_(&e)
}

func (m *trafficLightManagerStruct) StartFlashing()  {
    e := FrameEvent{Msg:"startFlashing"}
    m._mux_(&e)
}

func (m *trafficLightManagerStruct) ChangeFlashingAnimation()  {
    e := FrameEvent{Msg:"changeFlashingAnimation"}
    m._mux_(&e)
}

func (m *trafficLightManagerStruct) SystemError()  {
    e := FrameEvent{Msg:"systemError"}
    m._mux_(&e)
}

func (m *trafficLightManagerStruct) SystemRestart()  {
    e := FrameEvent{Msg:"systemRestart"}
    m._mux_(&e)
}

func (m *trafficLightManagerStruct) DestroyTrafficLight()  {
    e := FrameEvent{Msg:"destroyTrafficLight"}
    m._mux_(&e)
}

func (m *trafficLightManagerStruct) ConnectionClosed()  {
    e := FrameEvent{Msg:"connectionClosed"}
    m._mux_(&e)
}

//====================== Multiplexer ====================//

func (m *trafficLightManagerStruct) _mux_(e *FrameEvent) {
    switch m._compartment_.State {
    case TrafficLightManagerState_Start:
        m._TrafficLightManagerState_Start_(e)
    case TrafficLightManagerState_Create:
        m._TrafficLightManagerState_Create_(e)
    case TrafficLightManagerState_Load:
        m._TrafficLightManagerState_Load_(e)
    case TrafficLightManagerState_Working:
        m._TrafficLightManagerState_Working_(e)
    case TrafficLightManagerState_Save:
        m._TrafficLightManagerState_Save_(e)
    case TrafficLightManagerState_Stop:
        m._TrafficLightManagerState_Stop_(e)
    case TrafficLightManagerState_HandleExternalEvents:
        m._TrafficLightManagerState_HandleExternalEvents_(e)
    case TrafficLightManagerState_HandleControllerEvents:
        m._TrafficLightManagerState_HandleControllerEvents_(e)
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

func (m *trafficLightManagerStruct) _TrafficLightManagerState_Start_(e *FrameEvent) {
    switch e.Msg {
    case ">":
        if e.Params["createWorkflow"].(bool) {
            // Create\nWorkflow
            compartment := NewTrafficLightManagerCompartment(TrafficLightManagerState_Create)
            m._transition_(compartment)
        } else {
            // Load\nWorkflow
            compartment := NewTrafficLightManagerCompartment(TrafficLightManagerState_Load)
            m._transition_(compartment)
        }
        return
    }
}

func (m *trafficLightManagerStruct) _TrafficLightManagerState_Create_(e *FrameEvent) {
    switch e.Msg {
    case ">":
        m.trafficLight = NewTrafficLight(m)
        // Created
        compartment := NewTrafficLightManagerCompartment(TrafficLightManagerState_Save)
        m._transition_(compartment)
        return
    }
    m._TrafficLightManagerState_HandleExternalEvents_(e)
    
}

func (m *trafficLightManagerStruct) _TrafficLightManagerState_Load_(e *FrameEvent) {
    switch e.Msg {
    case ">":
        var workflowData  = m.getWorkflowFromRedis()
        m.trafficLight = LoadTrafficLight(m,workflowData)
        compartment := NewTrafficLightManagerCompartment(TrafficLightManagerState_Working)
        m._changeState_(compartment)
        return
    }
    m._TrafficLightManagerState_HandleExternalEvents_(e)
    
}

func (m *trafficLightManagerStruct) _TrafficLightManagerState_Working_(e *FrameEvent) {
    switch e.Msg {
    }
    m._TrafficLightManagerState_HandleExternalEvents_(e)
    
}

func (m *trafficLightManagerStruct) _TrafficLightManagerState_Save_(e *FrameEvent) {
    switch e.Msg {
    case ">":
        var workflowData  = m.trafficLight.Marshal()
        m.setWorkflowInRedis(workflowData)
        m.trafficLight = nil
        // Stop
        compartment := NewTrafficLightManagerCompartment(TrafficLightManagerState_Stop)
        m._transition_(compartment)
        return
    }
}

func (m *trafficLightManagerStruct) _TrafficLightManagerState_Stop_(e *FrameEvent) {
    switch e.Msg {
    }
}

func (m *trafficLightManagerStruct) _TrafficLightManagerState_HandleExternalEvents_(e *FrameEvent) {
    switch e.Msg {
    case "tick":
        m.trafficLight.Tick()
        // Tick
        compartment := NewTrafficLightManagerCompartment(TrafficLightManagerState_Save)
        m._transition_(compartment)
        return
    case "systemError":
        m.trafficLight.SystemError()
        // System\nError
        compartment := NewTrafficLightManagerCompartment(TrafficLightManagerState_Save)
        m._transition_(compartment)
        return
    case "systemRestart":
        m.trafficLight.SystemRestart()
        // System\nRestart
        compartment := NewTrafficLightManagerCompartment(TrafficLightManagerState_Save)
        m._transition_(compartment)
        return
    case "end":
        m.trafficLight.Stop()
        // End
        compartment := NewTrafficLightManagerCompartment(TrafficLightManagerState_Save)
        m._transition_(compartment)
        return
    case "connectionClosed":
        m.removeWorkflowFromRedis()
        // Connection Closed
        compartment := NewTrafficLightManagerCompartment(TrafficLightManagerState_Stop)
        m._transition_(compartment)
        return
    }
    m._TrafficLightManagerState_HandleControllerEvents_(e)
    
}

func (m *trafficLightManagerStruct) _TrafficLightManagerState_HandleControllerEvents_(e *FrameEvent) {
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

func (m *trafficLightManagerStruct) _transition_(compartment *TrafficLightManagerCompartment) {
    m._nextCompartment_ = compartment
}

func (m *trafficLightManagerStruct) _do_transition_(nextCompartment *TrafficLightManagerCompartment) {
    m._mux_(&FrameEvent{Msg: "<", Params: m._compartment_.ExitArgs, Ret: nil})
    m._compartment_ = nextCompartment
    m._mux_(&FrameEvent{Msg: ">", Params: m._compartment_.EnterArgs, Ret: nil})
}

func (m *trafficLightManagerStruct) _changeState_(compartment *TrafficLightManagerCompartment) {
    m._compartment_ = compartment
}

//===================== Actions Block ===================//


/********************************************************

// Unimplemented Actions

func (m *trafficLightManagerStruct) enterRed()  {}
func (m *trafficLightManagerStruct) enterGreen()  {}
func (m *trafficLightManagerStruct) enterYellow()  {}
func (m *trafficLightManagerStruct) enterFlashingRed()  {}
func (m *trafficLightManagerStruct) startWorkingTimer()  {}
func (m *trafficLightManagerStruct) stopWorkingTimer()  {}
func (m *trafficLightManagerStruct) startFlashingTimer()  {}
func (m *trafficLightManagerStruct) stopFlashingTimer()  {}
func (m *trafficLightManagerStruct) initTrafficLight()  {}
func (m *trafficLightManagerStruct) changeFlashingAnimation()  {}
func (m *trafficLightManagerStruct) destroyTrafficLight()  {}
func (m *trafficLightManagerStruct) removeWorkflowFromRedis()  {}
func (m *trafficLightManagerStruct) setWorkflowInRedis(data []byte)  {}
func (m *trafficLightManagerStruct) getWorkflowFromRedis() []byte {}

********************************************************/

//=============== Compartment ==============//

type TrafficLightManagerCompartment struct {
    State TrafficLightManagerState
    StateArgs map[string]interface{}
    StateVars map[string]interface{}
    EnterArgs map[string]interface{}
    ExitArgs map[string]interface{}
    _forwardEvent_ *FrameEvent
}

func NewTrafficLightManagerCompartment(state TrafficLightManagerState) *TrafficLightManagerCompartment {
    c := &TrafficLightManagerCompartment{State: state}
    c.StateArgs = make(map[string]interface{})
    c.StateVars = make(map[string]interface{})
    c.EnterArgs = make(map[string]interface{})
    c.ExitArgs = make(map[string]interface{})
    return c
}
