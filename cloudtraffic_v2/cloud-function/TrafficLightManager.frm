```
package trafficlight

import (
	"github.com/frame-lang/cloudtraffic/cloudtraffic_v2/framelang"
)
```

#TrafficLightManager >[createWorkflow:bool]

    -interface-

    initTrafficLight
    tick
    end
    enterRed
    enterGreen
    enterYellow
    enterFlashingRed
    startWorkingTimer
    stopWorkingTimer
    startFlashingTimer
    stopFlashingTimer
    startFlashing
    changeFlashingAnimation
    systemError
    systemRestart
    destroyTrafficLight
    connectionClosed

    -machine-

    $Start
        |>|[createWorkflow:bool] 
            createWorkflow ? 
            	-> "Create\nWorkflow" $Create 
            : 
            	-> "Load\nWorkflow" $Load 
            :: ^

	$Create => $HandleExternalEvents
    	|>|
            trafficLight = NewTrafficLight(#)
            trafficLight.Start() 
             -> "Created" $Save ^
            
    $Load => $HandleExternalEvents
    	|>|
            var workflowData = getWorkflowFromRedis()
            trafficLight = LoadTrafficLight(# workflowData)
            ->> "Loaded" $Working ^
 
 	$Working => $HandleExternalEvents
    
    $Save
        |>|
            var workflowData = trafficLight.Marshal() 
            setWorkflowInRedis(workflowData)
            trafficLight = nil 
            -> "Stop" $Stop ^
            
    $Stop

        
    $HandleExternalEvents => $HandleControllerEvents
        |tick|
            trafficLight.Tick() -> "Tick" $Save ^
        |systemError|
            trafficLight.SystemError() -> "System\nError" $Save ^
        |systemRestart|
            trafficLight.SystemRestart() -> "System\nRestart" $Save ^
        |end|
            trafficLight.Stop() -> "End" $Save ^
        |connectionClosed|
            removeWorkflowFromRedis() -> "Connection Closed" $Stop ^
 
    $HandleControllerEvents
        |initTrafficLight| initTrafficLight() ^
        |enterRed| enterRed() ^
        |enterGreen| enterGreen()  ^
        |enterYellow| enterYellow() ^
        |enterFlashingRed| enterFlashingRed() ^
        |startWorkingTimer| startWorkingTimer() ^
        |stopWorkingTimer| stopWorkingTimer() ^
        |startFlashingTimer| startFlashingTimer() ^
        |stopFlashingTimer| stopFlashingTimer() ^
        |changeFlashingAnimation| changeFlashingAnimation() ^
        |destroyTrafficLight| destroyTrafficLight() ^

    -actions-

    enterRed
    enterGreen
    enterYellow
    enterFlashingRed
    startWorkingTimer
    stopWorkingTimer
    startFlashingTimer
    stopFlashingTimer   
    initTrafficLight
    changeFlashingAnimation
    destroyTrafficLight
    removeWorkflowFromRedis
    setWorkflowInRedis [data:`[]byte`]
    getWorkflowFromRedis: `[]byte`

    -domain-

    var trafficLight:TrafficLight = null
##