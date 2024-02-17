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

  const onClick = () => {
    play();
  }

  return (
    <>
      {src}
      <button onClick={onClick}>In Case Of Emergency</button>
      <audio id="player" preload="false" autoPlay loop src="https://p.scdn.co/mp3-preview/cea8b314f8b3777c6b87a45187b7b522d3911fde?cid=79ca1b48bc314222a4699a62a130764c" />
    </>
  );
};
