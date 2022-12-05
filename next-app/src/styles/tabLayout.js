import styles from './tab.module.css'

export default function Layout({ children }) {
    return <li className={styles.li}>{children}</li>
}
