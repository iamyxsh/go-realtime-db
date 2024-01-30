import { Link } from 'react-router-dom'
import { logo } from '../../assets'

const navbarOptions = [
  {
    name: 'Features',
    link: '/#features',
  },
  {
    name: 'Documentation',
    link: '/docs',
  },
  {
    name: 'Signup',
  },
]

const Navbar = () => {
  return (
    <nav className="w-full bg-bg1 text-accent text-xl p-2">
      <div className="w-1/2 m-auto flex justify-between items-center">
        <div className="flex gap-2 justify-center items-center">
          <img src={logo} className="w-7" />
          tokyo
        </div>
        <div className="flex gap-5 items-center">
          {navbarOptions.map((item) => {
            if (item.link) {
              return (
                <Link to={item.link} className="text-[1rem]">
                  {item.name}
                </Link>
              )
            } else {
              return <button className="text-[1rem]">{item.name}</button>
            }
          })}
        </div>
      </div>
    </nav>
  )
}

export default Navbar
