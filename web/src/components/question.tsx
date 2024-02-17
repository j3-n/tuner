import { Question } from "../types/Question";

type Props = {
  question: Question;
};

export const QuestionComponent: React.FC<Props> = ({
  question,
}): JSX.Element => {
  return (
    <>{JSON.stringify(question)}</>
  );
}
