import { getSeachResultInfoOnLocalStorage, searchResultInfo } from 'entity/localStorage'
import { FC, useState, SetStateAction, Dispatch, useEffect } from 'react'
import { useRouter } from 'next/router'
type Props = {
  setSortOrder: Dispatch<SetStateAction<searchResultInfo>>
}

const CATEGORY = [
  { label: '新しい順', value: 'created_time-desc' },
  { label: 'おすすめ順', value: 'score-desc' },
  { label: '価格の安い順', value: 'price-asc' },
  { label: '価格の高い順', value: 'price-desc' },
  { label: 'いいね順', value: 'like_count-desc' }
]

const SearchSortOrder: FC<Props> = ({ setSortOrder }) => {
  const [selected, setSelected] = useState<string>(CATEGORY[0].value)
  const router = useRouter()
  useEffect(() => {
    if (router.pathname === '/') return
    const sortOrder = getSeachResultInfoOnLocalStorage('sortOrder')
    setSelected(sortOrder)
    setSortOrder(prevState => ({ ...prevState, sortOrder: sortOrder }))
  }, [])
  return (
    <select
      value={selected}
      name="status"
      onChange={e => {
        setSelected(e.target.value)
        setSortOrder(prevState => ({ ...prevState, sortOrder: e.target.value }))
      }}
    >
      {CATEGORY.map((v, i) => (
        <option key={i} value={v.value}>
          {v.label}
        </option>
      ))}
    </select>
  )
}
export default SearchSortOrder
