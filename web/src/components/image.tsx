type Props = {
  src: string;
  width: number;
  height: number;
  rounded?: boolean;
};

export const ImageComponent: React.FC<Props> = ({
  src,
  width,
  height,
  rounded = false,
}): JSX.Element => {
  return (
    <div className={`w-${width} h-${height} relative`}>
      <img
        src={src}
        className={`absolute inset-0 w-full h-full object-cover ${rounded ? 'rounded-full' : ''}`}
        width={width}
        height={height}
        alt="profile picture"
      />
    </div>
  );
};
