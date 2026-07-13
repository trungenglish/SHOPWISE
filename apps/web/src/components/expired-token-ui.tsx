import React, { useEffect, useRef } from 'react';
import { Button } from '@shopwise/ui/components/button';
import { Card, CardFooter } from '@shopwise/ui/components/card';
import { AlertCircle, Clock, RefreshCw, XCircle, ShieldAlert } from 'lucide-react';
import { useNavigate } from '@tanstack/react-router';

export type TokenErrorCode = 
  | 'token_missing' 
  | 'token_expired' 
  | 'token_consumed' 
  | 'token_revoked' 
  | 'token_invalid' 
  | 'network_error';

interface ExpiredTokenUIProps {
  errorCode: TokenErrorCode;
  onRetry?: () => void;
  onRequestNewLink?: () => void;
}

export function ExpiredTokenUI({ errorCode, onRetry, onRequestNewLink }: ExpiredTokenUIProps) {
  const navigate = useNavigate();
  const headerRef = useRef<HTMLHeadingElement>(null);

  useEffect(() => {
    // Focus management for accessibility
    if (headerRef.current) {
      headerRef.current.focus();
    }
  }, [errorCode]);

  const getStateDetails = () => {
    switch (errorCode) {
      case 'token_expired':
        return {
          title: 'Link Expired',
          description: 'This resume link has expired because it is older than 24 hours. Please request a new link to continue your session.',
          icon: <Clock className="w-12 h-12 text-muted-foreground mb-4" aria-hidden="true" />,
          action: 'request_new'
        };
      case 'token_consumed':
        return {
          title: 'Link Already Used',
          description: 'This resume link has already been used. Resume links are single-use for your security. Please request a new link.',
          icon: <RefreshCw className="w-12 h-12 text-muted-foreground mb-4" aria-hidden="true" />,
          action: 'request_new'
        };
      case 'token_revoked':
        return {
          title: 'Link Revoked',
          description: 'This link was revoked because a newer resume link was generated for this session. Please use the most recent link.',
          icon: <ShieldAlert className="w-12 h-12 text-destructive mb-4" aria-hidden="true" />,
          action: 'request_new'
        };
      case 'token_invalid':
      case 'token_missing':
        return {
          title: 'Invalid Link',
          description: 'The provided resume link is invalid or malformed. Please ensure you clicked the full link from your message.',
          icon: <XCircle className="w-12 h-12 text-destructive mb-4" aria-hidden="true" />,
          action: 'home'
        };
      case 'network_error':
      default:
        return {
          title: 'Connection Error',
          description: 'We encountered a temporary network issue while verifying your link. Please try again.',
          icon: <AlertCircle className="w-12 h-12 text-destructive mb-4" aria-hidden="true" />,
          action: 'retry'
        };
    }
  };

  const details = getStateDetails();

  return (
    <div className="flex items-center justify-center min-h-screen bg-background p-4">
      <Card className="w-full max-w-md shadow-lg border-muted">
        <div className="flex flex-col items-center pt-6 px-6 pb-2 text-center">
          {details.icon}
          <h1 
            ref={headerRef} 
            tabIndex={-1} 
            className="text-2xl font-bold tracking-tight focus:outline-none"
            aria-live="polite"
          >
            {details.title}
          </h1>
          <p className="text-muted-foreground mt-2" aria-live="polite">
            {details.description}
          </p>
        </div>
        <CardFooter className="flex flex-col gap-3 mt-6">
          {details.action === 'request_new' && onRequestNewLink && (
            <Button className="w-full" onClick={onRequestNewLink} autoFocus>
              Request New Link
            </Button>
          )}
          {details.action === 'retry' && onRetry && (
            <Button className="w-full" onClick={onRetry} autoFocus>
              Try Again
            </Button>
          )}
          <Button variant="outline" className="w-full" onClick={() => navigate({ to: '/' })}>
            Return to Home
          </Button>
        </CardFooter>
      </Card>
    </div>
  );
}
