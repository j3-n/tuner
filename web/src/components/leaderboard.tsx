import { Leaderboard } from "../types/Leaderboard";

type Props = {
  leaderboard: Leaderboard;
};

export const LeaderboardComponent: React.FC<Props> = ({ leaderboard }): JSX.Element => {
  return (
    <>
      {JSON.stringify(leaderboard)}
    </>
  );
}
