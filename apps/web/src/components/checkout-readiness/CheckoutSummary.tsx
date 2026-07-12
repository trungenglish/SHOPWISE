import React from 'react';

export const CheckoutSummary: React.FC<{
    productName: string;
    price: number;
    inventoryMessage: string;
}> = ({ productName, price, inventoryMessage }) => {
    return (
        <div className="checkout-summary">
            <h3>{productName}</h3>
            <p>Price: {price} VND</p>
            <p>Status: {inventoryMessage}</p>
        </div>
    );
};
