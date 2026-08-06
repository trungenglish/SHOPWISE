import React from "react";
import { Clock } from "lucide-react";

interface CountdownDisplayProps {
  remainingSeconds: number;
}

export const CountdownDisplay: React.FC<CountdownDisplayProps> = ({ remainingSeconds }) => {
  const minutes = Math.floor(remainingSeconds / 60);
  const seconds = remainingSeconds % 60;
  const timeString = `${minutes.toString().padStart(2, "0")}:${seconds.toString().padStart(2, "0")}`;

  // Add a warning color if under 2 minutes
  const isUrgent = remainingSeconds > 0 && remainingSeconds <= 120;

  return (
    <div 
      className={`flex items-center space-x-2 font-mono text-xl font-extrabold tracking-tight ${isUrgent ? 'text-rose-400 drop-shadow-[0_0_6px_rgba(251,113,133,0.5)] animate-pulse' : 'text-slate-100 drop-shadow-[0_0_6px_rgba(255,255,255,0.2)]'}`}
      aria-live="polite"
      aria-atomic="true"
    >
      <Clock className={`w-4 h-4 ${isUrgent ? 'text-rose-400' : 'text-cyan-400'}`} />
      <span className="tabular-nums">{timeString}</span>
      <span className="sr-only">remaining</span>
    </div>
  );
};
