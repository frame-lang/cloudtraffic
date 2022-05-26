```
package trafficlight

import (
	"github.com/frame-lang/cloudtraffic/cloudtraffic_v1/framelang"
	"github.com/gorilla/websocket"
)
```

#TrafficLightMom[clientId: string connection: `*websocket.Conn`]

    -interface-

    start @(|>>|)
    stop
    initTrafficLight
    tick    
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
    getConnection:`*websocket.Conn`

    -machine-

    $New => $TrafficLightApi
        |>>|
            trafficLight = NewTrafficLight(#)
            trafficLight.Start()
            -> "Traffic Light\nStarted" $Saving ^
        |getConnection|:`*websocket.Conn` @^ = connection ^
 
    $Saving 
        |>|
            var jsonData = trafficLight.Marshal() 
            saveInDisk(jsonData)
            trafficLight = nil
            -> "Saved" $Persisted ^

    $Persisted 
        |tick| -> "Tick"  =>  $Working ^
        |systemError| -> "System Error" =>  $Working ^
        |systemRestart| -> "System Error" =>  $Working ^
        |getConnection|:`*websocket.Conn` @^ = connection ^
        |stop| -> "Stop" $End ^

    $Working => $TrafficLightApi
        |>|
            var savedData = loadFromDisk(#.clientId)
            trafficLight = LoadTrafficLight(# savedData) ^
        |tick| trafficLight.Tick() -> "Done" $Saving ^
        |systemError| trafficLight.SystemError() -> "Done" $Saving ^
        |systemRestart| trafficLight.SystemRestart() -> "Done" $Saving ^

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

    $End => $TrafficLightApi
        |>|
            var savedData = loadFromDisk(#.clientId)
            trafficLight = LoadTrafficLight(# savedData)
            trafficLight.Stop()
            var jsonData = trafficLight.Marshal() 
            saveInDisk(jsonData)
            trafficLight = nil
            -> $New ^

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
    var connection:`*websocket.Conn` = connection
    var clientId:string = clientId
    var stopper:`chan<- bool` = null
##