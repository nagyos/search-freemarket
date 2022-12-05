import { SearchInput, Title } from './index'
import styles from './../styles/navigation.module.css'

import type { Dispatch, FC, RefObject, SetStateAction } from 'react'
import { AllItems,Item } from '../entity/item'

type Props = {
  headerRef: RefObject<HTMLElement>
  setAllItems: Dispatch<SetStateAction<AllItems>>
}

const NavigatorTop: FC<Props> = ({ headerRef, setAllItems }) => {
  return (
    <header id="header" className={styles.navigation} ref={headerRef}>
      <Title title={'まとめて検索'} />
      <SearchInput setAllItems={setAllItems} />
    </header>
  )
}
export default NavigatorTop
