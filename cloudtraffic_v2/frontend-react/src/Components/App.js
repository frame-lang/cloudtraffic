import { useState, useEffect } from 'react';

import useWebSocket, { ReadyState } from 'react-use-websocket';

import '../App.css';
import Header from './Header';
import Main from './Main';
import Footer from './Footer';

const socketUrl = `ws://${window.location.hostname}:9000/ws`;

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

  const reconnect = () => {
    window.location.reload(false);
  }

  useEffect(() => {
    if (!lastMessage) return;
    const data = JSON.parse(lastMessage.data);
    setState(data)
    console.log('Data received:', data);
  }, [lastMessage]);

  useEffect(() => {
    sendMessage(JSON.stringify({event: 'createWorkflow'}))
  }, [])

  return (
    <div className="App">
      <Header state={state} sendMessage={sendMessage} connectionStatus={connectionStatus} reconnect={reconnect}/>
      <Main state={state} sendMessage={sendMessage}/>
      <Footer connectionStatus={connectionStatus} />
    </div>
  );
}

export default App;
