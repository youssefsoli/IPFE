import SearchIcon from '@mui/icons-material/Search';
import { useState } from 'react';

export const SearchBar = ({ handleSubmit }) => {
  const [input, setInput] = useState('');

  return (
    <form
      className={`w-[582px] h-[44px] border border-googlegray rounded-full flex
      focus-within:bg-neutral-800 focus-within:border-neutral-800 ${
        input ? 'bg-neutral-800 border-neutral-800' : ''
      }`}
      onSubmit={(e) => {
        e.preventDefault();
        handleSubmit(input);
      }}
    >
      <SearchIcon
        className="my-auto ml-3 mr-2"
        sx={{ color: '#9AA0A6', strokeWidth: 1 }}
      />
      <input
        className="w-full mr-6 bg-transparent text-white"
        value={input}
        onChange={(e) => {
          setInput(e.target.value);
        }}
      />
    </form>
  );
};
