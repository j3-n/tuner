import { Player } from "./Player";

export type Result = {
  lobby: string,
  players: Player[],
  points: Record<string, number>,
};