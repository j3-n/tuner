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
import { GuessComponent } from '../components/guess';
import { ResultComponent } from '../components/result';
import { LeaderboardComponent } from '../components/leaderboard';

export const Route = createLazyFileRoute('/create/$lobbyId')({
  component: Page
});

function Page() {
  const { lobbyId } = Route.useParams();
  const socketUrl = `ws://${import.meta.env.VITE_HOST_ADDRESS}/create`;

  const [state, setState] = useState<State>(State.Waiting);
  const [lobby, setLobby] = useState<Lobby>();
  const [question, setQuestion] = useState<Question>();
  const [result, setResult] = useState<Result>();
  const [leaderboard, setLeaderboard] = useState<Leaderboard>();

  const { sendJsonMessage } = useWebSocket(socketUrl, {
    onOpen: () => {
      console.log("connected")
    },
    onMessage: (event: WebSocketEventMap['message']) => {
      const message = event.data;
      const command = JSON.parse(message);

      console.log(JSON.stringify(command));

      try {
        switch (command?.command) {
          case "WAITING":
            // the default state
            console.log(command.body)
            setLobby(command.body as Lobby);
            break;
          case "QUESTION":
            // follows the default state, will alternate with result state till finished state
            setQuestion(command.body as Question);
            setState(State.Answering);
            break;
          case "RESULT":
            // result state, shows the current user points
            setResult(command.body as Result);
            setState(State.Result);
            break;
          case "FINISHED":
            // shows the end of game leaderboard
            setLeaderboard(command.body as Leaderboard);
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
    const message = {
      command: "START",
      body: "",
    };
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
            <ButtonComponent onClick={onClickPlay}>Play!</ButtonComponent>
          </center>
        </div>
      </div>
    </div>
  );
}
