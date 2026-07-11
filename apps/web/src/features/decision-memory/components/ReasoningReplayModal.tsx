import React from 'react';
import { SessionMessage } from '../../../api/decisionMemory';

interface ReasoningStep {
  summary: string;
  evidence?: string[];
}

interface ReasoningReplayModalProps {
  message: SessionMessage;
  onClose: () => void;
}

export const ReasoningReplayModal: React.FC<ReasoningReplayModalProps> = ({ message, onClose }) => {
  let steps: ReasoningStep[] = [];
  try {
    if (message.ReasoningGraph) {
      steps = JSON.parse(message.ReasoningGraph) as ReasoningStep[];
    }
  } catch (e) {
    console.error('Failed to parse reasoning graph', e);
  }

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50">
      <div className="bg-white rounded-lg shadow-xl w-full max-w-lg overflow-hidden">
        <div className="px-6 py-4 border-b flex justify-between items-center bg-gray-50">
          <h3 className="text-lg font-semibold text-gray-800">Reasoning Replay</h3>
          <button onClick={onClose} className="text-gray-500 hover:text-gray-700">
            &times;
          </button>
        </div>
        <div className="p-6 max-h-96 overflow-y-auto">
          {steps.length === 0 ? (
            <p className="text-gray-500 italic">No reasoning steps available for this message.</p>
          ) : (
            <div className="space-y-6">
              {steps.map((step, index) => (
                <div key={index} className="flex space-x-4">
                  <div className="flex flex-col items-center">
                    <div className="flex items-center justify-center w-8 h-8 bg-blue-100 text-blue-600 rounded-full font-bold">
                      {index + 1}
                    </div>
                    {index < steps.length - 1 && (
                      <div className="w-px h-full bg-gray-200 my-2"></div>
                    )}
                  </div>
                  <div className="flex-1 pt-1">
                    <p className="text-gray-800 font-medium">{step.summary}</p>
                    {step.evidence && step.evidence.length > 0 && (
                      <ul className="mt-2 list-disc list-inside text-sm text-gray-600 space-y-1">
                        {step.evidence.map((ev, i) => (
                          <li key={i}>{ev}</li>
                        ))}
                      </ul>
                    )}
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>
        <div className="px-6 py-4 border-t bg-gray-50 text-right">
          <button
            onClick={onClose}
            className="px-4 py-2 bg-gray-200 hover:bg-gray-300 text-gray-800 rounded text-sm font-medium transition-colors"
          >
            Close
          </button>
        </div>
      </div>
    </div>
  );
};
