import { ReactNode } from "react";

type Props = {
  onClick?: () => void;
  color?: string;
  children?: ReactNode;
}

export const ButtonComponent: React.FC<Props> = ({
  onClick,
  color = "teal",
  children
}): JSX.Element => {
  return (
    <div>
      <button className={`rounded-lg p-2 border-2 border-${color}-200 bg-${color}-100 hover:bg-${color}-200`} onClick={onClick}>
        <div>
        {children && children}
        </div>
      </button>
    </div>
  );
};
