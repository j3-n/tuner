type Props = {
  src: string;
};

export const ImageComponent: React.FC<Props> = ({
  src,
}): JSX.Element => {
  return (
    <>
      <img src={src}></img>
    </>
  );
};
