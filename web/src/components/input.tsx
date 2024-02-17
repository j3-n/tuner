type Props = {
  placeholder?: string;
  value?: string;
  onChange?: React.ChangeEventHandler<HTMLInputElement>;
};


export const InputComponent: React.FC<Props> = ({ placeholder, value, onChange }) => {
  return (
    <div>
      <input
        placeholder={placeholder}
        value={value}
        onChange={onChange}
      />
    </div>
  );
};
