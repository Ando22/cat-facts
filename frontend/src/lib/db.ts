let db: IDBDatabase | null = null;
let version = 1;

export interface Fact {
  id: number;
  fact: string;
  length: number;
}

export enum Stores {
  Facts = 'facts',
}

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

