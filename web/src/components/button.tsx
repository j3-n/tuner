import { ReactNode } from "react";

type Props = {
  onClick?: () => void;
  children?: ReactNode;
}

export const ButtonComponent: React.FC<Props> = ({
  onClick,
  children
}): JSX.Element => {
  return (
    <div className="rounded-lg p-2 border-2 border-teal-200 bg-teal-100 hover:bg-teal-200">
      <button onClick={onClick}>
        {children && children}
      </button>
    </div>
  );
};
