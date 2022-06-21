const Footer = ({connectionStatus}) => {
    return (
        <div style={{
            height: '25px',
            backgroundColor: '#E7E7E7',
            display: 'flex',
            alignItems: 'center',
            paddingLeft: '20px'
        }}>
            <span style={{
                color: '#737272'
            }}>Connection Status:</span>
            <span style={{
                color: '#494949',
                marginLeft: '6px',
                fontWeight: '500'

            }}>{connectionStatus}</span>
        </div>
    )
}

export default Footer;