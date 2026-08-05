import { fireEvent, render, screen } from '@testing-library/react'
import { expect, it, vi, beforeEach } from 'vitest'

window.HTMLElement.prototype.scrollIntoView = vi.fn()

import type { AgentEnvelope } from '@/api/decision-memory'
import AuditTrail from './audit-trail'

type QuestionEnvelope = Extract<AgentEnvelope, { type: 'question' }>

const questionEnvelope: QuestionEnvelope = {
  type: 'question',
  message: 'What gaming style do you prefer?',
  schema_version: '1.1',
  turn_id: 'turn-q1',
  revision: 1,
  conversation_state: 'collecting_requirements',
  ui_state: { status: 'ready' },
  ui_operations: [],
  question: {
    id: 'gaming-style-q',
    mode: 'single',
    free_text_allowed: true,
    input_label: 'Your choice',
    input_placeholder: 'e.g. AAA games',
    submit_label: 'Continue',
    options: [
      { id: 'aaa', label: 'AAA story games' },
      { id: 'esports', label: 'Competitive esports' },
      { id: 'casual', label: 'Casual / indie' },
    ],
  },
}

const defaultProps = {
  logs: [
    { time: '10:00', message: 'Analyzing request', status: 'done' as const },
    { time: '10:01', message: 'Loading catalog', status: 'done' as const },
  ],
  userIntent: 'I need a gaming laptop.',
  onInjectConstraint: vi.fn(),
  isLoading: false,
  onReplay: vi.fn(),
}

it('renders active question card after completed log entries', () => {
  render(
    <AuditTrail
      {...defaultProps}
      activeQuestion={questionEnvelope}
      onInteraction={vi.fn()}
    />,
  )

  // All log items are present
  expect(screen.getByText('Analyzing request')).toBeInTheDocument()
  expect(screen.getByText('Loading catalog')).toBeInTheDocument()

  // Active question message is rendered inside the ClarificationCard
  expect(
    screen.getByText('What gaming style do you prefer?'),
  ).toBeInTheDocument()

  // The three quick-reply options appear as labels
  expect(screen.getByText('AAA story games')).toBeInTheDocument()
  expect(screen.getByText('Competitive esports')).toBeInTheDocument()
  expect(screen.getByText('Casual / indie')).toBeInTheDocument()
})

it('calls onInteraction with selected option after the user clicks the submit button', () => {
  const onInteraction = vi.fn()
  render(
    <AuditTrail
      {...defaultProps}
      activeQuestion={questionEnvelope}
      onInteraction={onInteraction}
    />,
  )

  // The ClarificationCard uses option.label as visible label text inside <label>
  // The submit button label comes from question.submit_label
  const option = screen.getByDisplayValue('aaa')
  fireEvent.click(option)
  expect(onInteraction).not.toHaveBeenCalled()

  fireEvent.click(screen.getByRole('button', { name: 'Continue' }))
  expect(onInteraction).toHaveBeenCalledWith(
    expect.objectContaining({
      interaction_id: 'gaming-style-q',
      action: 'question.answer',
      payload: expect.objectContaining({ selected_option_ids: ['aaa'] }),
    }),
  )
})

it('does not render a ClarificationCard when activeQuestion is null', () => {
  render(
    <AuditTrail
      {...defaultProps}
      activeQuestion={null}
      onInteraction={vi.fn()}
    />,
  )

  expect(
    screen.queryByText('What gaming style do you prefer?'),
  ).not.toBeInTheDocument()

  // Regular logs still render
  expect(screen.getByText('Analyzing request')).toBeInTheDocument()
})

it('calls onInjectConstraint when the free-text composer form is submitted', () => {
  const onInjectConstraint = vi.fn()
  render(
    <AuditTrail
      {...defaultProps}
      onInjectConstraint={onInjectConstraint}
      activeQuestion={null}
    />,
  )

  const input = screen.getByPlaceholderText(
    /Add new constraint/i,
  ) as HTMLInputElement
  fireEvent.change(input, { target: { value: 'OLED screen' } })
  // Submit the form by pressing Enter (the submit button is icon-only)
  fireEvent.submit(input.closest('form') as HTMLFormElement)

  expect(onInjectConstraint).toHaveBeenCalledWith('OLED screen')
})
