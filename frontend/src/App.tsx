import React, { useEffect, useState } from 'react';
import { Fact, getStoreData, addData, Stores, initDB } from './lib/db';

import 'tailwindcss/tailwind.css';


const App = () => {
  const [messages, setMessages] = useState<Fact[]>([]);

  useEffect(() => {
    // Retrieve data from IndexedDB
    retrieveFromIndexedDB();

    // Fetch remaining data from Cloud Database
    fetchFromCloudDatabase();
  }, []);

  // Function to fetch data from Cloud Database
  const fetchFromCloudDatabase = async () => {
    try {
      const response = await fetch(
        'http://localhost:8000/api/facts'
      );
      const data = await response.json();
      const facts = data.map((fact: Fact) => ({
        id: fact.id,
        fact: fact.fact,
        length: fact.length,
      }));
      setMessages(facts);
      saveToIndexedDB(facts);
    } catch (error) {
      console.error('Error fetching data from Cloud Database:', error);
    }
  };

  // Function to retrieve data from IndexedDB
  const retrieveFromIndexedDB = async () => {
    const db = await initDB();
    if (db) {
      const storedData = await getStoreData<Fact>(Stores.Facts);
      setMessages(storedData);
    }
  };


  // Function to save data to IndexedDB
  const saveToIndexedDB = async (data: Fact[]) => {
    const db = await initDB();
    if (db) {
      await addData('facts', data);
    } else {
      console.error('Failed to save data to IndexedDB.');
    }
  };

  return (
    <div className="flex flex-col justify-center items-center py-8 px-4 gap-4">
      <h1 className="text-2xl font-semibold text-center">Cat Facts Library</h1>

      <div className="flex flex-col gap-2">
        {messages.map((message: Fact) => (
          <div key={message.id} className="rounded-lg shadow-md p-4 border border-gray-200">
            <p>{message.fact}</p>
          </div>
        ))}
      </div>

      <button onClick={
        () => fetchFromCloudDatabase()
      }
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded w-full"
      >Get New Cat Fact</button>
    </div>
  );
};

export default App;
