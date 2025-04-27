import React, { useState } from 'react';
import axios from 'axios';

const FileUpload = () => {
    const [selectedFile, setSelectedFile] = useState(null);
    const [uploadStatus, setUploadStatus] = useState('');

    const handleFileChange = (e) => {
        setSelectedFile(e.target.files[0]);
    };

    const handleUpload = async () => {
        if (!selectedFile) {
            setUploadStatus('Please select a file first!');
            return;
        }

        const formData = new FormData();
        formData.append('file', selectedFile);

        try {
            const response = await axios.post('http://localhost:8181/upload', formData, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                },
            });
            setUploadStatus(response.data.message);
        } catch (error) {
            setUploadStatus('Upload failed: ' + (error.response?.data?.error || error.message));
        }
    };

    return (
        <div className="flex flex-col items-center justify-center my-10 space-y-6">
            <div className="flex flex-col items-center space-y-4 p-6 bg-white rounded-xl shadow-lg">
                <label className="text-lg font-semibold text-gray-700">Select a file to upload</label>
                <input
                    type="file"
                    onChange={handleFileChange}
                    className="text-gray-500 file:mr-4 file:py-2 file:px-6 file:rounded-full file:border-0 file:text-sm file:font-semibold file:bg-purple-50 file:text-purple-700 hover:file:bg-purple-100"
                />
                <button
                    onClick={handleUpload}
                    className="bg-purple-500 hover:bg-purple-600 text-white px-8 py-2 rounded-full transition-all duration-300"
                >
                    Upload File
                </button>
            </div>

            {uploadStatus && (
                <p className="text-sm text-gray-700 mt-4">{uploadStatus}</p>
            )}
        </div>
    );
};

export default FileUpload;
