import { createLazyFileRoute } from '@tanstack/react-router';
import { Blink } from '../components/blink';
import logoImg from "../assets/logo.png"

export const Route = createLazyFileRoute('/')({
  component: Page,
})

function Page() {
  return (
    <div className="p-2 fixed justify-center items-center w-full">

      <img src={logoImg} className="h-48 rounded-3xl inline ml-10"></img>

      <h1 className="text-8xl font-bold inline ml-10">TUNER</h1>

      <div className="w-full text-slate-100 px-10 mt-10">
        <div className="align-middle gap-x-10 w-full h-96 grid grid-cols-2 mb-20">
          <div className="bg-sky-600 rounded-xl"><Blink to='/create'><p className="text-6xl">Create Game</p></Blink></div>
          <div className="bg-fuchsia-700 rounded-xl"><Blink to='/lobby'><p className="text-6xl">Join Game</p></Blink></div>
        </div>
        <div className="bg-green-600 rounded-xl"><Blink to='http://localhost:4444/login'><p className="text-6xl">Login with Spotify</p></Blink></div>
      </div>
    </div>
  );
}
