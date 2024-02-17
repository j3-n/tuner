import { ReactNode } from "react";

type Props = {
  children?: ReactNode;
}

export const ButtonComponent: React.FC<Props> = ({ children }) => {
  return (
    <div className="rounded-lg p-2 border-2 border-teal-200 bg-teal-100 hover:bg-teal-200">
      <button>
        {children && children}
      </button>
    </div>
  );
};
