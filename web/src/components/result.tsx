import { Question } from "../types/Question";
import { Result } from "../types/Result";
import Cookies from "js-cookie";

type Props = {
  question: Question;
  result: Result,
  answer: string,
};

export const ResultComponent: React.FC<Props> = ({ result, question, answer }): JSX.Element => {
  let sess = Cookies.get("TUNER_SESSION");
  return (
    <>
    <div className="pt-5 text-6xl mx-auto text-slate-100">
      <ResultText answer={answer} question={question} />
    </div>
    <div className="pt-10 text-slate-100 mx-auto">
      {sess &&
        <p className="text-4xl font-bold">Your Score: {result.points[sess] || 0}</p>
      }
    </div>
    </>
  );
};

type OtherProps = {
  question: Question,
  answer: string,
}
const ResultText: React.FC<OtherProps> = ({ question, answer }): JSX.Element => {
  if(answer === "-1") {
    return <p>Too Slow!</p>
  }
  if(question.correct.toString() === answer) {
    return <p>Hook, Line and Sinker!</p>
  } else {
    return <p>Womp Womp :/</p>
  }
}
