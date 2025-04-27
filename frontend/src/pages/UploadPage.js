import React from 'react';
import FileUpload from '../components/FileUpload';

const UploadPage = () => {
    return (
        <div className="container mx-auto">
            <h1 className="text-3xl font-bold text-center my-8">Upload File</h1>
            <FileUpload />
        </div>
    );
};

export default UploadPage;
