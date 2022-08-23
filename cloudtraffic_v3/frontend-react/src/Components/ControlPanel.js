import { useState } from 'react';
import { Button, Slider } from '@mui/material';
import { isMobile } from 'react-device-detect';
import './ControlPanel.css'

import {
    SLIDER_MARKS,
    DEFAULT_WOKRING_INTERVAL,
    DEFAULT_FLASHING_INTERVAL
} from '../Utils/Constants';

export default function ControlPanel({state, sendMessage, connectionStatus}) {
    const [workingInterval, setWorkingInterval] = useState(DEFAULT_WOKRING_INTERVAL);
    const [flashingInterval, setFlashingInterval] = useState(DEFAULT_FLASHING_INTERVAL);

    const handleSystemError = () => {
        sendMessage(JSON.stringify({event: 'error'}))
      }
    
      const handleSystemRestart = () => {
        sendMessage(JSON.stringify({event: 'restart'}))
      }
    
      const handleWorkingTimerChange = (event, newValue) => {
        if (typeof newValue === 'number') {
          setWorkingInterval(newValue)
          sendMessage(JSON.stringify({
            event: 'updateWorkingTimer',
            state: state.name,
            data: `${newValue}s`
          }))
        }
      };
    
      const handleErrorTimerChange = (event, newValue) => {
        if (typeof newValue === 'number') {
          setFlashingInterval(newValue)
          sendMessage(JSON.stringify({
            event: 'updateFlashingTimer',
            state: state.name,
            data: `${newValue}s`
          }))
        }
      };
    
      const handleResetTimers = () => {
        setWorkingInterval(DEFAULT_WOKRING_INTERVAL)
        setFlashingInterval(DEFAULT_FLASHING_INTERVAL)
        sendMessage(JSON.stringify({
          event: 'updateWorkingTimer',
          state: state.name,
          data: `${DEFAULT_WOKRING_INTERVAL}s`
        }))
        sendMessage(JSON.stringify({
          event: 'updateFlashingTimer',
          state: state.name,
          data: `${DEFAULT_FLASHING_INTERVAL}s`
        }))
      }

    return (
        <>
            <div className='control-section-items'>
                <div>
                <span>Connection Status: </span>
                <span className={`connection-status ${(connectionStatus === 'Open') ? 'connection-status-open': 'connection-status-closed'}`}>{connectionStatus}</span>
                </div>
    
                <Slider
                    aria-label="Always visible"
                    defaultValue={DEFAULT_WOKRING_INTERVAL}
                    step={0.25}
                    marks={SLIDER_MARKS}
                    value={workingInterval}
                    valueLabelDisplay={isMobile ? 'off' : 'on'}
                    min={0.25}
                    max={5}
                    onChange={handleWorkingTimerChange}
                    disabled={(state && (state.name === 'error' || state.name === 'working')) ? false : true}
                    className={`${isMobile ? '' : 'interval-slider'}`}
                />
                <div className='control-title'>Working State Timer</div>
                <hr className='line-break' />

                <Slider
                    aria-label="Always visible"
                    defaultValue={DEFAULT_FLASHING_INTERVAL}
                    step={0.25}
                    value={flashingInterval}
                    marks={SLIDER_MARKS}
                    valueLabelDisplay={isMobile ? 'off' : 'on'}
                    min={0.25}
                    max={5}
                    onChange={handleErrorTimerChange}
                    color="secondary"
                    disabled={(state && (state.name === 'error' || state.name === 'working')) ? false : true}
                    className={`${isMobile ? '' : 'interval-slider'}`}
                />
                <div className='control-title'>Error State Timer</div>
                <hr className='line-break' />
                
                <Button
                    className='control-panel-button'
                    variant="contained"
                    color='success'
                    onClick={handleResetTimers}
                    disabled={(
                        state && (
                        workingInterval !== DEFAULT_WOKRING_INTERVAL ||
                        flashingInterval !== DEFAULT_FLASHING_INTERVAL
                        ) && (state.name === 'error' ||
                        state.name === 'working'
                        )) ? false : true}
                >
                    Reset Timers
                </Button>

                <Button
                    className='control-panel-button'
                    variant="contained"
                    onClick={handleSystemRestart}
                    disabled={(state && state.name === 'error') ? false : true}
                >
                    System Restart
                </Button>

                <Button
                    className='control-panel-button'
                    variant="contained"
                    color='error'
                    style={{
                        margin: '5px 0'
                    }}
                    onClick={handleSystemError}
                    disabled={(state && state.name === 'working') ? false : true}
                >
                    System Error
                </Button>
            </div>

            {!isMobile && (<div className='control-section-title'>
                Control Panel
            </div>)}
        </>
    )
} 