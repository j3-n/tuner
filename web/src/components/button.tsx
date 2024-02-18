import { ReactNode } from "react";

type Props = {
  onClick?: () => void;
  color?: string;
  children?: ReactNode;
}

export const ButtonComponent: React.FC<Props> = ({
  onClick,
  color = "",
  children
}): JSX.Element => {
  return (
    <div>
      <button className={`rounded-lg w-full h-full p-2 bg-${color}-100 hover:bg-${color}-200`} onClick={onClick}>
        <div className="w-full h-full text-slate-100 font-bold text-2xl">
          {children && children}
        </div>
      </button>
    </div>
  );
};
