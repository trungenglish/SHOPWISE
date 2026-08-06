import React, { useState } from 'react';
import { branchSession, DecisionSession } from '../../../api/decision-memory';

interface BranchSessionButtonProps {
  sessionId: string;
  currentTitle: string;
  onBranchSuccess: (newSession: DecisionSession) => void;
}

export const BranchSessionButton: React.FC<BranchSessionButtonProps> = ({
  sessionId,
  currentTitle,
  onBranchSuccess,
}) => {
  const [isBranching, setIsBranching] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleBranch = async () => {
    setIsBranching(true);
    setError(null);
    try {
      const newTitle = `${currentTitle} (Branch)`;
      const newSession = await branchSession(sessionId, newTitle);
      onBranchSuccess(newSession);
    } catch (err) {
      console.error('Failed to branch session', err);
      setError('Failed to branch session.');
    } finally {
      setIsBranching(false);
    }
  };

  return (
    <div className="flex items-center">
      <button
        onClick={handleBranch}
        disabled={isBranching}
        className="text-sm px-3 py-1 bg-indigo-100 hover:bg-indigo-200 text-indigo-700 rounded transition-colors disabled:opacity-50"
        title="Create an independent copy of this session"
      >
        {isBranching ? 'Branching...' : 'Branch'}
      </button>
      {error && <span className="ml-2 text-xs text-red-500">{error}</span>}
    </div>
  );
};
