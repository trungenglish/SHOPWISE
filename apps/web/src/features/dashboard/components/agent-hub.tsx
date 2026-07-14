import React, { useState } from "react";
import {
  Cpu,
  ShoppingCart,
  Archive,
  ShieldCheck,
  Landmark,
  X,
  Bot,
} from "lucide-react";
import { AgentStatus } from "../types";

interface AgentHubProps {
  agents: AgentStatus[];
}

export default function AgentHub({ agents }: AgentHubProps) {
  const [isOpen, setIsOpen] = useState(false);

  // Helper to map type to icon
  const getIcon = (type: string, className?: string) => {
    switch (type) {
      case "shopping":
        return (
          <ShoppingCart size={14} className={className || "text-[#4F7CFF]"} />
        );
      case "inventory":
        return <Archive size={14} className={className || "text-secondary"} />;
      case "warranty":
        return (
          <ShieldCheck size={14} className={className || "text-yellow-400"} />
        );
      case "finance":
        return (
          <Landmark size={14} className={className || "text-purple-400"} />
        );
      default:
        return <Cpu size={14} className={className} />;
    }
  };

  // Helper to get border color and text classes
  const getStatusColorClass = (type: string, progress: number) => {
    if (progress === 0) return "border-l-outline-variant/30";
    switch (type) {
      case "shopping":
        return "border-l-[#4F7CFF]";
      case "inventory":
        return "border-l-[#c8c6c8]";
      case "warranty":
        return "border-l-yellow-400";
      case "finance":
        return "border-l-purple-400";
      default:
        return "border-l-outline";
    }
  };

  const getProgressBg = (type: string) => {
    switch (type) {
      case "shopping":
        return "bg-[#4F7CFF]";
      case "inventory":
        return "bg-secondary";
      case "warranty":
        return "bg-yellow-400";
      case "finance":
        return "bg-purple-400";
      default:
        return "bg-outline";
    }
  };

  const hasActiveAgents = agents.some(
    (a) => a.progress > 0 && a.progress < 100
  );

  return (
    <aside
      className={`fixed right-6 bottom-6 z-50 flex flex-col transition-[transform,opacity,background-color,border-color,box-shadow,max-height,width] duration-200 ease-[var(--ease-out)] select-none ${
        isOpen
          ? "glass-panel border-outline-variant/15 max-h-[70vh] w-[320px] rounded-2xl border p-5 shadow-[0_15px_40px_rgba(0,0,0,0.6)]"
          : "glass-panel border-outline-variant/15 hover:bg-surface-high h-14 w-14 cursor-pointer items-center justify-center rounded-full border shadow-[0_10px_25px_rgba(0,0,0,0.5)] hover:scale-105 max-md:hover:scale-100"
      }`}
      onClick={() => !isOpen && setIsOpen(true)}
    >
      {!isOpen ? (
        <div className="relative flex h-full w-full items-center justify-center text-[#4F7CFF]">
          <Bot size={24} />
          {hasActiveAgents && (
            <span className="absolute top-0 right-0 flex h-3 w-3">
              <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-[#4F7CFF] opacity-75"></span>
              <span className="relative inline-flex h-3 w-3 rounded-full bg-[#4F7CFF]"></span>
            </span>
          )}
        </div>
      ) : (
        <>
          <div className="mb-6 flex items-start justify-between">
            <div>
              <h2 className="font-display text-on-surface text-base font-bold">
                Agent Hub
              </h2>
              <p className="mt-1.5 flex items-center gap-2 font-mono text-[11px] text-[#4F7CFF]">
                <span className="h-2.5 w-2.5 animate-pulse rounded-full bg-[#4F7CFF] shadow-[0_0_8px_#4F7CFF]"></span>
                Inter-Agent Collaboration
              </p>
            </div>
            <button
              onClick={(e) => {
                e.stopPropagation();
                setIsOpen(false);
              }}
              className="hover:bg-surface-high text-on-surface-variant cursor-pointer rounded p-1 transition-colors"
              title="Close Agent Hub"
            >
              <X size={18} />
            </button>
          </div>

          <div
            className="flex-1 overflow-y-auto pr-1"
            onClick={(e) => e.stopPropagation()}
          >
            <div className="flex flex-col">
              {agents.map((agent, index) => {
                const isQueued = agent.progress === 0;
                const isCompleted = agent.progress === 100;

                return (
                  <React.Fragment key={agent.id}>
                    {/* Agent Card */}
                    <div
                      className={`bg-surface-low border-outline-variant/15 hover:border-outline-variant/30 relative flex flex-col gap-2 overflow-hidden rounded-xl border border-l-4 p-3.5 transition-[border-color,background-color,opacity] duration-200 ease-[var(--ease-out)] ${getStatusColorClass(
                        agent.type,
                        agent.progress
                      )} ${isQueued ? "opacity-60" : ""}`}
                    >
                      <div className="flex items-center justify-between">
                        <div className="flex items-center gap-2.5">
                          {getIcon(agent.type)}
                          <span className="font-display text-on-surface text-xs font-bold">
                            {agent.name}
                          </span>
                        </div>
                        <span className="text-on-surface-variant font-mono text-[10px] font-bold">
                          {isQueued ? "QUEUED" : `${agent.progress}%`}
                        </span>
                      </div>

                      <p className="text-on-surface-variant font-mono text-[11px] leading-relaxed">
                        {agent.statusMessage}
                      </p>

                      <div className="bg-surface-highest mt-1 h-1 w-full overflow-hidden rounded-full">
                        <div
                          className={`h-full transition-[width] duration-300 ease-[var(--ease-out)] ${getProgressBg(
                            agent.type
                          )}`}
                          style={{ width: `${agent.progress}%` }}
                        />
                      </div>
                    </div>

                    {/* Connection Stream line between agents (except the last one) */}
                    {index < agents.length - 1 && (
                      <div className="relative flex h-6 w-full justify-center">
                        <div className="bg-outline-variant/20 relative h-full w-px">
                          {!isQueued && (
                            <>
                              <div
                                className="comm-dot"
                                style={{ animationDelay: "0s" }}
                              ></div>
                              <div
                                className="comm-dot"
                                style={{ animationDelay: "1s" }}
                              ></div>
                            </>
                          )}
                        </div>
                      </div>
                    )}
                  </React.Fragment>
                );
              })}
            </div>
          </div>
        </>
      )}
    </aside>
  );
}
