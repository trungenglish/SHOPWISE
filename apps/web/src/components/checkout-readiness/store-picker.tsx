import React, { useState } from 'react';

export const StorePicker: React.FC<{
    onSelectStore: (storeId: string) => void;
}> = ({ onSelectStore }) => {
    const [useGeolocation, setUseGeolocation] = useState(false);

    return (
        <div className="store-picker">
            <h3>Select a Store</h3>
            <button onClick={() => setUseGeolocation(true)}>
                {useGeolocation ? "Using Location..." : "Use My Location"}
            </button>
            <ul>
                <li onClick={() => onSelectStore("store_1")}>Store 1 (District 1)</li>
                <li onClick={() => onSelectStore("store_2")}>Store 2 (District 3)</li>
            </ul>
        </div>
    );
};
