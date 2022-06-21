import begin from '../Assets/Images/begin.svg'
import red from '../Assets/Images/red.svg'
import yellow from '../Assets/Images/yellow.svg'
import green from '../Assets/Images/green.svg'
import flashingRed from '../Assets/Images/flashing_red.svg'
import end from '../Assets/Images/end.svg'

const mapping = {
    begin,
    red,
    yellow,
    green,
    flashingRed,
    end
};

export default function ({ img }) {
    return (
        <img className='umg-img' src={mapping[img]}/>
    )
}