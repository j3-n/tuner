import { Question } from "../types/Question";
import { ButtonComponent } from "./button";

type Props = {
  text: string;
  question: Question;
};

export const GuessComponent: React.FC<Props> = ({
  text
}): JSX.Element => {
  return (
    <>
      {text}

      <ButtonComponent>one</ButtonComponent>
      <ButtonComponent>two</ButtonComponent>
      <ButtonComponent>three</ButtonComponent>
      <ButtonComponent>four</ButtonComponent>
    </>
  );
};
