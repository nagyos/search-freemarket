import { useEffect, useState } from 'react'

export const showItemImages = () => {
  const [image, setImage] = useState('')

    const showItemImages = async (url:string) => {
        const images: string[] = await fetchItemImages(url)
        setImage(images[0])
    }

    const fetchItemImages = async (url: string): Promise<string[]> => {
        const response = await fetch('http://localhost:8000/getImages')
        const result = await response.json()
        return result
    }

    const changeShowImage = () => {
        setImage("faf")
    }

    return (
        <div className='modal-window'>
            <div className='main-image'>
                <img src={image} />
            </div>
            <div className='sub-images'>
                <ul>
                    <li onClick={changeShowImage}></li>
                </ul>
            </div>
        </div>
    )

}