import { AllItems } from './item'

export const parentKey = 'searchResultInfo'

export const getSeachResultInfoOnLocalStorage = (key: string) => {
    const jsonValue = window.localStorage.getItem(parentKey)
    if (jsonValue !== null) return JSON.parse(jsonValue)[key]

    return null
}

export interface searchResultInfo {
  items: AllItems
  keyword: string
  minPrice: string
  maxPrice: string
  sortOrder: string
  status: string
}
