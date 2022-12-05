import { handleErrors } from './error'

const baseEndPointWithColly = 'http://localhost:8000/v1/getItemImages'

export const fetchCrawlingItemImages = async (itemUrl:string): Promise<string[]> => {
  const response = await fetch(baseEndPointWithColly + '/' + itemUrl)
    .catch(e => {
      throw Error(e)
    })
    .then(handleErrors)
    .then(res => res.json())
  return await response
}
