import { ButtonComponent } from "./button";

type Props = {
  orangeText: string;
  purpleText: string;
  greenText: string;
  blueText: string;
  onClick: (id: string) => void;
};

export const AnswerComponent: React.FC<Props> = ({
  orangeText,
  purpleText,
  greenText,
  blueText,
  onClick,
}): JSX.Element => {
  return (
    <div className="flex flex-col h-screen">
      <div className="flex flex-col flex-grow justify-center items-center">
        <div className="fixed justify-between w-1/2 h-full">
          <div className=" bg-orange-500">
            <ButtonComponent
              onClick={() => onClick("0")}
            >
              {orangeText}
            </ButtonComponent>
          </div>
          <div className=" bg-purple-500">
            <ButtonComponent
              onClick={() => onClick("1")}
            >
              {purpleText}
            </ButtonComponent>
          </div>
          <div className="bg-green-500">
            <ButtonComponent
              onClick={() => onClick("2")}
            >
              {greenText}
            </ButtonComponent>
          </div>
          <div className="bg-blue-500">
            <ButtonComponent
              onClick={() => onClick("3")}
            >
              {blueText}
            </ButtonComponent>
          </div>
        </div>
      </div>
    </div>
  );
};
