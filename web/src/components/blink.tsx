import { Link } from "@tanstack/react-router";
import { ReactNode } from "react";

type Props = {
  to: string;
  children?: ReactNode;
}

export const Blink: React.FC<Props> = ({
  to,
  children
}) => {
  return (
    <Link className="" to={to}>
      <div className="w-full h-full text-center justify-center items-center flex rounded-lg p-10">
          {children && children}
      </div>
    </Link>
  );
};
