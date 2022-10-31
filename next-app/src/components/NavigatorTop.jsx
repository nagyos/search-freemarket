import { Title, SearchInput } from './index'
import styles from './../styles/navigation.module.css'


const NavigatorTop = props => {
  return (
    // <div className="navigation-top">
    <div className={styles.navigation}>
      <Title title={'まとめて検索'} />
      <SearchInput />
    </div>
  )
}
export default NavigatorTop
