import { ButtonHTMLAttributes } from 'react'

interface Props extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: 'filled' | 'outline'
}

export const Button = ({ variant, children }: Props) => {
  return (
    <div className="bg-bg1 text-lg px-2 py-[1px] border-solid border-accent border-[1px] rounded-md">
      {children}
    </div>
  )
}
