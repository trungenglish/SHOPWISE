import React, { useState } from "react";
import {
  RotateCw,
  HelpCircle,
  Send,
  CheckCircle2,
  Loader2,
  Sparkles,
} from "lucide-react";
import { AuditLog } from "../types";

interface AuditTrailProps {
  logs: AuditLog[];
  userIntent: string;
  onInjectConstraint: (constraint: string) => void;
  isLoading: boolean;
  onReplay: () => void;
  suggestions?: string[];
  hasMenuButton?: boolean;
}

export default function AuditTrail({
  logs,
  userIntent,
  onInjectConstraint,
  isLoading,
  onReplay,
  suggestions,
  hasMenuButton,
}: AuditTrailProps) {
  const [inputValue, setInputValue] = useState("");

  const handleSubmit = (e: React.SubmitEvent) => {
    e.preventDefault();
    if (inputValue.trim()) {
      onInjectConstraint(inputValue.trim());
      setInputValue("");
    }
  };

  return (
    <section className="glass-panel border-outline-variant/15 flex h-screen w-[320px] shrink-0 flex-col border-r select-none">
      {/* Header */}
      <div className={`border-outline-variant/20 flex items-center justify-between border-b p-4 pt-6 ${hasMenuButton ? "pl-16" : ""}`}>
        <h2 className="font-display text-on-surface flex items-center gap-1.5 text-sm font-semibold tracking-wide uppercase">
          <Sparkles size={14} className="text-[#4F7CFF]" />
          Audit Trail
        </h2>
        <HelpCircle
          size={15}
          className="text-on-surface-variant cursor-pointer transition-colors hover:text-[#4F7CFF]"
        />
      </div>

      {/* Main Content Scroll Container */}
      <div className="flex flex-1 flex-col gap-4 overflow-y-auto p-4">
        {/* Replay decision path */}
        <button
          onClick={onReplay}
          disabled={isLoading}
          className="group flex w-full cursor-pointer items-center justify-center gap-2 rounded-lg border border-[#4F7CFF]/20 bg-[#4F7CFF]/5 px-4 py-3 font-mono text-[11px] font-bold text-[#4F7CFF] shadow-[0_0_15px_rgba(79,124,255,0.05)] transition-all hover:border-[#4F7CFF]/40 hover:bg-[#4F7CFF]/15 disabled:cursor-not-allowed disabled:opacity-50"
        >
          <RotateCw
            size={14}
            className={`transition-all duration-500 group-hover:rotate-180 ${isLoading ? "animate-spin" : ""}`}
          />
          REPLAY DECISION FLOW
        </button>

        {/* User Intent Box */}
        <div className="bg-surface-low border-outline-variant/15 hover:bg-surface-high rounded-lg border p-3 transition-colors">
          <div className="mb-1.5 flex items-center justify-between">
            <p className="text-on-surface-variant font-mono text-[10px] font-bold tracking-wider uppercase">
              USER INTENT
            </p>
          </div>
          <p className="text-on-surface font-sans text-xs leading-relaxed font-medium">
            {userIntent || "Waiting for shopping request setup..."}
          </p>
        </div>

        {/* Logs Timeline */}
        <div className="relative mt-2 flex flex-1 flex-col pl-1">
          <div className="bg-outline-variant/20 absolute top-[10px] bottom-[15px] left-[7px] w-0.5"></div>

          <div className="flex flex-col gap-5">
            {logs.map((log, index) => {
              const isRunning = log.status === "running";
              const isDone = log.status === "done";

              return (
                <div
                  key={index}
                  className="group relative flex flex-col gap-1 pl-6"
                >
                  {/* Status Node */}
                  <div className="absolute top-[2px] left-0 z-10 flex items-center justify-center">
                    {isDone ? (
                      <CheckCircle2
                        size={15}
                        className="bg-surface rounded-full text-green-400"
                      />
                    ) : isRunning ? (
                      <Loader2
                        size={15}
                        className="animate-spin text-[#4F7CFF] drop-shadow-[0_0_8px_rgba(79,124,255,0.8)]"
                      />
                    ) : (
                      <div className="bg-surface-low border-outline-variant/60 h-[11px] w-[11px] rounded-full border-2"></div>
                    )}
                  </div>

                  <div className="flex items-center justify-between">
                    <span
                      className={`font-mono text-[10px] font-semibold ${isRunning ? "text-[#4F7CFF]" : "text-on-surface-variant"}`}
                    >
                      {log.time} {isRunning && "(RUNNING)"}
                    </span>
                  </div>
                  <p
                    className={`font-sans text-[12px] leading-normal font-medium ${isRunning ? "text-[#4F7CFF]" : isDone ? "text-on-surface" : "text-on-surface-variant"}`}
                  >
                    {log.message}
                  </p>
                </div>
              );
            })}
          </div>
        </div>
      </div>

      {/* Input section */}
      <div className="border-outline-variant/15 bg-surface-lowest border-t p-4 flex flex-col gap-3">
        {suggestions && suggestions.length > 0 && (
          <div className="flex w-full gap-2 overflow-x-auto pb-1 scrollbar-hide snap-x">
            {suggestions.map((s, idx) => (
              <button
                key={idx}
                onClick={() => onInjectConstraint(s)}
                className="floating-chip bg-primary/10 border-primary/20 hover:border-primary/50 flex cursor-pointer shrink-0 items-center gap-2 rounded-full border px-3 py-1.5 text-left font-mono text-[10px] font-bold tracking-wide text-[#4F7CFF] backdrop-blur-md transition-all hover:scale-[1.03] active:scale-[0.97] snap-start"
              >
                <Sparkles size={11} />
                <span className="whitespace-nowrap">{s}</span>
              </button>
            ))}
          </div>
        )}

        <form
          onSubmit={handleSubmit}
          className="bg-surface border-outline-variant/30 flex items-center rounded-lg border p-1.5 transition-all focus-within:border-[#4F7CFF]/70 focus-within:ring-1 focus-within:ring-[#4F7CFF]/30"
        >
          <input
            value={inputValue}
            onChange={(e) => setInputValue(e.target.value)}
            disabled={isLoading}
            className="text-on-surface placeholder:text-on-surface-variant/40 w-full border-none bg-transparent px-2 py-1.5 font-sans text-xs outline-none focus:ring-0 focus:outline-none"
            placeholder="Add new constraint (e.g., OLED screen, under 80.000.000 ₫)..."
            type="text"
          />
          <button
            type="submit"
            disabled={isLoading || !inputValue.trim()}
            className="hover:bg-surface-high flex h-8 w-8 cursor-pointer items-center justify-center rounded-md text-[#4F7CFF] transition-colors disabled:opacity-30 disabled:hover:bg-transparent"
          >
            {isLoading ? (
              <Loader2
                size={16}
                className="text-on-surface-variant animate-spin"
              />
            ) : (
              <Send size={15} />
            )}
          </button>
        </form>
      </div>
    </section>
  );
}
