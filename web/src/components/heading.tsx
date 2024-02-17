import { ReactNode } from "react";

type Props = {
  children?: ReactNode;
}

export const H1Component: React.FC<Props> = ({
  children,
}): JSX.Element => {
  return (
    <h1 className="font-bold">
      {children && children}
    </h1>
  );
};

export const H2Component: React.FC<Props> = ({
  children,
}): JSX.Element => {
  return (
    <h2 className="font-bold">
      {children && children}
    </h2>
  );
};

export const H3Component: React.FC<Props> = ({
  children,
}): JSX.Element => {
  return (
    <h3 className="font-bold">
      {children && children}
    </h3>
  );
};
