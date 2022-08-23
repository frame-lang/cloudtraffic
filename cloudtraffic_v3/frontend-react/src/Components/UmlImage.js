import begin from '../Assets/Images/begin.svg'
import red from '../Assets/Images/red.svg'
import yellow from '../Assets/Images/yellow.svg'
import green from '../Assets/Images/green.svg'
import flashingRed from '../Assets/Images/flashing_red.svg'
import end from '../Assets/Images/end.svg'

import { isMobile } from 'react-device-detect';
import './UmlImage.css'

const mapping = {
    begin,
    red,
    yellow,
    green,
    flashingRed,
    end
};

export default function UmlImage  ({ img }) {
    return (
        <img className={`${isMobile ? 'uml-img-mobile' : 'uml-img'}`} alt='umlImage' src={mapping[img]}/>
    )
}