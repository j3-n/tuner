import { Link } from "@tanstack/react-router";
import { ReactNode } from "react";

type Props = {
  to: string;
  size?: "sm" | "lg";
  children?: ReactNode;
}

export const LinkComponent: React.FC<Props> = ({
  to,
  size,
  children
}): JSX.Element => {
  let className = "";

  if (size === "sm") {
    className += "rounded-lg p-2 border-2 border-teal-200 bg-teal-100 hover:bg-teal-200"
  } else {
    className += "w-full h-full text-center justify-center items-center flex rounded-lg p-10"
  }

  return (
    <Link className="" to={to}>
      <div className={className}>
        {children && children}
      </div>
    </Link>
  );
};
