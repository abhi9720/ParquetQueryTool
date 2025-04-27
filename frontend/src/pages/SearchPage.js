import React, { useState } from 'react';
import axios from 'axios';
import SearchBar from '../components/SearchBar';
import ResultsList from '../components/ResultsList';

const SearchPage = () => {
    const [results, setResults] = useState(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
    const [metadata, setMetadata] = useState({ matches: 0, timeTaken: 0 });

    const handleSearch = async (query) => {
        setLoading(true);
        setError(null);
        try {
            const response = await axios.get(`http://localhost:8181/search?query=${query}`);
            setResults(response.data.results);
            setMetadata({
                matches: response.data.count,
                timeTaken: response.data.search_time_ms,
            });
        } catch (error) {
            setError('Error searching, please try again.');
            setResults([]);
            setMetadata({ matches: 0, timeTaken: 0 });
            console.error('Error searching:', error);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="container mx-auto">
            <SearchBar onSearch={handleSearch} />

            {loading && <div className="text-center">Searching...</div>}
            {error && <div className="text-red-500 text-center">{error}</div>}

            {!loading && (
                <div>
                    <div className="my-4">
                        <span className="text-md">Showing : {results?.length || 0}/{metadata.matches} </span>
                        <span className="ml-5 text-md">Time taken: {metadata.timeTaken} ms</span>
                    </div>
                    <ResultsList results={results} />
                </div>
            )}
        </div>
    );
};

export default SearchPage;
