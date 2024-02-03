import { useState } from 'react'
import { Button, Input } from '../components'
import { useUserStore } from '../stores'
import { toast } from 'react-toastify'
import { useNavigate } from 'react-router-dom'

export const AuthPage = () => {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [name, setName] = useState('')

  const [type, setType] = useState<'signin' | 'signup'>('signin')

  const navigate = useNavigate()
  const signin = useUserStore((state) => state.signin)
  const signup = useUserStore((state) => state.signup)

  const handleSignin = async () => {
    try {
      await signin(email, password)
      toast.success('Sigin completed')
      navigate('/dashboard')
    } catch (err) {
      toast.error('Something went wrong')
    }
  }

  const handleSignup = async () => {
    try {
      await signup(name, email, password)
      toast.success('Sigin completed')
      navigate('/dashboard')
    } catch (err) {
      toast.error('Something went wrong')
    }
  }

  return (
    <div className="w-screen h-screen flex justify-center items-center bg-bg1">
      {type === 'signin' ? (
        <div className="w-[30%] max-w-[30rem] min-w-[10rem] flex flex-col justify-start items-start gap-5">
          <div className="flex flex-col gap-2">
            <h1 className="text-5xl text-accent">Welcome!</h1>
            <h1 className="text-2xl text-accent">Login to your account</h1>
          </div>
          <Input
            onChange={(e) => setEmail(e.target.value)}
            label="Email"
            type="email"
          />
          <Input
            onChange={(e) => setPassword(e.target.value)}
            label="Password"
            type="password"
          />
          <Button onClick={() => handleSignin()}>
            <h1 className="text-xl px-2 py-1 text-accent">Sign In</h1>
          </Button>
          <h1
            className="cursor-pointer text-lg text-accent underline underline-offset-2"
            onClick={() => setType('signup')}
          >
            Signup
          </h1>
        </div>
      ) : (
        <div className="w-[30%] max-w-[30rem] min-w-[10rem] flex flex-col justify-start items-start gap-5">
          <div className="flex flex-col gap-2">
            <h1 className="text-5xl text-accent">Welcome!</h1>
            <h1 className="text-2xl text-accent">Login to your account</h1>
          </div>
          <Input onChange={(e) => setName(e.target.value)} label="Name" />
          <Input
            onChange={(e) => setEmail(e.target.value)}
            label="Email"
            type="email"
          />
          <Input
            onChange={(e) => setPassword(e.target.value)}
            label="Password"
            type="password"
          />
          <Button onClick={() => handleSignup()}>
            <h1 className="text-xl px-2 py-1 text-accent">Sign In</h1>
          </Button>
          <h1
            className="cursor-pointer text-lg text-accent underline underline-offset-2"
            onClick={() => setType('signin')}
          >
            Signin
          </h1>
        </div>
      )}
    </div>
  )
}
