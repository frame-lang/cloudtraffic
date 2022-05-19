import { AppBar, Box, Toolbar, Typography } from '@mui/material';
import { LoadingButton } from '@mui/lab';
import TrafficIcon from '@mui/icons-material/Traffic';

import { STATES } from '../Utils/Constants';

export default function Header({state, sendMessage}) {

    function handleClick() {
        if (state.name === STATES['END_STATE']) {
            sendMessage('init')
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
                </Box>
            </Toolbar>
            </AppBar>
        </Box>
    );
}
