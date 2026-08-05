import React from 'react';
import type {
  ErrorStateProps,
  RetryActionProps,
  ContinueCachedDataProps,
  ModifySearchProps,
  AlternativeRecommendationProps
} from '@shopwise/ui-protocol/types/error';
import { dynamicUIEventDispatcher } from '../EventDispatcher';

export const ErrorState: React.FC<ErrorStateProps> = ({ title, message, code, severity = 'error' }) => (
  <div className={`error-state severity-${severity}`}>
    <h3>{title}</h3>
    <p>{message}</p>
    {code && <small>Error Code: {code}</small>}
  </div>
);

export const RetryAction: React.FC<RetryActionProps & { id: string }> = ({ id, label = 'Retry', actionId }) => (
  <button
    className="btn btn-primary retry-btn"
    onClick={() => dynamicUIEventDispatcher.dispatch(id, 'RETRY', { actionId })}
  >
    {label}
  </button>
);

export const ContinueCachedData: React.FC<ContinueCachedDataProps & { id: string }> = ({ id, label = 'Continue Offline' }) => (
  <button
    className="btn btn-secondary continue-cached-btn"
    onClick={() => dynamicUIEventDispatcher.dispatch(id, 'CONTINUE_CACHED', {})}
  >
    {label}
  </button>
);

export const ModifySearch: React.FC<ModifySearchProps & { id: string }> = ({ id, label = 'Modify Search', originalQuery }) => (
  <button
    className="btn btn-outline modify-search-btn"
    onClick={() => dynamicUIEventDispatcher.dispatch(id, 'MODIFY_SEARCH', { originalQuery })}
  >
    {label}
  </button>
);

export const AlternativeRecommendation: React.FC<AlternativeRecommendationProps & { id: string }> = ({ id, message, productIds }) => (
  <div className="alternative-recommendation">
    {message && <p>{message}</p>}
    <button
      className="btn btn-secondary"
      onClick={() => dynamicUIEventDispatcher.dispatch(id, 'VIEW_ALTERNATIVES', { productIds })}
    >
      View Alternatives
    </button>
  </div>
);
