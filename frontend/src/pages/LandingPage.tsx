import { authentication, database } from '../assets'
import { Page } from '../components'
import { FeatureCard } from '../components/molecules/FeatureCard'

export const feature = [
  {
    name: 'Database',
    img: database,
    points: ['Realtime', 'SQL', 'Easy-to-use API'],
  },
  {
    name: 'Authentication',
    img: authentication,
    points: ['Realtime', 'SQL', 'Easy-to-use API'],
  },
]

const LandingPage = () => {
  return (
    <Page>
      <div className="flex flex-col w-screen">
        <div className="w-screen h-[50rem] bg-accent flex flex-col gap-10 justify-center items-center">
          <h1 className="text-7xl text-primaryColor">Empower Your App</h1>
          <h1 className="text-4xl text-bg1">
            Use our service to build your app in a weekend.
          </h1>
        </div>
        <div className="flex flex-col">
          <div className="flex gap-10 justify-center items-center p-10">
            {feature.map((feat) => (
              <FeatureCard
                img={feat.img}
                name={feat.name}
                points={feat.points}
                key={feat.name}
              />
            ))}
          </div>
        </div>
      </div>
    </Page>
  )
}

export default LandingPage
