import React from 'react';

export const CustomerInformationForm: React.FC<{
    onSubmit: (info: any) => void;
}> = ({ onSubmit }) => {
    return (
        <form className="customer-info-form" onSubmit={(e) => {
            e.preventDefault();
            onSubmit({});
        }}>
            <h3>Customer Information</h3>
            <input type="text" placeholder="Name" />
            <input type="text" placeholder="Phone Number" />
            <button type="submit">Submit</button>
        </form>
    );
};
