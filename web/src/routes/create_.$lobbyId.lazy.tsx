import { createLazyFileRoute } from '@tanstack/react-router'
import { useState } from 'react';
import useWebSocket from 'react-use-websocket';
import { Lobby } from '../types/Lobby';
import { H1Component } from '../components/heading';
import { PlayerComponent } from '../components/player';
import { Player } from '../types/Player';
import { ButtonComponent } from '../components/button';
import { State } from '../lib/State';
import { Question } from '../types/Question';
import { Result } from '../types/Result';
import { Leaderboard } from '../types/Leaderboard';

export const Route = createLazyFileRoute('/create/$lobbyId')({
  component: Page
});

type Command = {
  command: string;
  body: string;
};

function Page() {
  const { lobbyId } = Route.useParams();
  const socketUrl = `ws://${import.meta.env.VITE_HOST_ADDRESS}/create`;

  const [state, setState] = useState<State>(State.Waiting);
  const [command, setCommand] = useState<Command>();
  const [lobby, setLobby] = useState<Lobby>();
  const [question, setQuestion] = useState<Question>();
  const [result, setResult] = useState<Result>();
  const [leaderboard, setLeaderboard] = useState<Leaderboard>();

  const { sendMessage, sendJsonMessage } = useWebSocket(socketUrl, {
    onOpen: () => {
      console.log("connected")
    },
    onMessage: (event: WebSocketEventMap['message']) => {
      const message = event.data;
      console.log(message);
      setCommand(JSON.parse(message));
      console.log(command);

      console.log(JSON.stringify(command));

      try {
        switch (command?.command) {
          case "WAITING":
            // the defa(ult state
            console.log(command.body)
            setLobby(JSON.parse(command.body));
            break;
          case "QUESTION":
            // follows the default state, will alternate with result state till finished state
            setQuestion(JSON.parse(command.body));
            setState(State.Answering);
            break;
          case "RESULT":
            // result state, shows the current user points
            setResult(JSON.parse(command.body));
            setState(State.Result);
            break;
          case "FINISHED":
            // shows the end of game leaderboard
            setLeaderboard(JSON.parse(command.body));
            setState(State.Finished);
            break;
        }
      } catch (error) {
        console.log("error parsing command,", error)
      }
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

        {state === State.Waiting && lobby &&
          <div className="pt-20 items-center w-1/2 grid grid-flow-col">
            {lobby && lobby.players.map((player: Player, index: number) =>
              <div key={index}>
                <PlayerComponent player={player}></PlayerComponent>
              </div>
            )}
          </div>}

        <div className="fixed items-center w-1/2 bottom-0">
          <center>
            <ButtonComponent onClick={onClickPlay}>Play!</ButtonComponent>
          </center>
        </div>
      </div>
    </div>
  );
}
