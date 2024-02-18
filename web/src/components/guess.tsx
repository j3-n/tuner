import { SendJsonMessage } from "react-use-websocket/dist/lib/types";
import { ButtonComponent } from "./button";

type Props = {
  orangeText: string;
  purpleText: string;
  greenText: string;
  blueText: string;
  sendJsonMessage: SendJsonMessage;
};

export const GuessComponent: React.FC<Props> = ({
  orangeText,
  purpleText,
  greenText,
  blueText,
  sendJsonMessage,
}): JSX.Element => {
  const onClickOrange = () => {
    const message = {
      "command": "GUESS",
      "body": {
        "answerId": "0"
      }
    }
    sendJsonMessage(message);
  }

  const onClickPurple = () => {
    const message = {
      "command": "GUESS",
      "body": {
        "answerId": "1"
      }
    }
    sendJsonMessage(message);
  }

  const onClickGreen = () => {
    const message = {
      "command": "GUESS",
      "body": {
        "answerId": "2"
      }
    }
    sendJsonMessage(message);
  }

  const onClickBlue = () => {
    const message = {
      "command": "GUESS",
      "body": {
        "answerId": "3"
      }
    }
    sendJsonMessage(message);
  }

  return (
    <div className="flex flex-col h-screen">
      <div className="flex flex-col flex-grow justify-center items-center">
        <div className="flex justify-between w-full h-full">
          <ButtonComponent
            color="orange"
            onClick={onClickOrange}
          >
            {orangeText}
          </ButtonComponent>
          <ButtonComponent
            color="purple"
            onClick={onClickPurple}
          >
            {purpleText}
          </ButtonComponent>
        </div>
        <div className="flex justify-between w-full h-full">
          <ButtonComponent
            color="green"
            onClick={onClickGreen}
          >
            {greenText}
          </ButtonComponent>
          <ButtonComponent
            color="blue"
            onClick={onClickBlue}
          >
            {blueText}
          </ButtonComponent>
        </div>
      </div>
    </div>
  );
};
