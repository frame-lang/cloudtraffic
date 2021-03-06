import { AppBar, Box, Toolbar, Typography, Button } from '@mui/material';
import { LoadingButton } from '@mui/lab';
import TrafficIcon from '@mui/icons-material/Traffic';
import AutorenewIcon from '@mui/icons-material/Autorenew';
import { STATES } from '../Utils/Constants';
import { isMobile } from 'react-device-detect';

export default function Header({
    state,
    sendMessage,
    connectionStatus,
    reconnect
}) {
    function handleClick() {
        if (state.name === STATES['END_STATE']) {
        sendMessage(JSON.stringify({event: 'createWorkflow'}))
            return
        }
        sendMessage(JSON.stringify({event: 'end'}))
    }

    return (
        <Box sx={{ flexGrow: 1 }}>
            <AppBar position="static">
            <Toolbar>
                <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
                    Traffic Light App V2
                </Typography>

                <Box sx={{ '& > button': { m: 1 } }}>
                    {connectionStatus !== 'Closed' ? (
                        <LoadingButton
                            color="secondary"
                            onClick={handleClick}
                            loading={state.loading}
                            loadingPosition="start"
                            startIcon={<TrafficIcon />}
                            variant="contained"
                        >
                            {state.loading && 'Loading...'}
                            {(!state.loading && state.name === STATES['END_STATE']) && `Create ${isMobile ? '' : 'Workflow'}` }
                            {(!state.loading && state.name !== STATES['END_STATE']) && `Destroy ${isMobile ? '' : 'Workflow'}` }
                        </LoadingButton>
                    ) : (
                        <Button onClick={reconnect} variant="contained" color='success' startIcon={<AutorenewIcon />}>
                            Reconnect
                        </Button>
                    )}
                </Box>
            </Toolbar>
            </AppBar>
        </Box>
    );
}
