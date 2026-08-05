import React from 'react';
import type {
  ChatMessageProps,
  ThinkingIndicatorProps,
  RecommendationExplanationProps,
  DecisionReasoningProps,
  ConfidenceIndicatorProps
} from '@shopwise/ui-protocol/types/commerce';

export const ChatMessage: React.FC<ChatMessageProps> = ({ text, sender }) => (
  <div className={`chat-message ${sender}`}>
    <p>{text}</p>
  </div>
);

export const ThinkingIndicator: React.FC<ThinkingIndicatorProps> = ({ statusText }) => (
  <div className="thinking-indicator">
    <span className="spinner"></span>
    {statusText && <span className="status">{statusText}</span>}
  </div>
);

export const RecommendationExplanation: React.FC<RecommendationExplanationProps> = ({ text, highlights }) => (
  <div className="recommendation-explanation">
    {text && <p>{text}</p>}
    {highlights && (
      <ul>
        {highlights.map((h, i) => <li key={i}>{h}</li>)}
      </ul>
    )}
  </div>
);

export const DecisionReasoning: React.FC<DecisionReasoningProps> = ({ reasoning }) => (
  <div className="decision-reasoning">
    {reasoning && <p><strong>Reasoning:</strong> {reasoning}</p>}
  </div>
);

export const ConfidenceIndicator: React.FC<ConfidenceIndicatorProps> = ({ score, label }) => (
  <div className="confidence-indicator">
    <div className="score-bar" style={{ width: `${score ?? 0}%` }}></div>
    <span className="label">{label || `Confidence: ${score}%`}</span>
  </div>
);
