import { useEffect, useState } from 'react';
import './Main.css'
import { Button, Slider } from '@mui/material';

import TrafficLightImage from './TrafficLightImage';
import UmlImage from './UmlImage';
import {
  BEGIN_STATE,
  WORKING_STATE,
  STATES,
  SYSTEM_ERROR_STATE,
  END_STATE,
  SLIDER_MARKS,
  DEFAULT_WOKRING_INTERVAL,
  DEFAULT_FLASHING_INTERVAL
} from '../Utils/Constants';

export default function Layout({state, sendMessage}) {
  const [activeState, setActiveState] = useState(BEGIN_STATE);
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

  useEffect(() => {
    if (state.name === STATES['BEGIN_STATE']) {
      setActiveState(BEGIN_STATE);
    } else if (state.name === STATES['WORKING_STATE']) {
      setActiveState(WORKING_STATE[state.color]);
    } else if (state.name === STATES['ERROR_STATE']) {
      setActiveState(SYSTEM_ERROR_STATE[state.color]);
    } else if (state.name === STATES['END_STATE']) {
      setActiveState(END_STATE);
    }
  }, [state.name, state.color])

    return (
      <div className='container'>
        <div className='img-section'>
          <TrafficLightImage color={activeState['color']} />
        </div>
        <div className='uml-section'>
          <UmlImage img={activeState['umlImgName']} />
        </div>
        <div className='control-section'>
          <div className='control-section-items'>
            <Slider
              aria-label="Always visible"
              defaultValue={DEFAULT_WOKRING_INTERVAL}
              step={0.25}
              marks={SLIDER_MARKS}
              value={workingInterval}
              valueLabelDisplay="on"
              min={0.25}
              max={5}
              onChange={handleWorkingTimerChange}
              disabled={(state && (state.name === 'error' || state.name === 'working')) ? false : true}
              style={{'marginTop': '20px'}}
            />
            <div className='control-title'>Working State Timer</div>
            <hr className='line-break' />

            <Slider
              aria-label="Always visible"
              defaultValue={DEFAULT_FLASHING_INTERVAL}
              step={0.25}
              value={flashingInterval}
              marks={SLIDER_MARKS}
              valueLabelDisplay="on"
              min={0.25}
              max={5}
              onChange={handleErrorTimerChange}
              color="secondary"
              disabled={(state && (state.name === 'error' || state.name === 'working')) ? false : true}
              style={{'marginTop': '45px'}}
            />
            <div className='control-title'>Error State Timer</div>
            <hr className='line-break' />
            
            <Button
              className='control-button'
              variant="contained"
              color='success'
              style={{
                margin: '30px 0 5px 0'
              }}
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
              className='control-button'
              variant="contained"
              style={{
                margin: '5px 0'
              }}
              onClick={handleSystemRestart}
              disabled={(state && state.name === 'error') ? false : true}
            >
              System Restart
            </Button>

            <Button
              className='control-button'
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

          <div className='control-section-title'>
            Control Panel
          </div>
        </div>
        {/* <div className='image-container'>
          <TrafficLightImage color={activeState['color']} />
        </div>

        <div className='code-container'>
          <div className='uml-container'>
            <UmlImage img={activeState['umlImgName']} />
          </div>

          <div className='action-container'>
            {(state.name === 'working' || state.name === 'error') && (
              <>
                <Button variant="outlined" onClick={handleSystemRestart}>System Restart</Button>
                <Button variant="outlined" onClick={handleSystemError}>System Error</Button>

                <Slider
                  aria-label="Always visible"
                  defaultValue={3}
                  // getAriaValueText={valuetext}
                  step={0.2}
                  marks={marks}
                  valueLabelDisplay="on"
                  min={0.2}
                  max={3}
                  onChange={handleChange}
                />
              </>
            )}
          </div>
        </div> */}
      </div>
    )
}