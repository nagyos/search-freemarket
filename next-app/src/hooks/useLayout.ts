import { useCallback, useEffect, useRef, useState } from 'react'

export const useLayout = () => {
  const [headerHeight, setHeaderHeight] = useState(0)
  const [itemContainerHeight, setItemContainerHeight] = useState(0)
  const headerRef = useRef<HTMLElement>(null)
  const setLayout = useCallback(() => {
    const header = headerRef.current
    if (header) {
      setHeaderHeight(header.clientHeight)
      setItemContainerHeight(window.innerHeight - header.clientHeight)
    }
  }, [])
  useEffect(() => {
    setLayout()
    window.addEventListener('resize', setLayout)
    return () => window.removeEventListener('resize', setLayout)
  }, [])
  return {
    headerHeight,
    headerRef,
    itemContainerHeight
  }
}
