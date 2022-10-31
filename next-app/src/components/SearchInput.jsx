import { useRouter } from 'next/router'
import { useState } from 'react'

const SearchInput = props => {
  const router = useRouter()
  const [input, setInput] = useState('')

  const clickButton = () => {
    if (!input) {
      return
    }

    // router.push({
    // //   pathname: 'localhost:8080/', //URL
    //   pathname: '/search', //URL
    //   query: { keyword: input } //検索クエリ
    // })

      fetch("http://localhost:8080/")
      .then(response => response.json())
      .then(data => {
            console.log(data)
            setInput(data)
        })

  }

  return (
    <div className="search-input">
      <form>
        <input
          type="text"
          //   value={text}
          placeholder="何をお探しですか？"
          onChange={e => setInput(e.target.value)}
        />
        <div className="search-button">
          <button type="submit" onClick={clickButton} disabled={!input}>
            <svg viewBox="0 0 24 24" class="search-icon" width={20} height={20}>
              <path
                className="cls-1"
                d="M20.79,19.8l-3.94-3.94a7.86,7.86,0,1,0-1,1l3.94,3.94a.7.7,0,0,0,.5.21.67.67,0,0,0,.49-.21A.69.69,0,0,0,20.79,19.8ZM6.28,15.39a6.44,6.44,0,1,1,9.11,0,6.43,6.43,0,0,1-9.11,0Z"
              ></path>
            </svg>
          </button>
        </div>
      </form>
    </div>
  )
}
export default SearchInput
