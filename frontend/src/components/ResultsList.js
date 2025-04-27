import React from 'react';

const ResultsList = ({ results }) => {
    if (!results || results === null || results.length === 0) {
        return <div className="text-center text-xl font-semibold">No results found.</div>;
    }

    return (
        <div className="overflow-x-auto max-w-full">
            <div className="bg-white shadow-md rounded-lg overflow-hidden">
                <table className="min-w-full table-auto border-collapse text-sm">
                    <thead className="bg-purple-200">
                        <tr>
                            <th className="border-b px-4 py-2 text-left font-semibold">Message</th>
                            <th className="border-b px-4 py-2 text-left font-semibold">Sender</th>
                            <th className="border-b px-4 py-2 text-left font-semibold">namespace</th>
                            <th className="border-b px-4 py-2 text-left font-semibold">Timestamp</th>
                        </tr>
                    </thead>
                    <tbody>
                        {results.map((result, index) => (
                            <tr key={index} className="hover:bg-gray-100">
                                <td className="border-b px-4 py-3">{result.message}</td>
                                <td className="border-b px-4 py-3">{result.sender}</td>
                                <td className="border-b px-4 py-3">{result.namespace}</td>
                                <td className="border-b px-4 py-3">{result.nano_timestamp}</td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>
        </div>
    );
};

export default ResultsList;
