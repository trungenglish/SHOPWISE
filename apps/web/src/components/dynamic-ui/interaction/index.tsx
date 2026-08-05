import React, { useState } from 'react';
import type {
  ButtonProps,
  QuickReplyProps,
  SelectProps,
  RadioGroupProps,
  CheckboxProps,
  TextInputProps,
  NumberInputProps,
  BudgetSliderProps,
  StorePickerProps
} from '@shopwise/ui-protocol/types/interaction';
import { dynamicUIEventDispatcher } from '../EventDispatcher';

const formatVND = (amount: number) => {
  return new Intl.NumberFormat('vi-VN', { style: 'currency', currency: 'VND' }).format(amount);
};

export const Button: React.FC<ButtonProps & { id: string }> = ({ id, text, variant = 'primary' }) => (
  <button
    className={`btn btn-${variant}`}
    onClick={() => dynamicUIEventDispatcher.dispatch(id, 'SUBMIT', { text })}
  >
    {text}
  </button>
);

export const QuickReply: React.FC<QuickReplyProps & { id: string }> = ({ id, label }) => (
  <button
    className="quick-reply"
    onClick={() => dynamicUIEventDispatcher.dispatch(id, 'QUICK_REPLY', { label })}
  >
    {label}
  </button>
);

export const Select: React.FC<SelectProps & { id: string }> = ({ id, options, placeholder }) => {
  const [val, setVal] = useState('');
  return (
    <select
      className="interaction-select"
      value={val}
      onChange={(e) => {
        setVal(e.target.value);
        dynamicUIEventDispatcher.dispatch(id, 'SELECT', { value: e.target.value });
      }}
    >
      <option value="" disabled>{placeholder || 'Select an option'}</option>
      {options?.map((opt, i) => (
        <option key={i} value={opt.value}>{opt.label}</option>
      ))}
    </select>
  );
};

export const RadioGroup: React.FC<RadioGroupProps & { id: string }> = ({ id, options, name }) => (
  <div className="radio-group">
    {options?.map((opt, i) => (
      <label key={i}>
        <input
          type="radio"
          name={name || id}
          value={opt.value}
          onChange={(e) => dynamicUIEventDispatcher.dispatch(id, 'SELECT', { value: e.target.value })}
        />
        {opt.label}
      </label>
    ))}
  </div>
);

export const Checkbox: React.FC<CheckboxProps & { id: string }> = ({ id, label, checked }) => {
  const [isChecked, setIsChecked] = useState(checked || false);
  return (
    <label className="checkbox-label">
      <input
        type="checkbox"
        checked={isChecked}
        onChange={(e) => {
          setIsChecked(e.target.checked);
          dynamicUIEventDispatcher.dispatch(id, 'TOGGLE', { checked: e.target.checked });
        }}
      />
      {label}
    </label>
  );
};

export const TextInput: React.FC<TextInputProps & { id: string }> = ({ id, placeholder, defaultValue, type = 'text' }) => {
  const [val, setVal] = useState(defaultValue || '');
  return (
    <input
      type={type}
      placeholder={placeholder}
      value={val}
      onChange={(e) => setVal(e.target.value)}
      onBlur={() => dynamicUIEventDispatcher.dispatch(id, 'CHANGE', { value: val })}
      className="text-input"
    />
  );
};

export const NumberInput: React.FC<NumberInputProps & { id: string }> = ({ id, placeholder, defaultValue }) => {
  const [val, setVal] = useState(defaultValue ?? '');
  return (
    <input
      type="number"
      placeholder={placeholder}
      value={val}
      onChange={(e) => setVal(e.target.value ? Number(e.target.value) : '')}
      onBlur={() => dynamicUIEventDispatcher.dispatch(id, 'CHANGE', { value: val })}
      className="number-input"
    />
  );
};

export const BudgetSlider: React.FC<BudgetSliderProps & { id: string }> = ({ id, minPriceVND, maxPriceVND, step = 100000 }) => {
  const [val, setVal] = useState(minPriceVND);
  return (
    <div className="budget-slider">
      <input
        type="range"
        min={minPriceVND}
        max={maxPriceVND}
        step={step}
        value={val}
        onChange={(e) => setVal(Number(e.target.value))}
        onMouseUp={() => dynamicUIEventDispatcher.dispatch(id, 'SET_BUDGET', { valueVND: val })}
        onTouchEnd={() => dynamicUIEventDispatcher.dispatch(id, 'SET_BUDGET', { valueVND: val })}
      />
      <span>Budget: {formatVND(val)}</span>
    </div>
  );
};

export const StorePicker: React.FC<StorePickerProps & { id: string }> = ({ id, stores }) => (
  <div className="store-picker">
    {stores?.map((store, i) => (
      <div
        key={i}
        className="store-item"
        onClick={() => dynamicUIEventDispatcher.dispatch(id, 'SELECT_STORE', { storeId: store.storeId })}
      >
        <strong>{store.name}</strong> - {store.distanceKm}km away
      </div>
    ))}
  </div>
);
