import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import { describe, it, expect, vi } from 'vitest';
import { ExpiredTokenUI } from './ExpiredTokenUI';

vi.mock('@tanstack/react-router', () => ({
  useNavigate: () => vi.fn(),
}));

describe('ExpiredTokenUI', () => {
  it('renders expired token state and handles focus', () => {
    const handleRequestNew = vi.fn();
    render(<ExpiredTokenUI errorCode="token_expired" onRequestNewLink={handleRequestNew} />);
    
    const heading = screen.getByRole('heading', { name: /link expired/i });
    expect(heading).toBeInTheDocument();
    
    // Check focus management
    expect(heading).toHaveFocus();
    
    // Check aria-live
    expect(heading).toHaveAttribute('aria-live', 'polite');

    const button = screen.getByRole('button', { name: /request new link/i });
    fireEvent.click(button);
    expect(handleRequestNew).toHaveBeenCalled();
  });

  it('renders revoked token state', () => {
    render(<ExpiredTokenUI errorCode="token_revoked" />);
    expect(screen.getByRole('heading', { name: /link revoked/i })).toBeInTheDocument();
  });

  it('renders consumed token state', () => {
    render(<ExpiredTokenUI errorCode="token_consumed" />);
    expect(screen.getByRole('heading', { name: /link already used/i })).toBeInTheDocument();
  });

  it('renders invalid token state without retry button', () => {
    render(<ExpiredTokenUI errorCode="token_invalid" />);
    expect(screen.getByRole('heading', { name: /invalid link/i })).toBeInTheDocument();
    expect(screen.queryByRole('button', { name: /request new link/i })).not.toBeInTheDocument();
    expect(screen.queryByRole('button', { name: /try again/i })).not.toBeInTheDocument();
    expect(screen.getByRole('button', { name: /return to home/i })).toBeInTheDocument();
  });

  it('renders network error state with retry button', () => {
    const handleRetry = vi.fn();
    render(<ExpiredTokenUI errorCode="network_error" onRetry={handleRetry} />);
    expect(screen.getByRole('heading', { name: /connection error/i })).toBeInTheDocument();
    
    const button = screen.getByRole('button', { name: /try again/i });
    fireEvent.click(button);
    expect(handleRetry).toHaveBeenCalled();
  });
});
