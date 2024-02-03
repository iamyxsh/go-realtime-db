import { authentication, database } from '../assets'

export const features = [
  {
    name: 'Database',
    img: database,
    points: ['Realtime', 'Postgres', 'Easy-to-use API'],
  },
  {
    name: 'Authentication',
    img: authentication,
    points: ['Signup and Login', 'JWT Verification', 'Easy-to-use API'],
  },
]

export const startBuilding = [
  {
    name: 'Database Example',
    link: '',
    gh: '',
    description:
      'Take a look at this example to see how to work with real-time DB',
  },
  {
    name: 'Authentication Example',
    link: '',
    gh: '',
    description:
      'Take a look at this example to see how to work with authentication',
  },
]

export const navbarOptions = [
  {
    name: 'Features',
    link: '/#features',
  },
  {
    name: 'Documentation',
    link: '/docs',
  },
  {
    name: 'Get Started',
    link: '/auth',
  },
]
