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
    <div>
      <Link className="" to={to}>
        {children && children}
      </Link>
    </div>
  );
};
