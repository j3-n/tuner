import { Link } from "@tanstack/react-router";
import { ReactNode } from "react";

type Props = {
  to: string;
  children?: ReactNode;
}

export const LinkComponent: React.FC<Props> = ({
  to,
  children
}) => {
  return (
    <div className="rounded-lg p-2 border-2 border-teal-200 bg-teal-100 hover:bg-teal-200">
      <Link className="" to={to}>
        {children && children}
      </Link>
    </div>
  );
};
