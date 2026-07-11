import React, { useState } from "react";
import { Sparkles, Send, Search, History } from "lucide-react";

interface EmptyWorkspaceProps {
  onSubmit: (query: string) => void;
}

export default function EmptyWorkspace({ onSubmit }: EmptyWorkspaceProps) {
  const [inputValue, setInputValue] = useState("");

  const handleSubmit = (e: React.SubmitEvent) => {
    e.preventDefault();
    if (inputValue.trim()) {
      onSubmit(inputValue.trim());
    }
  };

  const suggestions = [
    "I want to find laptop ABC",
    "Compare top gaming rigs under $2000",
    "Best thin & light for programming",
    "High refresh rate OLED screen laptops",
  ];

  return (
    <div className="flex h-full w-full flex-col items-center justify-center p-8 select-none">
      <div className="flex w-full max-w-3xl flex-col items-center gap-8">
        {/* Hero Section */}
        <div className="flex flex-col items-center text-center">
          <div className="mb-4 rounded-2xl bg-[#4F7CFF]/10 p-4 text-[#4F7CFF]">
            <Sparkles size={40} className="animate-pulse" />
          </div>
          <h1 className="font-display text-on-surface mb-2 text-4xl font-black tracking-tight">
            What are you looking for today?
          </h1>
          <p className="text-on-surface-variant font-sans text-base">
            Describe your ideal hardware, and SHOPWISE AI will orchestrate
            multiple agents to find the perfect match.
          </p>
        </div>

        {/* Search Input */}
        <div className="w-full">
          <form
            onSubmit={handleSubmit}
            className="bg-surface-highest border-outline-variant/30 flex w-full items-center rounded-2xl border p-2 shadow-2xl transition-all focus-within:border-[#4F7CFF] focus-within:shadow-[0_0_40px_rgba(79,124,255,0.15)]"
          >
            <div className="text-on-surface-variant flex h-12 w-12 items-center justify-center">
              <Search size={24} />
            </div>
            <input
              value={inputValue}
              onChange={(e) => setInputValue(e.target.value)}
              className="text-on-surface placeholder:text-on-surface-variant/40 h-14 w-full border-none bg-transparent px-2 font-sans text-lg outline-none focus:ring-0"
              placeholder="e.g. Intensive 3D rendering and local AI development..."
              type="text"
            />
            <button
              type="submit"
              disabled={!inputValue.trim()}
              className="disabled:bg-surface-highest disabled:text-on-surface-variant/30 flex h-12 w-12 cursor-pointer items-center justify-center rounded-xl bg-[#4F7CFF] text-white transition-colors hover:bg-[#4F7CFF]/90 disabled:cursor-not-allowed"
            >
              <Send size={20} />
            </button>
          </form>
        </div>

        {/* Suggestion Chips */}
        <div className="flex w-full flex-wrap items-center justify-center gap-3">
          {suggestions.map((suggestion, idx) => (
            <button
              key={idx}
              onClick={() => onSubmit(suggestion)}
              className="glass-card border-outline-variant/20 text-on-surface-variant hover:text-on-surface cursor-pointer rounded-xl border px-4 py-2.5 font-sans text-sm font-medium transition-all hover:border-[#4F7CFF]/40 hover:bg-[#4F7CFF]/5"
            >
              {suggestion}
            </button>
          ))}
        </div>

        {/* Recent Activity (Mock) */}
        <div className="mt-8 flex w-full flex-col gap-3">
          <div className="text-on-surface-variant flex items-center gap-2 font-mono text-xs font-bold tracking-wider uppercase">
            <History size={14} />
            <span>Continue previous sessions</span>
          </div>
          <div className="grid grid-cols-1 gap-3 md:grid-cols-2">
            <div className="border-outline-variant/15 bg-surface-low hover:bg-surface-high cursor-pointer rounded-xl border p-4 transition-colors">
              <h4 className="text-on-surface font-sans text-sm font-bold">
                Hardware Render OS
              </h4>
              <p className="text-on-surface-variant mt-1 text-xs">
                Last edited 2 hours ago
              </p>
            </div>
            <div className="border-outline-variant/15 bg-surface-low hover:bg-surface-high cursor-pointer rounded-xl border p-4 transition-colors">
              <h4 className="text-on-surface font-sans text-sm font-bold">
                MacBook Pro M3 Max Specs
              </h4>
              <p className="text-on-surface-variant mt-1 text-xs">
                Last edited yesterday
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
