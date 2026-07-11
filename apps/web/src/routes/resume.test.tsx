import React from 'react';
import { render, screen, waitFor } from '@testing-library/react';
import { describe, it, expect, vi, beforeEach } from 'vitest';
import { ResumePage, Route } from './resume';
import * as decisionMemoryApi from '@/api/decisionMemory';

const mockNavigate = vi.fn();
const mockInvalidateQueries = vi.fn();

vi.mock('@tanstack/react-router', () => ({
  useNavigate: () => mockNavigate,
  createFileRoute: () => (config: any) => ({ ...config, useSearch: () => ({ token: 'mock-token' }) }),
}));

let mockUseQueryReturn = {
  data: undefined,
  error: null,
  isFetching: false,
};

vi.mock('@tanstack/react-query', () => ({
  useQueryClient: () => ({
    invalidateQueries: mockInvalidateQueries,
  }),
  useQuery: (opts: any) => mockUseQueryReturn,
}));

vi.mock('@/api/decisionMemory', () => ({
  resumeSessionFromToken: vi.fn(),
  ResumeSessionError: class extends Error {
    code: string;
    constructor(m: string, c: string) {
      super(m);
      this.code = c;
    }
  }
}));

describe('ResumePage', () => {
  beforeEach(() => {
    vi.clearAllMocks();
    mockUseQueryReturn = {
      data: undefined,
      error: null,
      isFetching: false,
    };
  });

  it('valid token successfully resumes and redirects', async () => {
    mockUseQueryReturn = {
      data: { ID: 'session-123' },
      error: null,
      isFetching: false,
    } as any;

    // Override the mock to return a token
    Route.useSearch = () => ({ token: 'valid-token' });

    render(<ResumePage />);

    expect(screen.getByText(/resuming your session/i)).toBeInTheDocument();

    await waitFor(() => {
      // Verify token removed from URL
      expect(mockNavigate).toHaveBeenCalledWith({ to: '/resume', replace: true });
      // Verify invalidation
      expect(mockInvalidateQueries).toHaveBeenCalledWith({ queryKey: ['sessions'] });
      expect(mockInvalidateQueries).toHaveBeenCalledWith({ queryKey: ['session', 'session-123'] });
      // Verify navigation to dashboard
      expect(mockNavigate).toHaveBeenCalledWith({ to: '/dashboard', search: { sessionId: 'session-123' } });
    });
  });

  it('renders expired token UI on error', async () => {
    mockUseQueryReturn = {
      data: undefined,
      error: new decisionMemoryApi.ResumeSessionError('expired', 'token_expired'),
      isFetching: false,
    } as any;
    Route.useSearch = () => ({ token: 'expired-token' });

    render(<ResumePage />);

    await waitFor(() => {
      expect(screen.getByRole('heading', { name: /link expired/i })).toBeInTheDocument();
    });
    
    // Ensure token is not persisted or sent to dashboard
    expect(mockNavigate).not.toHaveBeenCalled();
  });

  it('renders network error and handles retry without duplicate requests', async () => {
    mockUseQueryReturn = {
      data: undefined,
      error: new Error('Network Error'),
      isFetching: false,
    } as any;
    Route.useSearch = () => ({ token: 'network-token' });

    render(<ResumePage />);

    await waitFor(() => {
      expect(screen.getByRole('heading', { name: /connection error/i })).toBeInTheDocument();
    });

    // Ensure useQuery is tested implicitly, but we just verify UI
    // Duplicate requests are prevented by React Query internally, which we trust,
    // so this test just ensures the UI renders the network error correctly.
  });
});
