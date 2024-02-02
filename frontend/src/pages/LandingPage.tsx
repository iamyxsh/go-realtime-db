import { Page } from '../components'
import { ExampleCard, FeatureCard } from '../components/molecules'
import { features, startBuilding } from '../constants'

export const LandingPage = () => {
  return (
    <Page>
      <div className="flex flex-col w-screen">
        <div className="w-screen h-[50rem] bg-accent flex flex-col gap-10 justify-center items-center">
          <h1 className="text-7xl text-primaryColor">Empower Your App</h1>
          <h1 className="text-4xl text-bg1">
            Use our service to build your app in a weekend.
          </h1>
        </div>
        <div className="flex flex-col mt-[5rem]">
          <div className="flex gap-10 justify-center items-center p-10">
            {features.map((feat) => (
              <FeatureCard
                img={feat.img}
                name={feat.name}
                points={feat.points as [string]}
                key={feat.name}
              />
            ))}
          </div>
        </div>
        <div className="flex flex-col items-center justify-start mt-[5rem] bg-bg1 p-[5rem]">
          <h1 className="text-5xl text-accent">Start Building</h1>
          <div className="flex gap-10 justify-center items-center p-10">
            {startBuilding.map((eg) => (
              <ExampleCard
                gh={eg.gh}
                name={eg.name}
                link={eg.link}
                description={eg.description}
                key={eg.name}
              />
            ))}
          </div>
        </div>
      </div>
    </Page>
  )
}
