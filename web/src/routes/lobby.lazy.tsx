import { createLazyFileRoute } from '@tanstack/react-router'
import { InputComponent } from '../components/input';
import { LinkComponent } from '../components/link';
import {
  useEffect,
  useState
} from 'react';
import { SongComponent } from '../components/song';

export const Route = createLazyFileRoute('/lobby')({
  component: Page
});

function Page() {
  const [lobbyCode, setLobbyCode] = useState<string>("");

  useEffect(() => {
    console.log(lobbyCode)
  }, [lobbyCode])

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setLobbyCode(event.target.value);
  };

  return (
    <div>
      <h1 className="">Hello /lobby!</h1>

      <div className="flex">
        <InputComponent placeholder="lobby code" id="input" onChange={handleInputChange} />
        <LinkComponent to={`/lobby/${lobbyCode}`}>go</LinkComponent>
        <SongComponent src="https://p.scdn.co/mp3-preview/cea8b314f8b3777c6b87a45187b7b522d3911fde?cid=79ca1b48bc314222a4699a62a130764c"></SongComponent>
      </div>
    </div>
  );
}
