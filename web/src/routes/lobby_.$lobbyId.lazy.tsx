import { createLazyFileRoute } from '@tanstack/react-router'
import { H1Component } from '../components/heading';
import { PlayerComponent } from '../components/player';
import { useState } from 'react';
import useWebSocket from 'react-use-websocket';
import { Lobby } from '../types/Lobby';
import { Player } from '../types/Player';
import { ButtonComponent } from '../components/button';
import { State } from '../lib/State';
import { GuessComponent } from '../components/guess';
import { Question } from '../types/Question';
import { Leaderboard } from '../types/Leaderboard';
import { Result } from '../types/Result';
import { ResultComponent } from '../components/result';
import { LeaderboardComponent } from '../components/leaderboard';

export const Route = createLazyFileRoute('/lobby/$lobbyId')({
  component: Page
});

type Command = {
  command: string;
  body: string;
};

function Page() {
  const { lobbyId } = Route.useParams();
  const socketUrl = `ws://${import.meta.env.VITE_HOST_ADDRESS}/play/${lobbyId}`;

  const [state, setState] = useState<State>(State.Waiting);
  const [command, setCommand] = useState<Command>();
  const [lobby, setLobby] = useState<Lobby>();
  const [question, setQuestion] = useState<Question>();
  const [result, setResult] = useState<Result>();
  const [leaderboard, setLeaderboard] = useState<Leaderboard>();

  const { getWebSocket } = useWebSocket(socketUrl, {
    onOpen: () => {
      console.log("connected");
    },
    onMessage: (event: WebSocketEventMap['message']) => {
      const message = event.data;
      console.log(JSON.stringify(message));
      try {
        setCommand(JSON.parse(message));
      } catch (error) {
        console.log("error parsing json", error)
      }

      if (command == null) {
        return;
      }

      try {
        switch (command.command) {
          case "WAITING":
            // the defa(ult state
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

  const onClickLeave = () => {
    getWebSocket()?.close()
  }

  return (
    <div className="max-h-screen">
      <div className="text-center items-center pt-20">
        <H1Component>lobby {lobbyId}</H1Component>
      </div>

      {state === State.Waiting && lobby &&
        <div className="pt-20 items-center w-1/2 grid grid-flow-col">
          {JSON.stringify(lobby)}
          {lobby.players.map((player: Player, index: number) =>
            <div key={index}>
              <PlayerComponent player={player}></PlayerComponent>
            </div>
          )}
        </div>}


      {state === State.Answering && question &&
        <div>
          <GuessComponent
            text='yo'
          >
          </GuessComponent>
        </div>
      }

      {state === State.Result && result &&
        <div>
          <ResultComponent
            result={result}
          >
          </ResultComponent>
        </div>
      }

      {state === State.Finished && leaderboard &&
        <div>
          <LeaderboardComponent
            leaderboard={leaderboard}
          >
          </LeaderboardComponent>
        </div>
      }

      <div className="fixed items-center w-1/2 bottom-0">
        <center>
          <ButtonComponent onClick={onClickLeave}>Leave!</ButtonComponent>
        </center>
      </div>
    </div>
  );
}
