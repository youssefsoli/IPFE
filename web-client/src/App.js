import icon from './images/icon.png';
import title from './images/title.svg';
import ipfe from './images/ipfe.svg';
import { SearchBar } from './components/SearchBar';
import './App.css';
import { useEffect, useState } from 'react';

export const App = () => {
  const [query, setQuery] = useState('');

  const handleSubmit = (newQuery) => {
    setQuery(newQuery);
  };

  useEffect(() => {}, [query]);

  return (
    <main className="bg-primary h-full w-full">
      {!query ? (
        <>
          <div className="flex h-1/2">
            <div className="m-auto mb-[40px]">
              <img className="mx-auto" src={icon} width="160px" />
              <img src={title} />
            </div>
          </div>
          <div className="flex h-1/2">
            <div className="mx-auto">
              <SearchBar handleSubmit={handleSubmit} />
            </div>
          </div>
        </>
      ) : (
        <div>
          <div>
            <img className="" src={icon} width="160px" />
            <img className="" src={ipfe} />
          </div>
          <SearchBar handleSubmit={handleSubmit} />
        </div>
      )}
    </main>
  );
};

export default App;
