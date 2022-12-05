import Link from 'next/link'

type Props = {
    title: string
}

const Title: React.FC<Props> = props => {
  return (
    <Link href="/">
      <a className="title">
        <h2>{props.title}</h2>
      </a>
    </Link>
  )
}
export default Title
