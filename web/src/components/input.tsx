type Props = {
  placeholder?: string;
};


export const InputComponent: React.FC<Props> = ({ placeholder }) => {
  return (
    <div>
      <input placeholder={placeholder} />
    </div>
  );
};
