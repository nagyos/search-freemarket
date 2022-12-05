export interface AllItems {
  mercari: Item[]
  paypayfleamarket: Item[]
  rakutenrakuma: Item[]
  yahooauction: Item[]
}

export interface Item {
  url: string
  name: string
  price: string
  image: string
}
