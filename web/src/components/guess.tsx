type Props = {
  text: string;
};

export const GuessComponent: React.FC<Props> = ({
  text
}): JSX.Element => {
  return (
    <>{text}</>
  );
};
