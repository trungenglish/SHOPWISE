import React, { useState } from "react";
import {
  Cpu,
  ShoppingCart,
  Archive,
  ShieldCheck,
  Landmark,
  ChevronRight,
  ChevronLeft,
} from "lucide-react";
import { AgentStatus } from "../types";

interface AgentHubProps {
  agents: AgentStatus[];
}

export default function AgentHub({ agents }: AgentHubProps) {
  const [isOpen, setIsOpen] = useState(true);

  // Helper to map type to icon
  const getIcon = (type: string, className?: string) => {
    switch (type) {
      case "shopping":
        return <ShoppingCart size={14} className={className || "text-[#4F7CFF]"} />;
      case "inventory":
        return <Archive size={14} className={className || "text-secondary"} />;
      case "warranty":
        return <ShieldCheck size={14} className={className || "text-yellow-400"} />;
      case "finance":
        return <Landmark size={14} className={className || "text-purple-400"} />;
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

  return (
    <aside
      className={`glass-panel border-outline-variant/15 flex h-full shrink-0 flex-col border-l shadow-[-10px_0_30px_rgba(0,0,0,0.5)] select-none transition-all duration-300 ${
        isOpen ? "w-[320px] p-5" : "w-16 items-center p-3"
      }`}
    >
      {!isOpen ? (
        <div className="flex h-full flex-col items-center">
          <button
            onClick={() => setIsOpen(true)}
            className="hover:bg-surface-high text-on-surface-variant cursor-pointer rounded-lg p-2 transition-colors"
            title="Expand Agent Hub"
          >
            <ChevronLeft size={20} />
          </button>

          <div className="mt-8 flex flex-1 flex-col items-center gap-6">
            {agents.map((agent) => (
              <div
                key={agent.id}
                className={`relative flex h-10 w-10 items-center justify-center rounded-full border border-transparent transition-all ${
                  agent.progress === 0
                    ? "opacity-60"
                    : "bg-surface-low shadow-[0_0_10px_rgba(0,0,0,0.2)]"
                }`}
                title={`${agent.name} - ${
                  agent.progress === 0 ? "QUEUED" : `${agent.progress}%`
                }`}
              >
                {getIcon(agent.type, "scale-125")}
                {agent.progress > 0 && agent.progress < 100 && (
                  <span className="absolute -top-1 -right-1 flex h-3 w-3">
                    <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-[#4F7CFF] opacity-75"></span>
                    <span className="relative inline-flex h-3 w-3 rounded-full bg-[#4F7CFF]"></span>
                  </span>
                )}
              </div>
            ))}
          </div>
        </div>
      ) : (
        <>
          <div className="mb-6 flex items-start justify-between">
            <div>
              <h2 className="font-display text-on-surface text-base font-bold">
                Agent Hub
              </h2>
              <p className="mt-1.5 flex items-center gap-2 font-mono text-[11px] text-[#4F7CFF]">
                <span className="shadow-[0_0_8px_#4F7CFF] h-2.5 w-2.5 animate-pulse rounded-full bg-[#4F7CFF]"></span>
                Inter-Agent Collaboration
              </p>
            </div>
            <button
              onClick={() => setIsOpen(false)}
              className="hover:bg-surface-high text-on-surface-variant cursor-pointer rounded p-1 transition-colors"
              title="Collapse Agent Hub"
            >
              <ChevronRight size={18} />
            </button>
          </div>

          <div className="flex-1 overflow-y-auto pr-1">
            <div className="flex flex-col">
              {agents.map((agent, index) => {
                const isQueued = agent.progress === 0;
                const isCompleted = agent.progress === 100;

                return (
                  <React.Fragment key={agent.id}>
                    {/* Agent Card */}
                    <div
                      className={`bg-surface-low border-outline-variant/15 hover:border-outline-variant/30 relative flex flex-col gap-2 overflow-hidden rounded-xl border border-l-4 p-3.5 transition-all duration-300 ${getStatusColorClass(
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
                          className={`h-full transition-all duration-500 ${getProgressBg(
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
