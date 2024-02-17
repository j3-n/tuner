import { createLazyFileRoute } from '@tanstack/react-router'
import { InputComponent } from '../components/input';
import { LinkComponent } from '../components/link';
import {
  useEffect,
  useState
} from 'react';

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
      </div>
    </div>
  );
}
