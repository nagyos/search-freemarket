import type { NextPage } from 'next'
import { useState } from 'react'
import Head from 'next/head'
import Image from 'next/image'
import { NavigatorTop, ItemContainer, Title, TabList, Item } from './../components/index'
import styles from './../styles/navigation.module.css'
import itemStyles from './../styles/item.module.css'

import { getWindowSize } from './../hooks/GetWindowSize'

//TODO:スクロールバーの実装
/*
タブの画面切り替え．SPA
その他の複数の画像出力→メルカリ以外番号で紐付いていない→クリックしたら商品URLに飛び．取得してくることになりそう．
Typescript化
モジュール分け
css
バックエンド側で並列処理
複数のJson受取
terraformを使ったインフラ構築
公開
*/
//TODO:Shadow-rootをつけるかどうか
interface Item {
  Url: string
  Name: string
  Price: string
  Image: string
}

const siteNames = ['yahooAuction', 'mercari', 'rakutenRakuma', 'paypayFleaMarket']
const Home: NextPage = () => {
  const [allItemsInfo, setAllItemsInfo] = useState({})
  const [items, setItems] = useState<Array<Item>>([])
  const [input, setInput] = useState('')

  const fetchCrawlingItems = async (): Promise<Item[]> => {
    const response = await fetch('http://localhost:8000')
    const result = await response.json()
    return result
  }

  const handleClick = async () => {
    const allItems: Item[] = await fetchCrawlingItems()
    setItems(allItems)
  }

  const { height, width } = getWindowSize()

  return (
    <div>
      {/* <div className="navigation-top"> */}
      <div className={styles.navigation}>
        <Title title={'まとめて検索'} />
        <div className={styles.form}>
          {/* <form> */}
          <input
            className={styles.input}
            type="text"
            placeholder="何をお探しですか？"
            onChange={e => setInput(e.target.value)}
          />
          <button type="submit" onClick={handleClick} disabled={!input}>
            <svg viewBox="0 0 24 24" className="search-icon" width={20} height={20}>
              <path
                className="cls-1"
                d="M20.79,19.8l-3.94-3.94a7.86,7.86,0,1,0-1,1l3.94,3.94a.7.7,0,0,0,.5.21.67.67,0,0,0,.49-.21A.69.69,0,0,0,20.79,19.8ZM6.28,15.39a6.44,6.44,0,1,1,9.11,0,6.43,6.43,0,0,1-9.11,0Z"
              ></path>
            </svg>
          </button>
          {/* </form> */}
        </div>
      </div>
      <div className="item-container">
        <div className="Tab-list">
          <TabList name={['すべて', 'メルカリ', 'ヤフオク', 'ラクマ', 'PayPayフリマ']} />
        </div>
        <div className={itemStyles.itemList}>
          //バーを持つmodule.cssを作ってここに貼る
          {/* <div className="item-grid"> */}
          <div className={itemStyles.container}>
            {items.map((data, _) => {
              return (
                <Item
                  url={'https://paypayfleamarket.yahoo.co.jp' + data.Url}
                  src={data.Image}
                  name={data.Name}
                  price={data.Price}
                />
              )
            })}
          </div>
          <div className={itemStyles.container}>
            {items.map((data, _) => {
              return (
                <Item
                  url={'https://paypayfleamarket.yahoo.co.jp' + data.Url}
                  src={data.Image}
                  name={data.Name}
                  price={data.Price}
                />
              )
            })}
          </div>
        </div>
      </div>
    </div>
  )
}

export default Home

// const Home: NextPage = () => {
//   return (
//     <div>
//       <NavigatorTop />
//       <ItemContainer />
//     </div>
//   )
// }
