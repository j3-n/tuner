import { ButtonComponent } from "./button";

type Props = {
  orangeText: string;
  purpleText: string;
  greenText: string;
  blueText: string;
  onClick: (id: string) => void;
};

export const GuessComponent: React.FC<Props> = ({
  orangeText,
  purpleText,
  greenText,
  blueText,
  onClick,
}): JSX.Element => {
  return (
    <div className="flex flex-col h-screen">
      <div className="flex flex-col flex-grow justify-center items-center">
        <div className="flex justify-between w-full h-full">
          <ButtonComponent
            color="orange"
            onClick={() => onClick("0")}
          >
            {orangeText}
          </ButtonComponent>
          <ButtonComponent
            color="purple"
            onClick={() => onClick("1")}
          >
            {purpleText}
          </ButtonComponent>
        </div>
        <div className="flex justify-between w-full h-full">
          <ButtonComponent
            color="green"
            onClick={() => onClick("2")}
          >
            {greenText}
          </ButtonComponent>
          <ButtonComponent
            color="blue"
            onClick={() => onClick("3")}
          >
            {blueText}
          </ButtonComponent>
        </div>
      </div>
    </div>
  );
};
