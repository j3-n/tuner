import { Leaderboard, PlayerScore } from "../types/Leaderboard";
import { PlayerLeaderboardComponent } from "./player";

type Props = {
  leaderboard: Leaderboard;
};

function compare( a: PlayerScore, b: PlayerScore ) {
  if ( a.score < b.score ){
    return -1;
  }
  if ( a.score > b.score ){
    return 1;
  }
  return 0;
}

function getSorted(arr: PlayerScore[]) {
  return arr.sort(compare)
}

export const LeaderboardComponent: React.FC<Props> = ({ leaderboard }): JSX.Element => {
  console.log(leaderboard);
  return (
    <>
      <div className="grid grid-cols-7 bg-slate-800 bg-opacity-75 rounded-xl p-5 mt-10 mx-10">
      {leaderboard && getSorted(leaderboard.scores).map((score: PlayerScore, index: number) =>
        <PlayerLeaderboardComponent key={index} place={index=1} name={score.playerName} url={score.iconURL} score={score.score}></PlayerLeaderboardComponent>
      )}
      </div>
    </>
  );
}
