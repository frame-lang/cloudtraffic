import { useState, useEffect } from 'react';

import useWebSocket, { ReadyState } from 'react-use-websocket';

import '../App.css';
import Header from './Header'
import Main from './Main'

const socketUrl = `ws://${window.location.hostname}:8000/ws`;

function App() {

  const {
    sendMessage,
    lastMessage,
    readyState,
  } = useWebSocket(socketUrl);

  const connectionStatus = {
    [ReadyState.CONNECTING]: 'Connecting',
    [ReadyState.OPEN]: 'Open',
    [ReadyState.CLOSING]: 'Closing',
    [ReadyState.CLOSED]: 'Closed',
    [ReadyState.UNINSTANTIATED]: 'Uninstantiated',
  }[readyState];

  const [state, setState] = useState({});

  useEffect(() => {
    if (!lastMessage) return;
    const data = JSON.parse(lastMessage.data);
    setState(data)
    console.log('Data received:', data);
  }, [lastMessage]);

  useEffect(() => {
    sendMessage('init')
  }, [])

  return (
    <div className="App">
      <Header state={state} sendMessage={sendMessage}/>
      <Main state={state} sendMessage={sendMessage}/>
    </div>
  );
}

export default App;
