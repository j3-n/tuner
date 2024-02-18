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
  const baseUrl = `http://${import.meta.env.VITE_WEB_ADDRESS}/lobby`
  const [lobbyCode, setLobbyCode] = useState<string>();

  useEffect(() => {
    console.log(lobbyCode)
  }, [lobbyCode])

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setLobbyCode(event.target.value);
  };

  return (
    <div>
        <InputComponent placeholder="Lobby Code" id="input" onChange={handleInputChange} /><br />
        <div className="bg-green-600 rounded-xl w-1/3 h-20 mx-auto">
          <LinkComponent to={`${baseUrl}/${lobbyCode}`} size='lg'>
            <p className="text-4xl text-slate-100">
              Join
            </p>
          </LinkComponent>
        </div>
    </div>
  );
}
