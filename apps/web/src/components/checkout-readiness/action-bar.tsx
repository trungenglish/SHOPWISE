import React from 'react';

export const ActionBar: React.FC<{
    onContinue: () => void;
    onRequestSimilar: () => void;
    onChangeStore: () => void;
    onRetry: () => void;
    onSaveSession: () => void;
    onResumeSession: () => void;
}> = (props) => {
    return (
        <div className="action-bar">
            <button onClick={props.onContinue}>Continue to Checkout</button>
            <button onClick={props.onRequestSimilar}>Request Similar Products</button>
            <button onClick={props.onChangeStore}>Change Store</button>
            <button onClick={props.onRetry}>Retry Validation</button>
            <button onClick={props.onSaveSession}>Save Session</button>
            <button onClick={props.onResumeSession}>Resume Session</button>
        </div>
    );
};
