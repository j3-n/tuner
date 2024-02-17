import { createLazyFileRoute } from '@tanstack/react-router'
import { H1Component } from '../components/heading';
import { PlayerComponent } from '../components/player';
import {
  useState
} from 'react';
import useWebSocket from 'react-use-websocket';
import { Lobby } from '../types/Lobby';
import { Player } from '../types/Player';
import { MessageType } from '../lib/Message';

export const Route = createLazyFileRoute('/lobby/$lobbyId')({
  component: Page
});

type Response = {
  type: MessageType;
  lobby?: Lobby;
}

function Page() {
  const { lobbyId } = Route.useParams();

  const [lobby, setLobby] = useState<Lobby>();

  const [socketUrl] = useState(`ws://${import.meta.env.VITE_HOST_ADDRESS}/game/${lobbyId}`);
  useWebSocket(socketUrl, {
    onOpen: () => {
      console.log("connected");
    },
    onMessage: (event: WebSocketEventMap['message']) => {
      const message = event.data;

      const response: Response = JSON.parse(message);
      if (response.type == MessageType.Lobby) {
        setLobby(response.lobby)
      }
    },
  });

  return (
    <div className="text-center items-center">
      <H1Component>lobby {lobbyId}</H1Component>

      {lobby && lobby.players.map((player: Player, index: number) =>
        <div key={index}>
          <PlayerComponent player={player}></PlayerComponent>
        </div>
      )}
    </div>
  );
}
