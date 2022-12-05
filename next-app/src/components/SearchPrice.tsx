import { getSeachResultInfoOnLocalStorage, searchResultInfo } from 'entity/localStorage'
import { FC, useState, Dispatch, SetStateAction, useEffect } from 'react'
import styles from './../styles/Price.module.css'
type Props = {
  setPrice: Dispatch<SetStateAction<searchResultInfo>>
}
import { useRouter } from 'next/router'
const SearchPrice: FC<Props> = ({ setPrice }) => {
  const [minPrice, setMinPrice] = useState<string>('')
  const [maxPrice, setMaxPrice] = useState<string>('')
  const router = useRouter()

  useEffect(() => {
    if (router.pathname === '/') return
    const minPrice = getSeachResultInfoOnLocalStorage('minPrice')
    const maxPrice = getSeachResultInfoOnLocalStorage('maxPrice')
    setMinPrice(minPrice)
    setPrice(prevState => ({ ...prevState, minPrice: minPrice }))
    setMaxPrice(maxPrice)
    setPrice(prevState => ({ ...prevState, maxPrice: maxPrice }))
  }, [])

  return (
    <form className={styles.container}>
      <div className={styles.price}>
        <input
          type="number"
          name="min-price"
          className={styles.input}
          value={minPrice}
          min="0"
          max="9999999"
          step="1000"
          placeholder="Min"
          onChange={e => {
            setMinPrice(e.target.value)
            setPrice(prevState => ({ ...prevState, minPrice: e.target.value }))
          }}
        />
        <input
          type="range"
          id="min-price-range"
          min="0"
          max="10000"
          step="100"
          onChange={e => {
            setMinPrice(e.target.value)
            setPrice(prevState => ({ ...prevState, minPrice: e.target.value }))
          }}
        />
      </div>
      -
      <div className={styles.price}>
        <input
          type="number"
          name="max-price"
          className={styles.input}
          value={maxPrice}
          min="0"
          max="9999999"
          step="1000"
          placeholder="Max"
          onChange={e => {
            setMaxPrice(e.target.value)
            setPrice(prevState => ({ ...prevState, maxPrice: e.target.value }))
          }}
        />
        <input
          type="range"
          id="max-price-range"
          min="0"
          max="10000"
          step="100"
          onChange={e => {
            setMaxPrice(e.target.value)
            setPrice(prevState => ({ ...prevState, maxPrice: e.target.value }))
          }}
        />
      </div>
    </form>
  )
}
export default SearchPrice
