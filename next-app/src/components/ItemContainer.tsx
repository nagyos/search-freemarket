import { FC, useContext } from 'react'
import { ItemList } from './index'
import itemStyles from './../styles/item.module.css'
import { AllItemsContext } from 'pages'
import { AllItems } from '../entity/item'

type Props = {
  headerHeight: number
  itemContainerHeight: number
}

const ItemContainer: FC<Props> = ({ headerHeight, itemContainerHeight }) => {
  const allItems = useContext<AllItems>(AllItemsContext)

  const isEmpty = (obj: AllItems) => {
    return (
      obj !== null &&
      typeof obj === 'object' &&
      obj.constructor === Object &&
      !Object.keys(obj).length
    )
  }
  const siteNameToJapanese = new Map<string, string>([
    ['mercari', 'メルカリ'],
    ['paypayfleamarket', 'PayPayフリマ'],
    ['rakutenrakuma', '楽天ラクマ'],
    ['yahooauction', 'ヤフオク!']
  ])

  const makeItemContainers = () => {
    const Containers: JSX.Element[] = []
    for (const siteName in allItems) {
      Containers.push(
        <div>
          <h4 style={{ textAlign: 'center', margin: '0' }}>{siteNameToJapanese.get(siteName)}</h4>
          <ItemList key={siteName} items={allItems[siteName as keyof AllItems]} />
        </div>
      )
    }
    // console.log(allItems, '=====container====')
    return (
      <div
        className={itemStyles.container}
        style={{
          position: 'fixed',
          top: headerHeight,
          height: itemContainerHeight - 24.5,
          marginBottom: '0.33em'
        }}
      >
        {Containers}
      </div>
    )
  }
  return (
    <div
      className="item-container"
      style={{
        padding: '0.33em 1px'
      }}
    >
      {!isEmpty(allItems) && makeItemContainers()}
    </div>
  )
}
export default ItemContainer
