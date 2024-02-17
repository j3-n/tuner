import { createLazyFileRoute } from '@tanstack/react-router'
import { ButtonComponent } from '../components/button';
import { InputComponent } from '../components/input';
import { useState } from 'react';

export const Route = createLazyFileRoute('/lobby')({
  component: Page
});

function Page() {
  const [lobbyCode, setLobbyCode] = useState('');

  return (
    <div>
      <h1 className="">Hello /lobby!</h1>

      <div className="flex">
        <InputComponent placeholder="lobby code" />
        <ButtonComponent onClick={() => console.log(lobbyCode)}>hello</ButtonComponent>
      </div>
    </div>
  );
}
