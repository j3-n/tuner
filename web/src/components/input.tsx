type Props = {
  placeholder?: string;
  id: string;
  onChange?: React.ChangeEventHandler<HTMLInputElement>;
};


export const InputComponent: React.FC<Props> = ({
  placeholder,
  id,
  onChange
}) => {
  return (
    <div>
      <input
        placeholder={placeholder}
        id={id}
        onChange={onChange}
      />
    </div>
  );
};
