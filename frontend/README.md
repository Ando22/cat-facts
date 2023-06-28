# React Chrome Extension - Manifest V3

## Project Overview
This project is a Chrome Extension developed using React and Manifest V3. The purpose of this extension is to provide users with cat facts fetched from a free open-source API (https://catfact.ninja/). The extension utilizes IndexedDB for storing and retrieving data locally.

## Built With
- [ReactJS](https://reactjs.org/): A JavaScript library for building user interfaces.
- [Tailwind CSS](https://tailwindcss.com/): A utility-first CSS framework for styling the components.
- [IndexedDB](https://developer.mozilla.org/en-US/docs/Web/API/IndexedDB_API): A browser-based database for storing structured data locally.
- [Webpack](https://webpack.js.org/): A module bundler used for building the extension.


## Getting Started
To run the project locally using webpack and add the `dist` folder to Chrome Extension, follow these steps:

1. Install the dependencies:
```
npm install
```

2. Run the project in watch mode:
```
npm run watch
```

3. To build the project locally:
```
npm run build
```

4. Add the `dist` folder to Chrome Extension:
- Open Chrome and go to `chrome://extensions/`
- Turn on developer mode
- Click on "Load unpacked"
- Select the `dist` folder

## IndexedDB Functionality
The extension utilizes IndexedDB for storing and retrieving data locally. Here are some functions that interact with IndexedDB:

```javascript
// Initialize the IndexedDB
export const initDB = (): Promise<boolean | IDBDatabase> => {
  return new Promise((resolve) => {
    if (db) {
      // Database already initialized, return the existing instance
      resolve(db);
    } else {
      const request = indexedDB.open('facts');

      request.onupgradeneeded = () => {
        db = request.result;

        if (!db.objectStoreNames.contains(Stores.Facts)) {
          console.log('Creating facts store');
          db.createObjectStore(Stores.Facts, { keyPath: 'id' });
        }
      };

      request.onsuccess = (e) => {
        db = request.result;
        version = db.version;
        resolve(request.result);
      };

      request.onerror = (e) => {
        resolve(false);
      };
    }
  });
};

// Add data to the IndexedDB
export const addData = (storeName: string, data: Fact[]): Promise<string | null> => {
  return new Promise((resolve) => {
    if (!db) {
      console.error('IndexedDB database is not initialized. Call initDB() first.');
      resolve(null);
      return;
    }

    const tx = db.transaction(storeName, 'readwrite');
    const store = tx.objectStore(storeName);

    // Clear the object store
    const clearRequest = store.clear();

    clearRequest.onsuccess = () => {
      // Add new data to the object store
      data.forEach((fact) => {
        store.add(fact);
      });

      resolve(null);
    };

    clearRequest.onerror = () => {
      const error = clearRequest.error?.message;
      if (error) {
        resolve(error);
      } else {
        resolve('Unknown error');
      }
    };
  });
};

// Retrieve data from the IndexedDB
export const getStoreData = <T>(storeName: Stores): Promise<T[]> => {
  return new Promise((resolve) => {
    if (!db) {
      console.error('IndexedDB database is not initialized. Call initDB() first.');
      resolve([]);
      return;
    }

    const tx = db.transaction(storeName, 'readonly');
    const store = tx.objectStore(storeName);
    const request = store.getAll();

    request.onsuccess = () => {
      resolve(request.result);
    };

    request.onerror = () => {
      console.error('Error retrieving data from IndexedDB:', request.error);
      resolve([]);
    };
  });
};
```

## API Documentation and Integrations
The extension integrates with the free open-source API, Cat Fact API. The API provides cat facts that are fetched and displayed in the extension.

To fetch data from the API, make a GET request to the following endpoint:

Endpoint: `/api/facts`

Example response:
```json
[
  {
    "id": 1,
    "fact": "Cats sleep for 70% of their lives.",
    "length": 32
  },
  {
    "id": 2,
    "fact": "Cats have five toes on their front paws and four toes on their back paws.",
    "length": 70
  },
  ...
]
```


## Integration with the Backend
To integrate the frontend with the backend and handle data from the backend, you can follow the below approach:

Use the useEffect hook to retrieve data from IndexedDB and fetch remaining data from the backend Cloud Database.
Fetch data from the Cloud Database using the appropriate endpoint (/api/facts) and save it to IndexedDB for future use.
Utilize the following functions to interact with IndexedDB and Cloud Database:

```javascript
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
```