import { useEffect, useState } from 'react';

import Button from '@mui/material/Button';

import TrafficLightImage from './TrafficLightImage';
import UmlImage from './UmlImage';
import { BEGIN_STATE, WORKING_STATE, STATES, SYSTEM_ERROR_STATE, END_STATE } from '../Utils/Constants';

export default function Layout({state, sendMessage}) {
  const [activeState, setActiveState] = useState(BEGIN_STATE);

  const handleSystemError = () => {
    sendMessage('error')
  }

  const handleSystemRestart = () => {
    sendMessage('restart')
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
      <main>
        <div className='image-container'>
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
              </>
            )}
          </div>
        </div>
      </main>
    )
}