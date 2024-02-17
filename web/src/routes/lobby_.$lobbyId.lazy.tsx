import { createLazyFileRoute } from '@tanstack/react-router'
import { H1Component } from '../components/heading';
import { PlayerComponent } from '../components/player';
import {
  useState
} from 'react';
import useWebSocket from 'react-use-websocket';
import { Lobby } from '../types/Lobby';
import { Player } from '../types/Player';
import { ButtonComponent } from '../components/button';
import { State } from '../lib/State';
import { GuessComponent } from '../components/guess';
import { Question } from '../types/Question';

export const Route = createLazyFileRoute('/lobby/$lobbyId')({
  component: Page
});

function Page() {
  const { lobbyId } = Route.useParams();
  const socketUrl = `ws://${import.meta.env.VITE_HOST_ADDRESS}/play/${lobbyId}`;

  const [state, setState] = useState<State>(State.Waiting);
  const [lobby, setLobby] = useState<Lobby>();
  const [question, setQuestion] = useState<Question>();

  const { getWebSocket } = useWebSocket(socketUrl, {
    onOpen: () => {
      console.log("connected");
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

  const onClickLeave = () => {
    getWebSocket()?.close()
  }

  const onClickGuess = () => {
    console.log("guessed")
  }

  return (
    <div className="max-h-screen">
      <div className="text-center items-center pt-20">
        <H1Component>lobby {lobbyId}</H1Component>
      </div>

      {state === State.Waiting &&
        <div className="text-center items-center pt-20">
          <div className="pt-20 items-center w-1/2 grid grid-flow-col">
            {lobby && lobby.players.map((player: Player, index: number) =>
              <div key={index}>
                <PlayerComponent player={player}></PlayerComponent>
              </div>
            )}
          </div>
        </div>
      }

      {state === State.Guessing && question &&
        <div>
          <GuessComponent question={question} text="yo"></GuessComponent>
        </div>}

      <div className="fixed items-center w-1/2 bottom-0">
        <center>
          <ButtonComponent onClick={onClickLeave}>Leave!</ButtonComponent>
        </center>
      </div>
    </div>
  );
}
