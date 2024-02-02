import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import { AuthPage, LandingPage } from './pages'

const router = createBrowserRouter([
  {
    path: '/',
    element: <LandingPage />,
  },
  {
    path: '/auth',
    element: <AuthPage />,
  },
])

function App() {
  return <RouterProvider router={router} />
}

export default App
