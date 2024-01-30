import Navbar from '../molecules/Navbar'

interface Props {
  children: React.ReactNode
}

export const Page = ({ children }: Props) => {
  return (
    <main className="flex flex-col justify-start items-center">
      <Navbar />
      <div>{children}</div>
    </main>
  )
}
