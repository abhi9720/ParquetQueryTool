import React, { useState } from 'react';

const SearchBar = ({ onSearch }) => {
    const [query, setQuery] = useState('');

    const handleChange = (e) => {
        setQuery(e.target.value);
    };

    const handleSearch = () => {
        if (query.trim()) {
            onSearch(query);
        }
    };

    return (
        <div className="h-64 flex items-center justify-center flex-col bg-purple-100">
            <h1 className="text-3xl font-bold text-center mb-6">Search Engine</h1>
            <div className="flex items-center space-x-4 bg-white p-6 rounded-2xl shadow-lg">
                <input
                    type="text"
                    value={query}
                    onChange={handleChange}
                    placeholder="Enter search keyword..."
                    className="border border-gray-300 rounded-xl px-4 py-2 w-80 focus:outline-none focus:ring-2 focus:ring-purple-400"
                />
                <button
                    onClick={handleSearch}
                    className="bg-purple-500 hover:bg-purple-600 text-white px-5 py-2 rounded-xl transition-all duration-300"
                >
                    Search
                </button>
            </div>
        </div>
    );
};

export default SearchBar;
