import { useEffect, useState } from 'react';
import { isMobile } from 'react-device-detect';
import './Main.css'


import TrafficLightImage from './TrafficLightImage';
import UmlImage from './UmlImage';
import ControlPanel from './ControlPanel';
import {
  BEGIN_STATE,
  WORKING_STATE,
  STATES,
  SYSTEM_ERROR_STATE,
  END_STATE
} from '../Utils/Constants';

export default function Layout({state, sendMessage, connectionStatus}) {
  const [activeState, setActiveState] = useState(BEGIN_STATE);

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
      <>
        {/* Desktop View */}
        {!isMobile && (
          <div className='container'>
            <div className='img-section'>
              <TrafficLightImage color={activeState['color']} />
            </div>
            <div className='uml-section'>
              <UmlImage img={activeState['umlImgName']} />
            </div>
            <div className={`${isMobile ? '' : 'control-section-full'} control-section`}>
              <ControlPanel
                state={state}
                sendMessage={sendMessage}
                connectionStatus={connectionStatus}
              />
            </div>
          </div>
        )}

        {/* Mobile View */}
        {isMobile && (
          <>
            <div style={{
              display: 'flex',
              justifyContent:'center',
              padding: '10px'
            }}>
              <TrafficLightImage color={activeState['color']} />
              <UmlImage img={activeState['umlImgName']} />
            </div>
            <hr style={{
              background: '#E4E3E3',
              height: '2px',
              border: 'none'
            }} />
            <div style={{
              margin: '10px',
              padding: '10px'
            }}>
              <ControlPanel
                  state={state}
                  sendMessage={sendMessage}
                  connectionStatus={connectionStatus}
                />
            </div>
          </>

        )}
      </>
    )
}