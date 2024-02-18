import { ReactNode } from "react";
import { Player } from "../types/Player";
import { ImageComponent } from "./image";

type Props = {
  player: Player;
  children?: ReactNode;
};

export const PlayerComponent: React.FC<Props> = ({
  player,
  children
}): JSX.Element => {
  return (
    <>
    <div className="col-span-1">
      {player.iconURL &&
        <ImageComponent
          src={player.iconURL}
          width={100}
          height={100}
          rounded={true}
        />}
      {children}
    </div>
    <div className="col-span-4">
      <h3 className="text-5xl font-bold">{player.displayName}</h3>
    </div>
    </>
  );
};
