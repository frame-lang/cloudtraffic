const Footer = ({connectionStatus}) => {
    return (
        <div style={{
            height: '25px',
            backgroundColor: '#E7E7E7',
            display: 'flex',
            alignItems: 'center',
            paddingLeft: '20px'
        }}>
            Connection Status: <b> {connectionStatus} </b> 
        </div>
    )
}

export default Footer;