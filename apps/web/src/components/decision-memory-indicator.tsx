import React from 'react';
import { CheckCircle2, Loader2, AlertCircle } from 'lucide-react';

export function DecisionMemoryIndicator({ isSaving, saveError }: { isSaving: boolean, saveError: string | null }) {
  if (saveError) {
    return (
      <div className="flex items-center gap-1 text-xs text-red-500 animate-in fade-in duration-300">
        <AlertCircle className="h-3 w-3" />
        <span>{saveError}</span>
      </div>
    );
  }

  if (isSaving) {
    return (
      <div className="flex items-center gap-1 text-xs text-muted-foreground animate-in fade-in duration-300">
        <Loader2 className="h-3 w-3 animate-spin" />
        <span>Saving...</span>
      </div>
    );
  }

  return (
    <div className="flex items-center gap-1 text-xs text-muted-foreground opacity-50 hover:opacity-100 transition-opacity">
      <CheckCircle2 className="h-3 w-3 text-green-500" />
      <span>Saved to Decision Memory</span>
    </div>
  );
}
