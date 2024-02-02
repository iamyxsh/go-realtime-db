type Props = {
  name: string
  img: string
  points: [string]
}

export const FeatureCard = ({ name, img, points }: Props) => {
  return (
    <div className=" flex flex-col items-center justify-start rounded-md gap-10 border-solid border-2 border-primaryColor w-[20%] h-[25rem] min-w-[20rem] max-w-[30rem] p-10">
      <img src={img} className="w-[5rem]" />
      <h1 className="text-2xl text-primaryColor text-start">{name}</h1>
      <div className="flex flex-col gap-5">
        {points.map((point) => {
          return (
            <div className="flex justify-start items-center  gap-2">
              <div className="rounded-full w-[9px] h-[8px] bg-black" />
              <h1 className="text-lg">{point}</h1>
            </div>
          )
        })}
      </div>
    </div>
  )
}
