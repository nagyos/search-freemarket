import { TabList, ItemList } from './index'

const ItemContainer = props => {
  return (
    <div className="item-container">
      <div className="Tab-list">
        <TabList name={['すべて', 'メルカリ', 'ヤフオク', 'ラクマ', 'PayPayフリマ']} />
      </div>
      <div className="item-list">
        <ItemList className="item-grid" />
        <ItemList className="item-grid" />
        <ItemList className="item-grid" />
      </div>
    </div>
  )
}
export default ItemContainer
