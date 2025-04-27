# Apica Search - Frontend

This is the frontend part of the Apica Search app. It allows users to upload files and search through them. The app is built using React and styled with Tailwind CSS.

## Features:
- **File Upload**: Allows users to upload Parquet files to the server for indexing.
- **Search**: Users can search for records by entering a search query.
- **Results**: Displays search results in a table format with relevant details like message, sender, and timestamp.

## How It Works:

1. **File Upload**:
   - The user selects a file and uploads it to the server. The status of the upload is displayed to the user.

2. **Search**:
   - The user enters a query in the search bar and the app fetches matching results from the backend.
   - The results are displayed in a table with relevant data (message, sender, namespace, timestamp).

## How to Run the App

### Prerequisites:
- Node.js (>=14.x)
- Yarn (optional, can also use npm)

### Steps to Run:

1. Clone the repository:
   ```bash
   git clone <repository_url>
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   # or if you're using yarn:
   yarn install
   ```

3. Start the development server:
   ```bash
   npm start
   # or if you're using yarn:
   yarn start
   ```

   The app will be available at `http://localhost:3000`.

4. Make sure the backend is running on `http://localhost:8181` to handle file uploads and search queries.

## File Structure:

- **src/components**: Contains reusable components like `FileUpload`, `SearchBar`, and `ResultsList`.
- **src/pages**: Contains page components for the Upload and Search pages.
- **src/App.js**: Main application file that sets up routing and displays pages.

## Features

### File Upload Page:
- Users can upload a Parquet file which is sent to the backend server.
- Displays upload status messages (success or failure).

### Search Page:
- Users can enter a search query and view the results in a table format.
- Shows metadata like the number of matches and the time taken for the search.

## Technology Stack:

- **React**: For building the user interface.
- **Tailwind CSS**: For styling.
- **Axios**: For making HTTP requests to the backend.
