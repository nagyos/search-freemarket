import { FC } from 'react'
import styles from './../styles/item.module.css'
import { Modal } from './index'
import { Item } from '../entity/item'

const Item: FC<Item> = props => {
  return (
    <>
      <li className={styles.item}>
        <a className={styles.a} href={props.url} target="_blank" rel="noopener noreferrer">
          <div className={styles.imgContainer}>
            <img className={styles.img} src={props.image} alt={props.name + 'のサムネイル'} />
            <span className={styles.price}>{props.price}</span>
          </div>
          <span className={styles.name}>{props.name} </span>
        </a>
        <Modal itemUrl={props.url} />
      </li>
    </>
  )
}
export default Item
