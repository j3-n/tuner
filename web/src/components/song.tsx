type Props = {
  link: string;
};

export const SongComponent: React.FC<Props> = ({
  link,
}) => {
  return (
    <>
      {link}
    </>
  );
};
