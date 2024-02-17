import { useEffect, useState } from "react";

type Props = {
  src: string;
};

export const SongComponent: React.FC<Props> = ({
  src,
}) => {
  const [playing, setPlaying] = useState<boolean>(true);

  useEffect(() => {
    const play = async () => {
      if (playing) {
        await new Audio(src).play();
        console.log(playing)
      }
    }

    play()
  })

  const onClick = () => {
    setPlaying(!playing)
  }

  return (
    <>
      {src}
      <button onClick={onClick}>button</button>
    </>
  );
};
