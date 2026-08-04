import React from 'react';

export const ErrorStateUI: React.FC<{ message: string }> = ({ message }) => {
    return (
        <div className="error-state-ui">
            <h3>Validation Error</h3>
            <p>{message}</p>
        </div>
    );
};
