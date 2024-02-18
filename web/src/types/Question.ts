import { Answer } from "./Answer";

export type Question = {
  question: string;
  answers : Answer[];
  correct: number;
};
