export type TriggerType = 'saved_product' | 'started_checkout';

export type PromotionStatus =
  | 'ACTIVE'
  | 'EXPIRED'
  | 'COMPLETED_PENDING_VOUCHER'
  | 'VOUCHER_ISSUED'
  | 'ERROR_RECOVERABLE'
  | 'ISSUE_RETRY_REQUIRED';

export interface StartPromotionRequest {
  trigger_type: TriggerType;
}

export interface PromotionStatusResponse {
  promotion_id?: string;
  status: PromotionStatus;
  expires_at?: string;
  campaign_code?: string;
  reward_type?: string;
  reward_value?: number;
  reward_scope?: string;
}
