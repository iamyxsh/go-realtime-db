type Props = {
  name: string
  img: string
  points: [string]
}

export const FeatureCard = ({ name, img, points }: Props) => {
  return (
    <div className=" flex flex-col rounded-md gap-10 border-solid border-2 border-primaryColor w-[50%] h-[25rem] p-10">
      <img src={img} className="w-[20%]" />
      <h1 className="text-2xl text-primaryColor text-start">{name}</h1>
    </div>
  )
}
