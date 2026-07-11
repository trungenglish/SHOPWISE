package uiprotocol

type ButtonProps struct {
	Text    string  `json:"text"`
	Variant *string `json:"variant,omitempty"`
}

type QuickReplyProps struct {
	Label string `json:"label"`
}

type SelectOption struct {
	Label *string `json:"label,omitempty"`
	Value *string `json:"value,omitempty"`
}

type SelectProps struct {
	Options     []SelectOption `json:"options,omitempty"`
	Placeholder *string        `json:"placeholder,omitempty"`
}

type RadioGroupProps struct {
	Options []SelectOption `json:"options,omitempty"`
	Name    *string        `json:"name,omitempty"`
}

type CheckboxProps struct {
	Label   *string `json:"label,omitempty"`
	Checked *bool   `json:"checked,omitempty"`
}

type TextInputProps struct {
	Placeholder  *string `json:"placeholder,omitempty"`
	DefaultValue *string `json:"defaultValue,omitempty"`
	Type         *string `json:"type,omitempty"`
}

type NumberInputProps struct {
	Placeholder  *string  `json:"placeholder,omitempty"`
	DefaultValue *float64 `json:"defaultValue,omitempty"`
}

type BudgetSliderProps struct {
	MinPriceVND int  `json:"minPriceVND"`
	MaxPriceVND int  `json:"maxPriceVND"`
	Step        *int `json:"step,omitempty"`
}

type Store struct {
	StoreID    *string  `json:"storeId,omitempty"`
	Name       *string  `json:"name,omitempty"`
	Address    *string  `json:"address,omitempty"`
	DistanceKm *float64 `json:"distanceKm,omitempty"`
}

type StorePickerProps struct {
	Stores []Store `json:"stores,omitempty"`
}
