import { useState, useCallback, useRef, useEffect } from 'react';
import {
  DecisionSession,
  createSession,
  getSession,
  autoSaveSession,
} from '../api/decisionMemory';

export function useDecisionMemory() {
  const [activeSession, setActiveSession] = useState<DecisionSession | null>(null);
  const [isSaving, setIsSaving] = useState(false);
  const [saveError, setSaveError] = useState<string | null>(null);
  
  const saveTimeoutRef = useRef<NodeJS.Timeout | null>(null);

  const startNewSession = useCallback(async (initialMessage: string) => {
    try {
      const session = await createSession(initialMessage);
      setActiveSession(session);
      return session;
    } catch (err) {
      console.error(err);
      throw err;
    }
  }, []);

  const loadSession = useCallback(async (id: string) => {
    try {
      const session = await getSession(id);
      setActiveSession(session);
      return session;
    } catch (err) {
      console.error(err);
      throw err;
    }
  }, []);

  const triggerAutoSave = useCallback((pinnedProducts: string[]) => {
    if (!activeSession) return;

    if (saveTimeoutRef.current) {
      clearTimeout(saveTimeoutRef.current);
    }

    saveTimeoutRef.current = setTimeout(async () => {
      setIsSaving(true);
      setSaveError(null);
      try {
        const timestamp = new Date().toISOString();
        const updated = await autoSaveSession(activeSession.ID, timestamp, pinnedProducts);
        setActiveSession(updated);
      } catch (err: any) {
        if (err.message === 'conflict') {
          setSaveError('Session updated elsewhere. Please refresh.');
        } else {
          setSaveError('Failed to autosave.');
        }
      } finally {
        setIsSaving(false);
      }
    }, 1000); // 1s debounce
  }, [activeSession]);

  useEffect(() => {
    return () => {
      if (saveTimeoutRef.current) clearTimeout(saveTimeoutRef.current);
    };
  }, []);

  return {
    activeSession,
    isSaving,
    saveError,
    startNewSession,
    loadSession,
    triggerAutoSave,
  };
}
