import { useState } from "react";
import { X, Check, AlertCircle } from "lucide-react";

interface RetailAccount {
  id: string;
  name: string;
  connected: boolean;
  username?: string;
  loyaltyPoints?: number;
  discountRate: number; // e.g. 0.05 for 5% off
}

interface RetailModalProps {
  isOpen: boolean;
  onClose: () => void;
  accounts: RetailAccount[];
  onToggleConnect: (id: string, username?: string) => void;
}

export default function RetailModal({
  isOpen,
  onClose,
  accounts,
  onToggleConnect,
}: RetailModalProps) {
  const [inputs, setInputs] = useState<{ [key: string]: string }>({});

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/75 p-4 backdrop-blur-md select-none">
      <div className="bg-surface-high border-outline-variant/30 flex max-h-[85vh] w-full max-w-md flex-col overflow-hidden rounded-2xl border shadow-2xl">
        {/* Header */}
        <div className="border-outline-variant/15 bg-surface-lowest flex items-center justify-between border-b p-5">
          <div>
            <h3 className="text-on-surface font-display text-lg font-bold">
              Retailer Integration
            </h3>
            <p className="text-on-surface-variant mt-1 text-xs font-medium">
              Link retailer accounts to check inventory & apply exclusive
              discounts
            </p>
          </div>
          <button
            onClick={onClose}
            className="text-on-surface-variant hover:text-on-surface hover:bg-surface-highest cursor-pointer rounded-lg p-1.5 transition-colors"
          >
            <X size={18} />
          </button>
        </div>

        {/* Content */}
        <div className="flex flex-1 flex-col gap-4 overflow-auto p-5">
          <div className="bg-primary/10 border-primary/20 flex items-start gap-2.5 rounded-xl border p-3">
            <AlertCircle size={16} className="mt-0.5 shrink-0 text-[#4F7CFF]" />
            <p className="text-on-surface-variant font-sans text-[11px] leading-relaxed">
              <strong>Maximum Security:</strong> SHOPWISE encrypts integration
              information directly with providers using secure protocols. Data
              is only used for applying discount codes and automatic inventory
              reconciliation.
            </p>
          </div>

          <div className="flex flex-col gap-3">
            {accounts.map((acc) => {
              const username = inputs[acc.id] || "";

              return (
                <div
                  key={acc.id}
                  className={`rounded-xl border p-4 transition-[background-color,border-color] duration-200 ease-[var(--ease-out)] ${
                    acc.connected
                      ? "border-[#4F7CFF]/50 bg-[#4F7CFF]/5"
                      : "border-outline-variant/20 bg-surface-low"
                  }`}
                >
                  <div className="mb-2 flex items-start justify-between">
                    <div>
                      <h4 className="font-display text-on-surface text-sm font-bold">
                        {acc.name}
                      </h4>
                      <p className="font-mono text-[10px] font-medium text-[#4F7CFF]">
                        Exclusive discount: {acc.discountRate * 100}%
                      </p>
                    </div>

                    {acc.connected ? (
                      <span className="flex items-center gap-1 rounded border border-green-500/20 bg-green-500/10 px-2 py-0.5 font-mono text-[9px] font-bold text-green-400">
                        <Check size={10} /> CONNECTED
                      </span>
                    ) : (
                      <span className="bg-surface-highest text-on-surface-variant rounded px-2 py-0.5 font-mono text-[9px] font-semibold">
                        NOT CONNECTED
                      </span>
                    )}
                  </div>

                  {acc.connected ? (
                    <div className="bg-surface-lowest/50 border-outline-variant/10 mt-3 flex items-center justify-between rounded-lg border p-2 font-mono text-xs">
                      <span className="text-on-surface-variant">
                        User: {acc.username}
                      </span>
                      <span className="font-bold text-yellow-400">
                        {acc.loyaltyPoints} Points
                      </span>
                      <button
                        onClick={() => onToggleConnect(acc.id)}
                        className="cursor-pointer text-[10px] font-bold text-red-400 uppercase hover:text-red-300"
                      >
                        Disconnect
                      </button>
                    </div>
                  ) : (
                    <div className="mt-3 flex gap-2">
                      <input
                        value={username}
                        onChange={(e) =>
                          setInputs({ ...inputs, [acc.id]: e.target.value })
                        }
                        placeholder="Enter Email or Username..."
                        className="bg-surface-lowest border-outline-variant/30 text-on-surface placeholder:text-on-surface-variant/40 flex-1 rounded-lg border px-2.5 py-1.5 text-xs outline-none focus:border-[#4F7CFF]/60 focus:ring-0"
                      />
                      <button
                        onClick={() => {
                          if (username.trim()) {
                            onToggleConnect(acc.id, username.trim());
                            setInputs({ ...inputs, [acc.id]: "" });
                          }
                        }}
                        className="shrink-0 cursor-pointer rounded-lg bg-[#4F7CFF] px-4 font-mono text-[11px] font-bold text-white transition-colors hover:bg-[#4F7CFF]/90"
                      >
                        Connect
                      </button>
                    </div>
                  )}
                </div>
              );
            })}
          </div>
        </div>

        {/* Footer */}
        <div className="border-outline-variant/15 bg-surface-lowest flex justify-end border-t p-4">
          <button
            onClick={onClose}
            className="cursor-pointer rounded-lg bg-[#4F7CFF] px-6 py-2.5 font-mono text-xs font-bold text-white shadow-lg transition-colors hover:bg-[#4F7CFF]/90"
          >
            Done
          </button>
        </div>
      </div>
    </div>
  );
}
