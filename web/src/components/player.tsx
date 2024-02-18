import { ReactNode } from "react";
import { Player } from "../types/Player";
import { ImageComponent } from "./image";

type Props = {
  player: Player;
  score?: number,
  showScore: boolean,
  children?: ReactNode;
};

export const PlayerComponent: React.FC<Props> = ({
  player,
  score,
  showScore,
  children
}): JSX.Element => {
  return (
    <>
    <div className="col-span">
      {player.iconURL &&
        <ImageComponent
          src={player.iconURL}
          width={100}
          height={100}
          rounded={true}
        />}
      {children}
    </div>
    <div className="col-span-4 text-slate-100">
      <h3 className="text-5xl font-bold">{player.displayName}</h3>
    </div>
    {
      showScore &&
      <div className="col-span-1 text-slate-100">
        <h3 className="text-5xl font-bold">{score}</h3>
      </div>
    }
    </>
  );
};

type OtherProps = {
  place: number,
  url: string,
  score: number,
  name: string,
};

export const PlayerLeaderboardComponent: React.FC<OtherProps> = ({
  place,
  url,
  score,
  name
}): JSX.Element => {
  return (
    <>
    <div className="col-span">
      {url &&
        <ImageComponent
          src={url}
          width={100}
          height={100}
          rounded={true}
        />}
    </div>
    <div className="col-span-1 text-slate-100 text-6xl font-bold">
      {place}
    </div>
    <div className="col-span-4 text-slate-100">
      <h3 className="text-5xl font-bold">{name}</h3>
    </div>
    <div className="col-span-1 text-slate-100">
      <h3 className="text-5xl font-bold">{score}</h3>
    </div>
    </>
  );
};
