```
package trafficlight

import (
	"github.com/frame-lang/cloudtraffic/cloudtraffic_v2/framelang"
)
```

#TrafficLightMom >[isInit:bool]

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

    $Entry => $TrafficLightApi
        |>|[isInit:bool] 
            isInit? 
                trafficLight = NewTrafficLight(#)
                trafficLight.Start()
                -> "Saving" $Save ^
            :
                var savedData = getFromRedis()
                trafficLight = LoadTrafficLight(# savedData) ^
            ::^
        |tick|
            trafficLight.Tick() -> "Done" $Save ^
        |systemError|
            trafficLight.SystemError() -> "Done" $Save ^
        |systemRestart|
            trafficLight.SystemRestart() -> "Done" $Save ^
        |end|
            trafficLight.Stop() -> "Done" $Save ^

    $Save
        |>|
            var jsonData = trafficLight.Marshal() 
            setInRedis(jsonData)
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
    setInRedis [data:`[]byte`]
    getFromRedis: `[]byte`

    -domain-

    var trafficLight:TrafficLight = null
##