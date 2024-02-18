import { createLazyFileRoute } from '@tanstack/react-router'
import { H1Component } from '../components/heading';
import { PlayerComponent } from '../components/player';
import { useState } from 'react';
import useWebSocket from 'react-use-websocket';
import { Lobby } from '../types/Lobby';
import { Player } from '../types/Player';
import { ButtonComponent } from '../components/button';
import { State } from '../lib/State';
import { AnswerComponent } from '../components/answer';
import { Question } from '../types/Question';
import { Leaderboard } from '../types/Leaderboard';
import { Result } from '../types/Result';
import { ResultComponent } from '../components/result';
import { LeaderboardComponent } from '../components/leaderboard';

export const Route = createLazyFileRoute('/lobby/$lobbyId')({
  component: Page
});

function Page() {
  const { lobbyId } = Route.useParams();
  const socketUrl = `ws://${import.meta.env.VITE_HOST_ADDRESS}/play/${lobbyId}`;

  const [state, setState] = useState<State>(State.Waiting);
  const [lobby, setLobby] = useState<Lobby>();
  const [question, setQuestion] = useState<Question>();
  const [answer, setAnswer] = useState<string>("-1");
  const [result, setResult] = useState<Result>();
  const [leaderboard, setLeaderboard] = useState<Leaderboard>();

  const { sendJsonMessage, getWebSocket } = useWebSocket(socketUrl, {
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
            setAnswer("-1");
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

  const onClickLeave = () => {
    getWebSocket()?.close()
    setState(State.Left)
  }

  const onClickAnswer = (id: string) => {
    const message = {
      "command": "GUESS",
      "body": {
        "answerId": id
      }
    }
    sendJsonMessage(message);
    setState(State.Answered);
    setAnswer(id);
  }

  return (
    <div className="max-h-screen">
      <div className="text-center items-center pt-20">
        <H1Component>lobby {lobbyId}</H1Component>
      </div>

      {state === State.Waiting && lobby &&
        <div className="pt-20 items-center w-1/2 grid grid-flow-col">
          {lobby.players.map((player: Player, index: number) =>
            <div key={index}>
              <PlayerComponent player={player}></PlayerComponent>
            </div>
          )}
        </div>}


      {state === State.Answering && question &&
        <div>
          <AnswerComponent
            orangeText={question.answers[0].song + " - " + question.answers[0].artist}
            purpleText={question.answers[1].song + " - " + question.answers[1].artist}
            greenText={question.answers[2].song + " - " + question.answers[2].artist}
            blueText={question.answers[3].song + " - " + question.answers[3].artist}
            onClick={onClickAnswer}
          >
          </AnswerComponent>
        </div>
      }

      {state === State.Result && question && result &&
        <div>
            <ResultComponent
              question={question}
              result={result}
              answer={answer}
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
