import Layout from './../styles/itemLayout'
import styles from './../styles/item.module.css'
import { Modal } from './index'


const Item = props => {

    return (
        <>
            <li className={styles.item}>
            <a className={styles.a} href={props.url} target="_blank" rel="noopener noreferrer">
                <div className={styles.imgContainer}>
                <img className={styles.img} src={props.src} alt={props.name + 'のサムネイル'} />
                <span className={styles.price}>{props.price}</span>
                </div>
                <span className={styles.name}>{props.name} </span>
                </a>
            {/* <button type="button" className={styles.showImagesIcon} onClick={showImages()}> */}
            {/* <button type="button" className={styles.showImagesButton} onClick={showItemImages}> */}

            <Modal itemBaseUrl={props.url}/>
            </li>
        </>
  )
}
export default Item

// const Item = props => {
//   return (
//     <a className={styles.a} href={props.url}>
//       <img className={styles.img} src={props.src} alt={props.name + 'のサムネイル'}>
//         <span className={styles.price}>{props.price}</span>
//         <span className={styles.name}>{props.name} </span>
//       </img>
//     </a>
//   )
// }
// export default Item
