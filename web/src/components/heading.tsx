import { ReactNode } from "react";

type Props = {
  children?: ReactNode;
}

export const H1Component: React.FC<Props> = ({
  children,
}): JSX.Element => {
  return (
    <h1 className="font-bold text-3xl">
      {children && children}
    </h1>
  );
};

export const H2Component: React.FC<Props> = ({
  children,
}): JSX.Element => {
  return (
    <h2 className="font-bold text-2xl">
      {children && children}
    </h2>
  );
};

export const H3Component: React.FC<Props> = ({
  children,
}): JSX.Element => {
  return (
    <h3 className="font-bold text-xl">
      {children && children}
    </h3>
  );
};
