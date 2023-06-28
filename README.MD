# Project Overview

This project is a Chrome extension developed on top of Google Chrome. It allows users to fetch cat facts from an external API, persist the data to a chosen database, and display the facts on a webpage. The extension utilizes both IndexedDB for local caching and an external Cloud Database for persistent storage.

## Objectives

The main objectives of this project are as follows:

- Make a Chrome extension on top of Google Chrome that makes API calls to an external source and then persists the fetched data to a chosen database. Webhooks are used to send the data to the database.
- Modify our existing Chrome Extension to fetch data from an external free API, save this data to both IndexedDB (for local caching) and an external Cloud Database (for persistent storage), and display a webpage that retrieves data first from IndexedDB for faster loading and subsequently fetches the remainder of the data from our Cloud Database.

## Project Directory

The project directory structure is as follows:
```            
backend
frontend
```

The `backend` directory contains the code for the backend Go server that interacts with the chosen database and handles the API calls. The `frontend` directory contains the code for the frontend React Chrome extension that fetches and displays the cat facts using IndexedDB and the Cloud Database.
Each directory may have its own subdirectories and files depending on the project requirements and architecture.
Please refer to the respective directories for detailed information on the code structure and functionality.

## Running the Project

To run the project, both the frontend and backend components need to be executed.

1. Start the Backend:
   - Navigate to the `backend` directory.
   - Follow the instructions provided in the backend's README to set up and run the backend Go server.

2. Start the Frontend:
   - Navigate to the `frontend` directory.
   - Follow the instructions provided in the frontend's README to set up and run the React Chrome extension.

Ensure that both the frontend and backend are running simultaneously to ensure proper functionality of the project.
