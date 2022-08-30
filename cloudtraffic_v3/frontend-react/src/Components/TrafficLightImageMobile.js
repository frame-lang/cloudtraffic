import "./TrafficLightImageMobile.css";

export default function TrafficLightImageMobile({ color }) {
    return (
        <div className="tl-container">
            <div className="trafficLight">
                <span style={{
                    background: color.red
                }}></span>
                <span style={{
                    background: color.yellow
                }}></span>
                <span style={{
                    background: color.green
                }}></span>
            </div>
            <div className="stick"></div>
        </div>
    )
}