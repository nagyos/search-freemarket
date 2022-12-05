import type { NextPage } from 'next'
import { useState, createContext, useEffect } from 'react'
import Head from 'next/head'
import { NavigatorTop, ItemContainer } from './../components/index'
import { useLayout } from './../hooks/useLayout'
import { AllItems } from '../entity/item'
import { getSeachResultInfoOnLocalStorage } from '../entity/localStorage'
import { useRouter } from 'next/router'

export const AllItemsContext = createContext<AllItems>({} as AllItems)
const Home: NextPage = () => {
  const [allItems, setAllItems] = useState<AllItems>({} as AllItems)
  const { headerHeight, headerRef, itemContainerHeight } = useLayout()
  const router = useRouter()

  useEffect(() => {
    if (router.pathname === '/') return
    setAllItems(getSeachResultInfoOnLocalStorage('items'))
  }, [])

  return (
    <div className="App" style={{ paddingTop: headerHeight }}>
      <Head>
        <title>まとめて検索</title>
        <link rel="icon" href="/icon.ico" />
      </Head>
      <NavigatorTop headerRef={headerRef} setAllItems={setAllItems} />
      <AllItemsContext.Provider value={allItems}>
        <ItemContainer headerHeight={headerHeight} itemContainerHeight={itemContainerHeight} />
      </AllItemsContext.Provider>
    </div>
  )
}

export default Home
