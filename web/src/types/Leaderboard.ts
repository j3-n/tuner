export type Leaderboard = {
  scores: PlayerScore[],
};

export type PlayerScore = {
  playerName: string,
  iconURL: string,
  score: number
};
