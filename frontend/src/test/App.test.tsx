import React from 'react';
import { render, screen, waitFor, fireEvent } from '@testing-library/react';
import App from '../App';
import { getStoreData, addData, initDB } from '../lib/db';

// Mock the functions from the db module
jest.mock('./lib/db', () => ({
  getStoreData: jest.fn(),
  addData: jest.fn(),
  initDB: jest.fn(),
}));

describe('App', () => {
  beforeEach(() => {
    // Reset the mock implementation before each test
    // @ts-ignore
    getStoreData.mockReset();
    // @ts-ignore
    addData.mockReset();
    // @ts-ignore
    initDB.mockImplementation(() => Promise.resolve({} as IDBDatabase));
  });

  
  it('should render without errors', () => {
    render(<App />);
    // Assert that the component renders without throwing any errors
  });

  it('should display data from IndexedDB', async () => {
    // Mock the data from IndexedDB
    const mockData = [
      { id: 1, fact: 'Fact 1', length: 10 },
      { id: 2, fact: 'Fact 2', length: 15 },
    ];
    // @ts-ignore
    getStoreData.mockResolvedValue(mockData);

    render(<App />);

    // Wait for the component to update with the data from IndexedDB
    await waitFor(() => {
      // Assert that the rendered data matches the mock data
      expect(screen.getByText('Fact 1')).toBeInTheDocument();
      // eslint-disable-next-line testing-library/no-wait-for-multiple-assertions
      expect(screen.getByText('Fact 2')).toBeInTheDocument();
    });
  });

  it('should fetch data from the Cloud Database and update displayed data', async () => {
    // Mock the response from the Cloud Database
    const mockData = [
      { id: 3, fact: 'Fact 3', length: 12 },
      { id: 4, fact: 'Fact 4', length: 20 },
    ];
    global.fetch = jest.fn().mockResolvedValueOnce({
      json: () => Promise.resolve(mockData),
    }) as jest.MockedFunction<typeof global.fetch>;

    render(<App />);

    // Wait for the component to update with the fetched data
    await waitFor(() => {
      // Assert that the rendered data matches the fetched data
      expect(screen.getByText('Fact 3')).toBeInTheDocument();
      // eslint-disable-next-line testing-library/no-wait-for-multiple-assertions
      expect(screen.getByText('Fact 4')).toBeInTheDocument();
    });

    // Assert that the saveToIndexedDB function was called with the fetched data
    expect(addData).toHaveBeenCalledWith('facts', mockData);
  });

  it('should save data to IndexedDB correctly', async () => {
    render(<App />);

    // Simulate the saveToIndexedDB function being called
    const mockData = [
      { id: 5, fact: 'Fact 5', length: 8 },
      { id: 6, fact: 'Fact 6', length: 11 },
    ];
    // @ts-ignore
    addData.mockImplementationOnce(() => Promise.resolve(null));
    fireEvent.click(screen.getByText('Get New Cat Fact'));

    // Wait for the saveToIndexedDB function to be called
    await waitFor(() => {
      // Assert that the saveToIndexedDB function was called with the correct data
      expect(addData).toHaveBeenCalledWith('facts', mockData);
    });
  });

  it('should update displayed data when "Get New Cat Fact" button is clicked', async () => {
    // Mock the response from the Cloud Database
    const mockData = [
      { id: 7, fact: 'Fact 7', length: 9 },
      { id: 8, fact: 'Fact 8', length: 14 },
    ];
    global.fetch = jest.fn().mockResolvedValueOnce({
      json: () => Promise.resolve(mockData),
    }) as jest.MockedFunction<typeof global.fetch>;

    render(<App />);

    // Click the "Get New Cat Fact" button
    fireEvent.click(screen.getByText('Get New Cat Fact'));

    // Wait for the component to update with the fetched data
    await waitFor(() => {
      // Assert that the rendered data matches the fetched data
      expect(screen.getByText('Fact 7')).toBeInTheDocument();
      // eslint-disable-next-line testing-library/no-wait-for-multiple-assertions
      expect(screen.getByText('Fact 8')).toBeInTheDocument();
    });

    // Assert that the saveToIndexedDB function was called with the fetched data
    expect(addData).toHaveBeenCalledWith('facts', mockData);
  });
});
