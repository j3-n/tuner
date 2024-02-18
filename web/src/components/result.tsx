import { Result } from "../types/Result";

type Props = {
  result: Result;
};

export const ResultComponent: React.FC<Props> = ({ result }): JSX.Element => {
  return (
    <>{JSON.stringify(result)}</>
  );
};
