import {
  Plus,
  MessageSquare,
  Bookmark,
  UserCheck,
  Settings,
  User,
} from "lucide-react";

interface SidebarProps {
  activeTab: string;
  setActiveTab: (tab: string) => void;
  onNewSession: () => void;
  savedCount: number;
  onClose?: () => void;
}

export default function Sidebar({
  activeTab,
  setActiveTab,
  onNewSession,
  savedCount,
  onClose,
}: SidebarProps) {
  return (
    <aside
      className="glass-panel flex h-full shrink-0 flex-col px-4 py-6 select-none"
      style={{ width: "260px" }}
    >
      {/* Brand Logo & Close Button */}
      <div className="mt-2 mb-8 flex items-center justify-between px-2">
        <div>
          <h1 className="font-display text-gradient text-[26px] leading-tight font-black tracking-tight">
            SHOPWISE
          </h1>
          <p className="font-display text-on-surface-variant text-[11px] font-medium tracking-widest uppercase opacity-80">
            AI Decision OS
          </p>
        </div>
        {onClose && (
          <button
            onClick={onClose}
            className="text-on-surface-variant hover:text-on-surface hover:bg-surface-high rounded p-1 transition-colors"
          >
            <Plus size={20} className="rotate-45" />
          </button>
        )}
      </div>

      {/* New Session Action */}
      <button
        onClick={onNewSession}
        className="bg-surface text-primary font-display mx-1 mb-6 flex cursor-pointer items-center justify-center gap-2 rounded-lg border border-[#4F7CFF]/30 px-4 py-3 text-xs font-semibold shadow-[0_0_15px_rgba(79,124,255,0.08)] transition-all hover:scale-[1.02] hover:border-[#4F7CFF]/60 hover:bg-[#4F7CFF]/15 active:scale-[0.98]"
      >
        <Plus size={16} />
        New Session
      </button>

      {/* Main Navigation */}
      <nav className="flex flex-1 flex-col gap-1.5">
        <button
          onClick={() => setActiveTab("sessions")}
          className={`flex cursor-pointer items-center gap-3 rounded-lg px-4 py-3 text-left transition-all ${
            activeTab === "sessions"
              ? "border-r-2 border-[#4F7CFF] bg-[#4F7CFF]/5 font-bold text-[#4F7CFF]"
              : "text-on-surface-variant hover:bg-surface-high hover:text-on-surface"
          }`}
        >
          <MessageSquare
            size={18}
            className={activeTab === "sessions" ? "text-[#4F7CFF]" : ""}
          />
          <span className="font-sans text-sm font-medium">
            Decision Session
          </span>
        </button>

        <button
          onClick={() => setActiveTab("saved")}
          className={`flex cursor-pointer items-center gap-3 rounded-lg px-4 py-3 text-left transition-all ${
            activeTab === "saved"
              ? "border-r-2 border-[#4F7CFF] bg-[#4F7CFF]/5 font-bold text-[#4F7CFF]"
              : "text-on-surface-variant hover:bg-surface-high hover:text-on-surface"
          }`}
        >
          <Bookmark
            size={18}
            className={activeTab === "saved" ? "text-[#4F7CFF]" : ""}
          />
          <span className="flex-1 font-sans text-sm font-medium">
            Saved Collection
          </span>
          {savedCount > 0 && (
            <span className="rounded-full bg-[#4F7CFF] px-2 py-0.5 font-mono text-[10px] font-bold text-white">
              {savedCount}
            </span>
          )}
        </button>

        <button
          onClick={() => setActiveTab("retail")}
          className={`flex cursor-pointer items-center gap-3 rounded-lg px-4 py-3 text-left transition-all ${
            activeTab === "retail"
              ? "border-r-2 border-[#4F7CFF] bg-[#4F7CFF]/5 font-bold text-[#4F7CFF]"
              : "text-on-surface-variant hover:bg-surface-high hover:text-on-surface"
          }`}
        >
          <UserCheck
            size={18}
            className={activeTab === "retail" ? "text-[#4F7CFF]" : ""}
          />
          <span className="font-sans text-sm font-medium">Retail Account</span>
        </button>
      </nav>

      {/* Bottom Profile and Settings */}
      <div className="border-outline-variant/20 mt-auto border-t pt-4">
        <a
          href="#settings"
          className="text-on-surface-variant hover:bg-surface-high hover:text-on-surface mb-2 flex items-center gap-3 rounded-lg px-4 py-2.5 transition-all"
        >
          <Settings size={18} />
          <span className="font-sans text-sm font-medium">Settings</span>
        </a>
        <div className="flex items-center gap-3 px-4 py-2">
          <div className="bg-surface-high border-outline-variant/30 flex h-8 w-8 items-center justify-center rounded-full border">
            <User size={16} className="text-on-surface-variant" />
          </div>
          <div className="text-on-surface-variant truncate font-sans text-xs font-semibold">
            quannguyenminh1704...
          </div>
        </div>
      </div>
    </aside>
  );
}
