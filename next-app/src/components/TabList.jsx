import { Tab } from './index'

const TabList = props => {

    var makeTabList = () => {
        const tabs = [];
        for (let i = 0; i <props.name.length; i++) {
            tabs.push(<Tab name={props.name[i]}/>)
        }
        return <ul className='tab-list'>{ tabs }</ul>;
    };

    return (
         makeTabList()

    // <ul className="tab-list">
    //    <Tab name={props.name[0]} />
    //    <Tab name={props.name[1]} />
    //    <Tab name={props.name[2]} />
    //    <Tab name={props.name[3]} />
    //    <Tab name={props.name[4]} />
    // </ul>
  )
}
export default TabList
