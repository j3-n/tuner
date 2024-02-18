type Props = {
  placeholder?: string;
  id: string;
  onChange?: React.ChangeEventHandler<HTMLInputElement>;
};


export const InputComponent: React.FC<Props> = ({
  placeholder,
  id,
  onChange
}): JSX.Element => {
  return (
    <div className="w-1/3 h-12 mx-auto">
      <input className="w-full h-full text-center border-solid border-2 border-slate-800 rounded-xl"
        placeholder={placeholder}
        id={id}
        onChange={onChange}
      />
    </div>
  );
};
