CREATE TABLE IF NOT EXISTS checkout_promotions (
    promotion_id UUID PRIMARY KEY,
    session_id VARCHAR(255) NOT NULL,
    user_id UUID,
    device_id VARCHAR(255),
    trigger_type VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    started_at TIMESTAMP WITH TIME ZONE NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    completed_at TIMESTAMP WITH TIME ZONE,
    voucher_id UUID,
    voucher_status VARCHAR(50),
    campaign_code VARCHAR(255) NOT NULL DEFAULT 'ACCESSORY_NEXT_PURCHASE_15',
    reward_type VARCHAR(50) NOT NULL DEFAULT 'percentage',
    reward_value INT NOT NULL DEFAULT 15,
    reward_scope VARCHAR(255) NOT NULL DEFAULT 'accessories_next_purchase',
    cooldown_until TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_checkout_promotions_session_id ON checkout_promotions(session_id);
CREATE INDEX IF NOT EXISTS idx_checkout_promotions_user_id ON checkout_promotions(user_id);
CREATE INDEX IF NOT EXISTS idx_checkout_promotions_device_id ON checkout_promotions(device_id);
