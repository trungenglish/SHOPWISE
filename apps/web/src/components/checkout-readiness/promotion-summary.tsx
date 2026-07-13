import React from 'react';

export const PromotionSummary: React.FC<{
    expiresInMinutes: number;
    promotionCode: string;
}> = ({ expiresInMinutes, promotionCode }) => {
    return (
        <div className="promotion-summary">
            <h3>Special Promotion: {promotionCode}</h3>
            <p className="countdown">
                Expires in: <strong>{expiresInMinutes} minutes</strong>
            </p>
        </div>
    );
};
