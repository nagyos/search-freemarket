import { useState } from 'react'
import itemstyles from './../styles/item.module.css'
import styles from './../styles/Modal.module.css'

// const Modal = props => {

type Props = {
  itemBaseUrl: string
}

const Modal: React.FC<Props> = props => {
  const [isOpen, setIsOpen] = useState(false)

  const toggleModal = () => {
    setIsOpen(!isOpen)
  }

  const [image, setImage] = useState('')

  const showItemImages = async (url: string) => {
    const images: string[] = await fetchItemImages(url)
    setImage(images[0])
  }

  const fetchItemImages = async (url: string): Promise<string[]> => {
    const response = await fetch('http://localhost:8000/getImages')
    const result = await response.json()
    return result
  }

  const changeShowImage = () => {
    setImage('faf')
  }

  return (
    <>
      <button
        type="button"
        className={itemstyles.showImagesButton}
        onClick={() => {
          showItemImages(props.itemBaseUrl)
          toggleModal()
        }}
      >
        <svg
          className={itemstyles.showImagesIcon}
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 576 512"
        >
          <path d="M512 32H160c-35.35 0-64 28.65-64 64v224c0 35.35 28.65 64 64 64H512c35.35 0 64-28.65 64-64V96C576 60.65 547.3 32 512 32zM528 320c0 8.822-7.178 16-16 16h-16l-109.3-160.9C383.7 170.7 378.7 168 373.3 168c-5.352 0-10.35 2.672-13.31 7.125l-62.74 94.11L274.9 238.6C271.9 234.4 267.1 232 262 232c-5.109 0-9.914 2.441-12.93 6.574L176 336H160c-8.822 0-16-7.178-16-16V96c0-8.822 7.178-16 16-16H512c8.822 0 16 7.178 16 16V320zM224 112c-17.67 0-32 14.33-32 32s14.33 32 32 32c17.68 0 32-14.33 32-32S241.7 112 224 112zM456 480H120C53.83 480 0 426.2 0 360v-240C0 106.8 10.75 96 24 96S48 106.8 48 120v240c0 39.7 32.3 72 72 72h336c13.25 0 24 10.75 24 24S469.3 480 456 480z" />
        </svg>
      </button>
      {isOpen && (
        <div className={styles.modal}>
          <div className={styles.overlay} onClick={toggleModal}></div>
          <div className={styles.modal_content}>
            <img src={props.itemBaseUrl} />
            <ul>
              <li></li>
            </ul>
            <button className={styles.close_modal} onClick={toggleModal}>
              X
            </button>
          </div>
        </div>
      )}
    </>
  )
}

export default Modal
