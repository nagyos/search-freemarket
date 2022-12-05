import { Dispatch, FC, SetStateAction, useContext, useEffect, useState } from 'react'
import styles from './../styles/navigation.module.css'
import { SearchKeyword, SearchPrice, SearchSortOrder } from './../components/index'
import { AllItems } from '../entity/item'
import { parentKey, searchResultInfo } from '../entity/localStorage'
import {
  crawleItemsWithColly,
  crawleItemsWithAgouti,
  getCrawledItems,
  setCrawledItems,
  getQueryParam,
  makeCrawlingItemsUrl,
  makeCrawlingItemsParam
} from '../handler/fetchCrawlingItems'
import Link from 'next/link'

type Props = {
  setAllItems: Dispatch<SetStateAction<AllItems>>
}

export const SearchInput: FC<Props> = ({ setAllItems }) => {
  const [searchInfo, setSearchInfo] = useState<searchResultInfo>({
    sortOrder: 'created_time-desc',
    status: 'on_sale'
  } as searchResultInfo)

  const storeSearchInfo = async (info: AllItems) => {
    console.log(info)
    const items = await JSON.parse(JSON.stringify(info))
    setSearchInfo(prevState => ({
      ...prevState,
      items: items
    }))
    var tempSearchInfo = await JSON.parse(JSON.stringify(searchInfo))
    tempSearchInfo.items = items
    window.localStorage.setItem(parentKey, JSON.stringify(tempSearchInfo))
  }

  const handleClick = async () => {
    setCrawledItems(await crawleItemsWithColly(searchInfo))
    setAllItems(await getCrawledItems())
    setCrawledItems(await crawleItemsWithAgouti(searchInfo))
    setAllItems(await JSON.parse(JSON.stringify(await getCrawledItems())))

    await storeSearchInfo(await getCrawledItems())
  }

  return (
    <div className={styles.form}>
      {/* <form className={styles.form}> */}
      <SearchKeyword
        placeHolder="何かお探しですか？"
        setKeyword={setSearchInfo}
        crawleItems={handleClick}
      />
      {/* <Link href="/search" as={'search?' + makeCrawlingItemsParam(searchInfo)}> */}
      <button type="submit" onClick={handleClick} disabled={!searchInfo.keyword}>
        <svg viewBox="0 0 24 24" className="search-icon" width={20} height={20}>
          <path
            className="cls-1"
            d="M20.79,19.8l-3.94-3.94a7.86,7.86,0,1,0-1,1l3.94,3.94a.7.7,0,0,0,.5.21.67.67,0,0,0,.49-.21A.69.69,0,0,0,20.79,19.8ZM6.28,15.39a6.44,6.44,0,1,1,9.11,0,6.43,6.43,0,0,1-9.11,0Z"
          ></path>
        </svg>
      </button>
      {/* </Link> */}
      <SearchPrice setPrice={setSearchInfo} />
      <SearchSortOrder setSortOrder={setSearchInfo} />
      {/* </form> */}
    </div>
  )
}
export default SearchInput
