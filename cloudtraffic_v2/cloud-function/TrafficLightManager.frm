```
package trafficlight

import (
	"github.com/frame-lang/cloudtraffic/cloudtraffic_v2/framelang"
)
```

#TrafficLightManager >[isNewWorkflow:bool]

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

    -machine-

    $Start => $HandleExternalEvents
        |>|[isNewWorkflow:bool] 
            isNewWorkflow ? 
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
            var savedData = getFromRedis()
            trafficLight = LoadTrafficLight(# savedData)
            ->> "Loaded" $Working ^
 
 	$Working => $HandleExternalEvents
    
    $Save
        |>|
            var jsonData = trafficLight.Marshal() 
            setInRedis(jsonData)
            trafficLight = nil 
            -> "Stop" $Stop ^
            
    $Stop => $HandleExternalEvents

        
    $HandleExternalEvents => $HandleControllerEvents
        |tick|
            trafficLight.Tick() -> "Tick" $Save ^
        |systemError|
            trafficLight.SystemError() -> "System\nError" $Save ^
        |systemRestart|
            trafficLight.SystemRestart() -> "System\nRestart" $Save ^
        |end|
            trafficLight.Stop() -> "End" $Save ^
 
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
    setInRedis [data:`[]byte`]
    getFromRedis: `[]byte`

    -domain-

    var trafficLight:TrafficLight = null
##