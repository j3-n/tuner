type Props = {
  images: string[];
}

export const BackgroundComponent: React.FC<Props> = ({
  images,
}): JSX.Element => {
  return (
    <div className="bg-slate-800 overflow-none brightness-[0.4] w-full h-full overflow-hidden fixed top-0 left-0 -z-10">
      {images.map((image: string, index: number) =>
        <img className="blur w-1/3 inline" src={image} key={index}></img>
      )}
    </div>
  );
};
