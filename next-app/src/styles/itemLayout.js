import styles from './item.module.css'

export default function Layout({ children }) {
  return <div className={styles.itemGrid}>{children}</div>
}