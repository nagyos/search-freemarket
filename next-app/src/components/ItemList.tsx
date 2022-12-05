import { FC, useContext } from 'react'
import { AllItemsContext } from 'pages'
import { Item as ItemCmponent } from './index'
import itemStyles from './../styles/item.module.css'
import { Item } from '../entity/item'

type Props = {
  items: Item[]
}

const ItemList: FC<Props> = ({ items }) => {
    const makeItemList = () => {
    const ItemsList: JSX.Element[] = []
    items.map((item: Item, _: number) => {
      ItemsList.push(
        <ItemCmponent
          url={item.url}
          image={item.image}
          name={item.name}
          price={item.price}
          key={_}
        />
      )
    })
    console.log('===itemList===')
    return ItemsList
  }

  return <div className={itemStyles.itemList}>{makeItemList()}</div>
}
export default ItemList
