import { X, Sparkles, BookOpen } from "lucide-react";

interface ReasoningModalProps {
  isOpen: boolean;
  onClose: () => void;
  reasoning: string;
}

export default function ReasoningModal({
  isOpen,
  onClose,
  reasoning,
}: ReasoningModalProps) {
  if (!isOpen) return null;

  // Simple and highly effective Markdown to styled JSX block parser
  const renderReasoning = (rawText: string) => {
    if (!rawText) return null;

    const lines = rawText.split("\n");
    return lines.map((line, idx) => {
      // 1. Headers (###)
      if (line.startsWith("### ")) {
        return (
          <h4
            key={idx}
            className="font-display border-outline-variant/10 mt-5 mb-2 border-b pb-1 text-sm font-black tracking-wider text-[#4F7CFF] uppercase"
          >
            {line.substring(4)}
          </h4>
        );
      }
      if (line.startsWith("## ")) {
        return (
          <h3
            key={idx}
            className="font-display text-on-surface mt-6 mb-3 text-base font-extrabold tracking-tight"
          >
            {line.substring(3)}
          </h3>
        );
      }
      if (line.startsWith("# ")) {
        return (
          <h2
            key={idx}
            className="font-display text-on-surface mt-6 mb-4 text-lg font-black tracking-tight"
          >
            {line.substring(2)}
          </h2>
        );
      }

      // 2. Bullet list
      if (line.startsWith("- ") || line.startsWith("* ")) {
        // Handle basic bold match in bullets e.g. **bold**: text
        const content = line.substring(2);
        return (
          <div key={idx} className="my-1 flex items-start gap-2 pl-2">
            <span className="mt-1 shrink-0 text-[#4F7CFF]">•</span>
            <p className="text-on-surface font-sans text-xs leading-relaxed">
              {renderBoldText(content)}
            </p>
          </div>
        );
      }

      // 3. Normal paragraph
      if (line.trim() === "") return <div key={idx} className="h-2" />;

      return (
        <p
          key={idx}
          className="text-on-surface-variant my-2 font-sans text-xs leading-relaxed"
        >
          {renderBoldText(line)}
        </p>
      );
    });
  };

  // Helper to parse **bold** parts
  const renderBoldText = (text: string) => {
    const parts = text.split(/\*\*(.*?)\*\*/g);
    return parts.map((part, i) => {
      if (i % 2 === 1) {
        return (
          <strong key={i} className="text-on-surface font-bold">
            {part}
          </strong>
        );
      }
      return part;
    });
  };

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/75 p-4 backdrop-blur-md select-none">
      <div className="bg-surface-high border-outline-variant/30 flex max-h-[85vh] w-full max-w-2xl flex-col overflow-hidden rounded-2xl border shadow-2xl">
        {/* Header */}
        <div className="border-outline-variant/15 bg-surface-lowest flex items-center justify-between border-b p-5">
          <div className="flex items-center gap-2">
            <Sparkles size={18} className="animate-pulse text-[#4F7CFF]" />
            <div>
              <h3 className="text-on-surface font-display text-lg font-bold">
                AI Reasoning Framework
              </h3>
              <p className="text-on-surface-variant mt-1 text-xs font-medium">
                Detailed reasoning, hardware compatibility analysis, and risk
                assessment
              </p>
            </div>
          </div>
          <button
            onClick={onClose}
            className="text-on-surface-variant hover:text-on-surface hover:bg-surface-highest cursor-pointer rounded-lg p-1.5 transition-colors"
          >
            <X size={18} />
          </button>
        </div>

        {/* Content */}
        <div className="bg-surface-lowest/40 flex-1 scrollbar-thin overflow-auto p-6">
          <div className="prose prose-invert max-w-none">
            {renderReasoning(reasoning)}
          </div>
        </div>

        {/* Footer */}
        <div className="border-outline-variant/15 bg-surface-lowest flex justify-end border-t p-4">
          <button
            onClick={onClose}
            className="cursor-pointer rounded-lg bg-[#4F7CFF] px-6 py-2.5 font-mono text-xs font-bold text-white shadow-lg transition-colors hover:bg-[#4F7CFF]/90"
          >
            Close
          </button>
        </div>
      </div>
    </div>
  );
}
