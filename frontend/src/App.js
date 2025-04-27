import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import SearchPage from './pages/SearchPage';
import UploadPage from './pages/UploadPage';

const App = () => {
  return (
    <Router>
      <div>
        {/* Navbar */}
        <nav className="bg-purple-500 p-4 shadow-lg">
          <div className="container mx-auto flex justify-between items-center">
            <div className="text-white text-2xl font-bold">MyApp</div>
            <div className="space-x-6">
              <Link to="/" className="text-white hover:underline">Search</Link>
              <Link to="/upload" className="text-white hover:underline">Upload</Link>
            </div>
          </div>
        </nav>

        {/* Pages */}
        <div className="container mx-auto p-6">
          <Routes>
            <Route path="/" element={<SearchPage />} />
            <Route path="/upload" element={<UploadPage />} />
          </Routes>
        </div>
      </div>
    </Router>
  );
};

export default App;
