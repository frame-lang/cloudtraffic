import { useState, useEffect } from 'react';

import useWebSocket, { ReadyState } from 'react-use-websocket';

import '../App.css';
import Header from './Header'
import Main from './Main'

const socketUrl = `ws://${window.location.hostname}:8000/ws`;
// const socketUrl = `ws://warm-bayou-79838.herokuapp.com/ws`;


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
    // const data = JSON.parse(JSON.parse(lastMessage.data));
    console.log('Data received: ', JSON.parse(lastMessage.data));
    // setState(data)
  }, [lastMessage]);

  useEffect(() => {
    sendMessage('start')
  }, [])

  return (
    <div className="App">
      <Header state={state} sendMessage={sendMessage}/>
      <Main state={state} sendMessage={sendMessage}/>
    </div>
  );
}

export default App;
