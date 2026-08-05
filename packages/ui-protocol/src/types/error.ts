export interface ErrorStateProps {
  title: string;
  message: string;
  code?: string;
  severity?: "info" | "warning" | "error" | "critical";
}

export interface RetryActionProps {
  label?: string;
  actionId?: string;
}

export interface ContinueCachedDataProps {
  label?: string;
}

export interface ModifySearchProps {
  label?: string;
  originalQuery?: string;
}

export interface AlternativeRecommendationProps {
  message?: string;
  productIds?: string[];
}
