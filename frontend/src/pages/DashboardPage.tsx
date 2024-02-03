import { Page } from '../components'

export const DashboardPage = () => {
  return (
    <Page>
      <div className="flex h-screen bg-bg1 w-screen">
        <div className="bg-bg1 text-white w-1/5 p-4">
          <h2 className="text-3xl mb-4">Menu</h2>
          <ul>
            <li className="text-xl py-2">Project</li>
            <li className="text-xl py-2">Profile</li>
            <li className="text-xl py-2">Settings</li>
          </ul>
        </div>
      </div>
    </Page>
  )
}
