import { AllItems } from 'entity/item'
import { searchResultInfo } from 'entity/localStorage'
import { handleErrors } from './error'

const baseEndPointWithColly = 'http://localhost:8000/v1/search-colly?'
const baseEndPointWithAgouti = 'http://localhost:8000/v1/search-aagouti?'

var crawledItems: AllItems = {} as AllItems

export const makeCrawlingItemsUrl = async (searchInfo: searchResultInfo): Promise<string> => {
  var params: string[] = []
  params.push(
    'keyword=' + searchInfo.keyword,
    'price-min=' + searchInfo.minPrice,
    'price-max=' + searchInfo.maxPrice,
    'sort=' + searchInfo.sortOrder.split('-')[0],
    'order=' + searchInfo.sortOrder.split('-')[1],
    'status=' + searchInfo.status
  )

  return params.join('&')
}
export const makeCrawlingItemsParam = (searchInfo: searchResultInfo): string => {
  var params: string[] = []
  params.push(
    'keyword=' + searchInfo.keyword,
    'price-min=' + searchInfo.minPrice,
    'price-max=' + searchInfo.maxPrice,
    'sort=' + searchInfo.sortOrder.split('-')[0],
    'order=' + searchInfo.sortOrder.split('-')[1],
    'status=' + searchInfo.status
  )

  return params.join('&')
}

export const getQueryParam = (searchInfo: searchResultInfo): Object => {
    return {
      keyword: searchInfo.keyword,
      'price-min': searchInfo.minPrice,
      'price-max': searchInfo.maxPrice,
      sort: searchInfo.sortOrder.split('-')[0],
      order: searchInfo.sortOrder.split('-')[1],
      status: searchInfo.status
    }
}
const fetchCrawlingAllItems = async (url: string): Promise<AllItems> => {
  const response = await fetch(url)
    .catch(e => {
      throw Error(e)
    })
    .then(handleErrors)
    .then(res => res.json())
  return await response
}

export const setCrawledItems = async (items: AllItems) => {
  for (const siteName in items)
    Object.keys(items).map((key: string, _: number) => {
      crawledItems[key as keyof AllItems] = items[key as keyof AllItems]
    })
}
export const getCrawledItems = async () => {
  return crawledItems
}

export const crawleItemsWithColly = async (searchInfo: searchResultInfo): Promise<AllItems> => {
  const url: string = baseEndPointWithColly + (await makeCrawlingItemsUrl(searchInfo))
  return await fetchCrawlingAllItems(url)
}

export const crawleItemsWithAgouti = async (searchInfo: searchResultInfo): Promise<AllItems> => {
  const url: string = baseEndPointWithAgouti + (await makeCrawlingItemsUrl(searchInfo))
  return await fetchCrawlingAllItems(url)
}
