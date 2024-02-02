import { Button, Input } from '../components'

export const AuthPage = () => {
  return (
    <div className="w-screen h-screen flex justify-center items-center bg-bg1">
      <div className="w-[30%] max-w-[30rem] min-w-[10rem] flex flex-col justify-start items-start gap-5">
        <div className="flex flex-col gap-2">
          <h1 className="text-5xl text-accent">Welcome!</h1>
          <h1 className="text-2xl text-accent">Signin to your account</h1>
        </div>
        <Input label="Email" type="email" />
        <Input label="Password" type="password" />
        <Button>
          <h1 className="text-xl px-2 py-1 text-accent">Sign In</h1>
        </Button>
      </div>
    </div>
  )
}
