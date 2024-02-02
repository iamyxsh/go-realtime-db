import { InputHTMLAttributes } from 'react'

interface Props extends InputHTMLAttributes<HTMLInputElement> {
  label?: string
}

export const Input = ({ label, ...props }: Props) => {
  return (
    <div className="flex flex-col items-start justify-start gap-2 w-full">
      <h1 className="text-lg text-accent">{label}</h1>
      <input {...props} className="bg-slate-300 p-1 px-2 w-full rounded-md" />
    </div>
  )
}
