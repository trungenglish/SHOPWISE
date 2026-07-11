import { createFileRoute, useNavigate } from '@tanstack/react-router';
import React, { useEffect, useState } from 'react';
import { useQuery, useQueryClient } from '@tanstack/react-query';
import { Loader2 } from 'lucide-react';
import { resumeSessionFromToken, ResumeSessionError } from '@/api/decisionMemory';
import { ExpiredTokenUI, TokenErrorCode } from '@/components/ExpiredTokenUI';

export const Route = createFileRoute('/resume')({
  component: ResumePage,
  validateSearch: (search: Record<string, unknown>) => {
    return {
      token: (search.token as string) || undefined,
    };
  },
});

export function ResumePage() {
  const { token } = Route.useSearch();
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const [manualRetry, setManualRetry] = useState(0);

  const { data: session, error, isFetching } = useQuery({
    queryKey: ['resumeSession', token, manualRetry],
    queryFn: () => resumeSessionFromToken(token as string),
    enabled: !!token,
    retry: false,
    refetchOnWindowFocus: false,
  });

  useEffect(() => {
    if (session) {
      // Remove token from URL
      navigate({ to: '/resume', replace: true });
      
      // Refresh session queries
      queryClient.invalidateQueries({ queryKey: ['sessions'] });
      queryClient.invalidateQueries({ queryKey: ['session', session.ID] });

      // Navigate to dashboard
      navigate({ to: '/dashboard', search: { sessionId: session.ID } });
    }
  }, [session, navigate, queryClient]);

  const handleManualRetry = () => {
    setManualRetry(prev => prev + 1);
  };

  const handleRequestNewLink = () => {
    navigate({ to: '/' });
  };

  if (!token) {
    return <ExpiredTokenUI errorCode="token_missing" />;
  }

  if (error) {
    let errorCode: TokenErrorCode = 'network_error';
    if (error instanceof ResumeSessionError) {
      errorCode = error.code as TokenErrorCode;
    }
    return (
      <ExpiredTokenUI 
        errorCode={errorCode} 
        onRetry={errorCode === 'network_error' ? handleManualRetry : undefined}
        onRequestNewLink={
          ['token_expired', 'token_consumed', 'token_revoked'].includes(errorCode)
            ? handleRequestNewLink
            : undefined
        }
      />
    );
  }

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-background">
      <Loader2 className="w-12 h-12 animate-spin text-primary mb-4" />
      <h2 className="text-xl font-semibold" aria-live="polite">Resuming your session...</h2>
      <p className="text-muted-foreground mt-2">Please wait while we securely restore your conversation.</p>
    </div>
  );
}
