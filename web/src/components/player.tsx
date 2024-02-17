import { ReactNode } from "react";
import { H3Component } from "./heading";
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
    <div>
      <H3Component>{player.displayName}</H3Component>
      {player.iconURL &&
        <ImageComponent
          src={player.iconURL}
          width={100}
          height={100}
          rounded={true}
        />}
      {children}
    </div>
  );
};
