type Props = {
  images: string[];
}

export const BackgroundComponent: React.FC<Props> = ({
  images,
}): JSX.Element => {
  return (
    <div className=" opacity-25 grid grid-cols-2 sm:grid-cols-3 md:grid-cols-3 lg:grid-cols-4 gap-4 w-full h-full overflow-hidden fixed top-0 left-0 -z-10">
      {images.map((image: string, index: number) =>
        <img src={image} key={index}></img>
      )}
    </div>
  );
};
