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
    <div className="grid grid-cols-2 gap-5 mt-20">
      <ButtonComponent
        onClick={() => onClick("0")}
      >
        <div className="h-20 rounded-xl bg-red-500 flex items-center">
            <p className="mx-auto">{orangeText}</p>
        </div>
      </ButtonComponent>
      <ButtonComponent
        onClick={() => onClick("1")}
      >
        <div className="h-20 rounded-xl bg-yellow-500 flex items-center">
        <p className="mx-auto">{purpleText}</p>
        </div>
      </ButtonComponent>
      <ButtonComponent
        onClick={() => onClick("2")}
      >
        <div className="h-20 rounded-xl bg-green-500 flex items-center">
        <p className="mx-auto">{greenText}</p>
        </div>
      </ButtonComponent>
      <ButtonComponent
        onClick={() => onClick("3")}
      >
        <div className="h-20 rounded-xl bg-blue-500 flex items-center">
        <p className="mx-auto">{blueText}</p>
        </div>
      </ButtonComponent>
    </div>
  );
};
