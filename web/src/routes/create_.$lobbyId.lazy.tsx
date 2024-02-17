import { createLazyFileRoute } from '@tanstack/react-router'
import { useState } from 'react';
import useWebSocket from 'react-use-websocket';
import { Lobby } from '../types/Lobby';
import { H1Component } from '../components/heading';
import { PlayerComponent } from '../components/player';
import { Player } from '../types/Player';
import { ButtonComponent } from '../components/button';
import { Command } from '../types/Command';

export const Route = createLazyFileRoute('/create/$lobbyId')({
  component: Page
});

function Page() {
  const { lobbyId } = Route.useParams();
  const socketUrl = `ws://${import.meta.env.VITE_HOST_ADDRESS}/create`;

  const [lobby, setLobby] = useState<Lobby>();

  const { sendMessage, sendJsonMessage } = useWebSocket(socketUrl, {
    onOpen: () => {
      console.log("connected")
    },
    onMessage: (event: WebSocketEventMap['message']) => {
      const message = event.data;
      console.log(message);
      setLobby(JSON.parse(message));
    },
    onClose: () => {
      console.log("disconnected")
    }
  });

  const onClickPlay = () => {
    const message: Command = {
      command: "play",
      body: "",
    };
    sendMessage("quiz");
    sendJsonMessage(message);
  };

  return (
    <div className="max-h-screen">
      <div className="text-center items-center">
        <H1Component>lobby {lobbyId}</H1Component>

        <div className="pt-20 items-center w-1/2 grid grid-flow-col">
          {lobby && lobby.players.map((player: Player, index: number) =>
            <div key={index}>
              <PlayerComponent player={player}></PlayerComponent>
            </div>
          )}
        </div>

        <div className="fixed items-center w-1/2 bottom-0">
          <center>
            <ButtonComponent onClick={onClickPlay}>Play!</ButtonComponent>
          </center>
        </div>
      </div>
    </div>
  );
}
