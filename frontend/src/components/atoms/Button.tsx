import { ButtonHTMLAttributes } from 'react'
import { useNavigate } from 'react-router-dom'

interface Props extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: 'filled' | 'outline'
  to?: string
}

export const Button = ({ children, to, onClick }: Props) => {
  const navigate = useNavigate()

  return (
    <button
      onClick={onClick ? onClick : () => navigate(to!)}
      className="bg-bg1 text-lg px-2 py-[1px] border-solid border-accent border-[1px] rounded-md cursor-pointer"
    >
      {children}
    </button>
  )
}
