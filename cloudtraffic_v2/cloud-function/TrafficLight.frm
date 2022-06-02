```
package trafficlight

import (
	"encoding/json"

	"github.com/frame-lang/cloudtraffic/cloudtraffic_v2/framelang"
)
```
#[derive(Marshal)]
#[managed(TrafficLightManager)]

#TrafficLight

    -interface-

    start @(|>>|)
    stop 
    tick
    systemError
    systemRestart
    changeColor[color:string]
    getColor:string

    -machine-

    $Begin
        |>>|
            initTrafficLight()
            startWorkingTimer()
            -> $Red ^

    $Red => $Working
        |>|
            enterRed() ^
        |tick|
            -> $Green ^

    $Green => $Working
        |>|
            enterGreen() ^
        |tick|
            -> $Yellow ^

    $Yellow => $Working
        |>|
            enterYellow() ^
        |tick|
            -> $Red ^

    $FlashingRed
        |>|
            enterFlashingRed()
            stopWorkingTimer()
            startFlashingTimer() ^
        |<|
            stopFlashingTimer()
            startWorkingTimer() ^
        |tick|
            changeFlashingAnimation() ^
        |changeColor|[color:string]
            flashColor = color ^
        |systemRestart|
            -> $Red  ^
        |stop|
            -> $End ^
        |getColor|:string @^ = flashColor ^
    
    $End
        |>|
            flashColor = ""
            stopWorkingTimer()
            destroyTrafficLight() ^

    $Working
        |stop|
            -> $End ^
        |systemError|
            -> $FlashingRed ^
        |changeColor|[color:string]
            flashColor = color ^
        |getColor|:string @^ = flashColor ^

    -actions-

    initTrafficLight
    enterRed
    enterGreen
    enterYellow
    enterFlashingRed
    startWorkingTimer
    stopWorkingTimer
    startFlashingTimer
    stopFlashingTimer
    changeFlashingAnimation
    destroyTrafficLight

    -domain-

    var flashColor:string = ""

##