"use client";

import { format } from "date-fns";
import { Calendar as CalendarIcon, ChevronDownIcon } from "lucide-react";

import { Button } from "@shopwise/ui/components/button";
import { Calendar } from "@shopwise/ui/components/calendar";
import { Field, FieldGroup, FieldLabel } from "@shopwise/ui/components/field";
import { Input } from "@shopwise/ui/components/input";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@shopwise/ui/components/popover";
import type { DateRange } from "react-day-picker";

interface DatePickerProps {
  readonly value: Date;
  readonly onChange: (date: Date) => void;
}

interface DatePickerWithRangeProps {
  readonly value: DateRange;
  readonly onChange: (date: DateRange) => void;
}

interface DatePickerTimeProps {
  readonly value: Date | undefined;
  readonly onChange: (date: Date) => void;
}

export const DatePicker = ({ value, onChange }: DatePickerProps) => {
  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          data-empty={!value}
          className="data-[empty=true]:text-muted-foreground w-[280px] justify-start text-left font-normal"
        >
          <CalendarIcon />
          {value ? format(value, "PPP") : <span>Pick a date</span>}
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-auto p-0">
        <Calendar mode="single" selected={value} onSelect={onChange} required />
      </PopoverContent>
    </Popover>
  );
};

export const DatePickerWithRange = ({
  value,
  onChange,
}: DatePickerWithRangeProps) => {
  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          id="date-picker-range"
          className="justify-start px-2.5 font-normal"
        >
          <CalendarIcon />
          {value?.from ? (
            value.to ? (
              <>
                {format(value.from, "LLL dd, y")} -{" "}
                {format(value.to, "LLL dd, y")}
              </>
            ) : (
              format(value.from, "LLL dd, y")
            )
          ) : (
            <span>Pick a date</span>
          )}
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-auto p-0" align="start">
        <Calendar
          mode="range"
          defaultMonth={value?.from}
          selected={value}
          onSelect={onChange}
          numberOfMonths={2}
          required
        />
      </PopoverContent>
    </Popover>
  );
};

export const DatePickerTime = ({ value, onChange }: DatePickerTimeProps) => {
  return (
    <FieldGroup className="mx-auto max-w-xs flex-row">
      <Field>
        <FieldLabel htmlFor="date-picker-optional">Date</FieldLabel>
        <Popover>
          <PopoverTrigger asChild>
            <Button
              variant="outline"
              id="date-picker-optional"
              className="w-32 justify-between font-normal"
            >
              {value ? format(value, "PPP") : "Select date"}
              <ChevronDownIcon />
            </Button>
          </PopoverTrigger>
          <PopoverContent className="w-auto overflow-hidden p-0" align="start">
            <Calendar
              mode="single"
              selected={value}
              captionLayout="dropdown"
              defaultMonth={value}
              onSelect={(date) => {
                if (date) {
                  onChange(date);
                }
              }}
            />
          </PopoverContent>
        </Popover>
      </Field>
      <Field className="w-32">
        <FieldLabel htmlFor="time-picker-optional">Time</FieldLabel>
        <Input
          type="time"
          id="time-picker-optional"
          step="1"
          defaultValue="10:30:00"
          className="bg-background appearance-none [&::-webkit-calendar-picker-indicator]:hidden [&::-webkit-calendar-picker-indicator]:appearance-none"
        />
      </Field>
    </FieldGroup>
  );
};
