import { AppBar, Box, Toolbar, Typography, Button } from '@mui/material';
import { LoadingButton } from '@mui/lab';
import TrafficIcon from '@mui/icons-material/Traffic';
import AutorenewIcon from '@mui/icons-material/Autorenew';
import { STATES } from '../Utils/Constants';

export default function Header({
    state,
    sendMessage,
    connectionStatus,
    reconnect
}) {
    function handleClick() {
        if (state.name === STATES['END_STATE']) {
            sendMessage('start')
            return
        }
        sendMessage('end')
    }

    return (
        <Box sx={{ flexGrow: 1 }}>
            <AppBar position="static">
            <Toolbar>
                <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
                    Traffic Light App
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
                            {(!state.loading && state.name === STATES['END_STATE']) && 'Create Traffic Light' }
                            {(!state.loading && state.name !== STATES['END_STATE']) && 'Destroy Traffic Light' }
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
