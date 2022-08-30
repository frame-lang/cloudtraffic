import { useEffect, useState } from 'react';
import { isMobile } from 'react-device-detect';
import './Main.css';

import TrafficLightImage from './TrafficLightImage';
import TrafficLightImageMobile from './TrafficLightImageMobile';
import UmlImage from './UmlImage';
import ControlPanel from './ControlPanel';
import {
  BEGIN_STATE,
  WORKING_STATE,
  STATES,
  SYSTEM_ERROR_STATE,
  END_STATE
} from '../Utils/Constants';

import Typography from '@mui/material/Typography';
import SwipeableDrawer from '@mui/material/SwipeableDrawer';
import Box from '@mui/material/Box';
import { styled } from '@mui/material/styles';
import { grey } from '@mui/material/colors';
import { Global } from '@emotion/react';

const drawerBleeding = 56;

const Puller = styled(Box)(({ theme }) => ({
  width: 30,
  height: 6,
  backgroundColor: theme.palette.mode === 'light' ? grey[300] : grey[900],
  borderRadius: 3,
  position: 'absolute',
  top: 8,
  left: 'calc(50% - 15px)',
}));

const StyledBox = styled(Box)(({ theme }) => ({
  backgroundColor: theme.palette.mode === 'light' ? '#fff' : grey[800],
}));

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

  const [open, setOpen] = useState(false);

  const toggleDrawer = (newOpen: boolean) => () => {
    setOpen(newOpen);
  };

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
            flexDirection: 'column',
            justifyContent:'center',
            alignItems: 'center',
            padding: '10px'
          }}>
            <TrafficLightImageMobile color={activeState['color']} />
            <hr />
            <UmlImage img={activeState['umlImgName']} />
            <Global
              styles={{
                '.MuiDrawer-root > .MuiPaper-root': {
                  height: `calc(50% + ${95}px)`,
                  overflow: 'visible',
                },
              }}
            />
          </div>

          <SwipeableDrawer
            anchor="bottom"
            open={open}
            onClose={toggleDrawer(false)}
            onOpen={toggleDrawer(true)}
            swipeAreaWidth={drawerBleeding}
            disableSwipeToOpen={false}
            ModalProps={{
              keepMounted: true,
            }}
          >
            <StyledBox
              sx={{
                position: 'absolute',
                top: -drawerBleeding,
                borderTopLeftRadius: 8,
                borderTopRightRadius: 8,
                visibility: 'visible',
                right: 0,
                left: 0,
                boxShadow: "-3px -4px 9px 0px rgba(194,194,194,1)"
              }}
            >
              <Puller />
              <Typography sx={{ p: 2, color: 'text.primary' }}>
                Control Panel
              </Typography>
            </StyledBox>
            <StyledBox
              sx={{
                px: 2,
                pb: 2,
                height: '100%',
                overflow: 'auto',
              }}
            >
                <ControlPanel
              state={state}
              sendMessage={sendMessage}
              connectionStatus={connectionStatus}
            />
            </StyledBox>
          </SwipeableDrawer>
        </>

      )}
    </>
  )
}