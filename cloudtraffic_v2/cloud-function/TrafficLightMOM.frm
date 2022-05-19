```
package trafficlight

import (
	"github.com/frame-lang/cloudtraffic/cloudtraffic_v2/framelang"
)
```

#TrafficLightMom

    -interface-

    start @(|>>|)
    stop
    initTrafficLight
    tick
    init
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

    $Entry => $TrafficLightApi
        |init|
            trafficLight = NewTrafficLight(#)
            trafficLight.Start()
            -> "Saving" $Save ^
        |tick|
            trafficLight.Tick() -> "Done" $Save ^
        |systemError|
            trafficLight.SystemError() -> "Done" $Save ^
        |systemRestart|
            trafficLight.SystemError() -> "Done" $Save ^
        |end|
            trafficLight.Stop() -> "Done" $Save ^

    $Save
        |>|
            var jsonData = trafficLight.Marshal() 
            saveInDisk(jsonData)
            trafficLight = nil ^

    $TrafficLightApi
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
    saveInDisk [data:`[]byte`]
    loadFromDisk [clientId: string]: `[]byte`

    -domain-

    var trafficLight:TrafficLight = null
##