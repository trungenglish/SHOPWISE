import { describe, it, expect, vi } from 'vitest';
import { render, screen, fireEvent } from '@testing-library/react';
import { CheckoutIncentivePanel } from './CheckoutIncentivePanel';
import { useCheckoutIncentive } from '../hooks/useCheckoutIncentive';

vi.mock('../hooks/useCheckoutIncentive');

const mockUseCheckoutIncentive = useCheckoutIncentive as any;

describe('CheckoutIncentivePanel', () => {
  it('renders nothing when HIDDEN', () => {
    mockUseCheckoutIncentive.mockReturnValue({
      displayState: 'HIDDEN',
    });
    const { container } = render(<CheckoutIncentivePanel />);
    expect(container.firstChild).toBeNull();
  });

  it('renders ACTIVE state with countdown', () => {
    mockUseCheckoutIncentive.mockReturnValue({
      displayState: 'ACTIVE',
      isCollapsed: false,
      currentPromotion: { status: 'ACTIVE', reward_value: 15 },
      remainingSeconds: 120,
    });
    render(<CheckoutIncentivePanel />);
    
    expect(screen.getByText(/Complete your payment within 15 minutes/i)).toBeInTheDocument();
    expect(screen.getByText('02:00')).toBeInTheDocument();
  });

  it('renders the collapsed badge when ACTIVE and isCollapsed is true', () => {
    mockUseCheckoutIncentive.mockReturnValue({
      displayState: 'ACTIVE',
      isCollapsed: true,
      currentPromotion: { status: 'ACTIVE', reward_value: 15 },
      remainingSeconds: 120,
      expand: vi.fn(),
    });
    render(<CheckoutIncentivePanel />);
    
    expect(screen.getByText(/15% Voucher · /i)).toBeInTheDocument();
    expect(screen.queryByText(/Complete your payment/i)).not.toBeInTheDocument();
  });

  it('renders SUCCESS state', () => {
    mockUseCheckoutIncentive.mockReturnValue({
      displayState: 'SUCCESS',
      isCollapsed: false,
      currentPromotion: { status: 'VOUCHER_ISSUED', reward_value: 15 },
    });
    render(<CheckoutIncentivePanel />);
    
    expect(screen.getByText('Voucher earned!')).toBeInTheDocument();
  });

  it('renders EXPIRED state', () => {
    mockUseCheckoutIncentive.mockReturnValue({
      displayState: 'EXPIRED',
      isCollapsed: false,
      currentPromotion: { status: 'EXPIRED' },
    });
    render(<CheckoutIncentivePanel />);
    
    expect(screen.getByText('Offer Expired')).toBeInTheDocument();
  });

  it('calls collapse when collapse button is clicked', () => {
    const mockCollapse = vi.fn();
    mockUseCheckoutIncentive.mockReturnValue({
      displayState: 'ACTIVE',
      isCollapsed: false,
      currentPromotion: { status: 'ACTIVE' },
      remainingSeconds: 60,
      collapse: mockCollapse,
    });
    render(<CheckoutIncentivePanel />);
    
    fireEvent.click(screen.getByLabelText('Collapse panel'));
    expect(mockCollapse).toHaveBeenCalledTimes(1);
  });
});
