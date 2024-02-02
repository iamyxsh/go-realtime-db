type Props = {
  name: string
  link: string
  gh: string
  description: string
}

export const ExampleCard = ({ name, link, gh, description }: Props) => {
  return (
    <div className="bg-accent flex flex-col items-center justify-start rounded-md gap-10 border-solid border-2 border-primaryColor w-[50%] h-[15rem] max-w-[30rem] p-10">
      <h1 className="text-2xl text-primaryColor text-start">{name}</h1>
      <div className="flex flex-col gap-5">
        <h1 className="text-lg">{description}</h1>
        <div className="flex justify-between items-center">
          <a href={link} className="text-bg1 text-xl">
            Demo
          </a>
          <a href={gh} className="text-bg1 text-xl">
            Github
          </a>
        </div>
      </div>
    </div>
  )
}
