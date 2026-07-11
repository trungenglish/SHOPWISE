import React, { useState } from 'react';
import { deletePreference } from '../../../api/decisionMemory';

interface PreferenceUndoToastProps {
  preferenceId: string;
  category: string;
  value: string;
  onDismiss?: () => void;
  onUndoSuccess?: () => void;
}

export const PreferenceUndoToast: React.FC<PreferenceUndoToastProps> = ({
  preferenceId,
  category,
  value,
  onDismiss,
  onUndoSuccess,
}) => {
  const [isUndoing, setIsUndoing] = useState(false);

  const handleUndo = async () => {
    try {
      setIsUndoing(true);
      await deletePreference(preferenceId);
      if (onUndoSuccess) onUndoSuccess();
    } catch (error) {
      console.error('Failed to undo preference', error);
    } finally {
      setIsUndoing(false);
      if (onDismiss) onDismiss();
    }
  };

  return (
    <div className="fixed bottom-4 right-4 bg-gray-800 text-white p-4 rounded shadow-lg flex items-center justify-between space-x-4">
      <div>
        <p className="text-sm font-medium">Auto-saved preference</p>
        <p className="text-xs text-gray-400">
          {category}: {value}
        </p>
      </div>
      <div className="flex space-x-2">
        <button
          onClick={handleUndo}
          disabled={isUndoing}
          className="text-xs bg-gray-700 hover:bg-gray-600 px-3 py-1 rounded"
        >
          {isUndoing ? 'Undoing...' : 'Undo'}
        </button>
        <button
          onClick={onDismiss}
          className="text-xs text-gray-400 hover:text-gray-200 px-2 py-1"
        >
          Dismiss
        </button>
      </div>
    </div>
  );
};
