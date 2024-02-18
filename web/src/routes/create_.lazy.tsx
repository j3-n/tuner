import { createLazyFileRoute } from '@tanstack/react-router'
import { useState } from 'react';
import useWebSocket from 'react-use-websocket';
import { Lobby } from '../types/Lobby';
import { PlayerComponent } from '../components/player';
import { Player } from '../types/Player';
import { ButtonComponent } from '../components/button';
import { State } from '../lib/State';
import { Question } from '../types/Question';
import { Result } from '../types/Result';
import { Leaderboard } from '../types/Leaderboard';
import { AnswerComponent } from '../components/answer';
import { ResultComponent } from '../components/result';
import { LeaderboardComponent } from '../components/leaderboard';
import { SongComponent } from '../components/song';
import { BackgroundComponent } from '../components/background';
import QRCode from 'react-qr-code';

export const Route = createLazyFileRoute('/create')({
  component: Page
});

function Page() {
  const socketUrl = `ws://${import.meta.env.VITE_HOST_ADDRESS}/create`;

  const [state, setState] = useState<State>(State.Waiting);
  const [lobby, setLobby] = useState<Lobby>();
  const [question, setQuestion] = useState<Question>();
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

  const onClickLeave = () => {
    getWebSocket()?.close()
    setState(State.Left)
  };

  const onClickAnswer = (id: string) => {
    const message = {
      "command": "GUESS",
      "body": {
        "answerId": id
      }
    }
    sendJsonMessage(message);
    setState(State.Answered)
  }

  return (
    <div className="h-screen">
      <div className="text-center items-center pt-20">
        <h1 className="text-slate-100 text-8xl font-bold bg-slate-800 bg-opacity-75 py-5">Lobby: {lobby?.lobbyId}</h1>
        <QRCode className="mx-auto mt-10" value={`https://${import.meta.env.VITE_WEB_ADDRESS}/${lobby?.lobbyId}`}></QRCode>

        {state === State.Waiting && lobby &&
          <div className="mt-16 gap-y-2 p-5 items-center w-1/2 mx-auto grid grid-cols-5 rounded-xl bg-opacity-75 bg-slate-800">
            {lobby && lobby.players.map((player: Player, index: number) =>
              <PlayerComponent key={index} player={player}></PlayerComponent>
            )}
          </div>}

        {state === State.Waiting &&
          <div className="fixed items-center w-1/2 bottom-0">
            <center>
              <ButtonComponent onClick={onClickPlay}>Play!</ButtonComponent>
            </center>
          </div>
        }

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

        {(state === State.Answering || state === State.Answered) && question &&
          <SongComponent src={question.question} />
        }

        {state === State.Answered &&
          <div>waiting...</div>
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

        {state != State.Waiting &&
          <div className="fixed items-center w-1/2 bottom-0">
            <center>
              <ButtonComponent onClick={onClickLeave}>Leave!</ButtonComponent>
            </center>
          </div>
        }

        {lobby &&
          <BackgroundComponent images={lobby.players[0].carousel}></BackgroundComponent>
        }
      </div>
    </div>
  );
}
