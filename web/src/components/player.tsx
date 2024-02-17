import { ReactNode } from "react";
import { H3Component } from "./heading";
import { Player } from "../types/Player";

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
      {children}
    </div>
  );
};
