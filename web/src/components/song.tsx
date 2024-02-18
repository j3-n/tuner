import { useEffect } from "react";

type Props = {
  src: string;
};

export const SongComponent: React.FC<Props> = ({
  src,
}) => {
  useEffect(() => {
    play();
  })

  const play = () => {
    const audio = document.getElementById("player") as HTMLAudioElement;
    if (audio != null) {
      audio.play();
    }
  }

  return (
    <>
      <audio id="player" preload="false" autoPlay loop src={src} />
    </>
  );
};
