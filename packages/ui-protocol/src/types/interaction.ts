export interface ButtonProps {
  text: string;
  variant?: "primary" | "secondary" | "danger" | "outline";
}

export interface QuickReplyProps {
  label: string;
}

export interface SelectOption {
  label?: string;
  value?: string;
}

export interface SelectProps {
  options?: SelectOption[];
  placeholder?: string;
}

export interface RadioGroupProps {
  options?: SelectOption[];
  name?: string;
}

export interface CheckboxProps {
  label?: string;
  checked?: boolean;
}

export interface TextInputProps {
  placeholder?: string;
  defaultValue?: string;
  type?: "text" | "email" | "password" | "tel";
}

export interface NumberInputProps {
  placeholder?: string;
  defaultValue?: number;
}

export interface BudgetSliderProps {
  minPriceVND: number;
  maxPriceVND: number;
  step?: number;
}

export interface Store {
  storeId?: string;
  name?: string;
  address?: string;
  distanceKm?: number;
}

export interface StorePickerProps {
  stores?: Store[];
}
