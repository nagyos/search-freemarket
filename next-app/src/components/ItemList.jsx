import { Item } from './index'

const ItemList = props => {
  return (
    <div className="item-grid">
      <Item
        src={
          'https://wiblok.com/wp-content/uploads/2022/01/png-transparent-go-software-framework-web-framework-application-programming-interface-programming-language-bee-doctor-honey-bee-food-computer-programming.png'
        }
        name={'蜂さん'}
        price={'2100'}
      />
    </div>
  )
}
export default ItemList
