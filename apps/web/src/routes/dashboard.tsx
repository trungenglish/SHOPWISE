import { createFileRoute } from "@tanstack/react-router";
import React, { useState, useEffect } from "react";
import { useMutation } from "@tanstack/react-query";
import Sidebar from "@/features/dashboard/components/Sidebar";
import AuditTrail from "@/features/dashboard/components/AuditTrail";
import AgentHub from "@/features/dashboard/components/AgentHub";
import SpatialWorkspace from "@/features/dashboard/components/SpatialWorkspace";
import ComparisonModal from "@/features/dashboard/components/ComparisonModal";
import AccessoriesModal from "@/features/dashboard/components/AccessoriesModal";
import ReasoningModal from "@/features/dashboard/components/ReasoningModal";
import RetailModal from "@/features/dashboard/components/RetailModal";
import CheckoutModal from "@/features/dashboard/components/CheckoutModal";
import PriceAlertModal from "@/features/dashboard/components/PriceAlertModal";
import {
  Laptop,
  AuditLog,
  AgentStatus,
  TrustFactors,
  PriceAlert,
} from "@/features/dashboard/types";
import {
  initialProducts,
  initialLogs,
  initialAgents,
  initialReasoning,
  initialAccessories,
} from "@/features/dashboard/mock-data";
import { Sparkles, Loader2, Bookmark, Bell } from "lucide-react";

export const Route = createFileRoute("/dashboard")({
  component: DashboardPage,
});

// A simulated API response for TanStack Query mutation
const mockDecisionApi = async (queryText: string) => {
  return new Promise<any>((resolve) => {
    setTimeout(() => {
      resolve({
        sessionTitle: "AI Custom Decision",
        products: initialProducts.map((p) => ({
          ...p,
          matchScore: Math.floor(Math.random() * 15) + 84,
          matchExplanation: `Recommended for: "${queryText}". Demonstrates solid real-world performance.`,
        })),
        logs: [
          {
            time: "01:14",
            message: "Parsed dynamic requirements for " + queryText,
            status: "done",
          },
          {
            time: "01:15",
            message: "Generated customized feature alignment nodes",
            status: "done",
          },
          {
            time: "01:16",
            message: "Evaluated thermal cooling benchmarks for workspace",
            status: "done",
          },
          { time: "01:16", message: "Ready with results", status: "done" },
        ],
        agents: initialAgents,
        trustScore: 95,
        trustFactors: {
          benchmarkSources: 10,
          reviewCoverage: 90,
          retailConsensus: "Extremely well matched",
          confidenceEvolution: [30, 45, 60, 80, 95],
        },
        graphNodes: ["High Power", "Optimization"],
        reasoningExplanation: initialReasoning,
        accessories: initialAccessories,
      });
    }, 2500); // simulate 2.5s network delay
  });
};

function DashboardPage() {
  const [activeTab, setActiveTab] = useState<string>("sessions");
  const [userIntent, setUserIntent] = useState<string>(
    "Intensive 3D rendering and local AI development. Focus: cooling and GPU."
  );
  const [sessionTitle, setSessionTitle] =
    useState<string>("Hardware Render OS");
  const [products, setProducts] = useState<Laptop[]>(initialProducts);
  const [logs, setLogs] = useState<AuditLog[]>(initialLogs);
  const [agents, setAgents] = useState<AgentStatus[]>(initialAgents);
  const [trustScore, setTrustScore] = useState<number>(94);
  const [trustFactors, setTrustFactors] = useState<TrustFactors>({
    benchmarkSources: 8,
    reviewCoverage: 85,
    retailConsensus: "Highly recommended",
    confidenceEvolution: [20, 40, 50, 75, 94],
  });
  const [graphNodes, setGraphNodes] = useState<string[]>([
    "GPU Priority",
    "Thermals",
  ]);
  const [reasoning, setReasoning] = useState<string>(initialReasoning);
  const [accessories, setAccessories] = useState<any[]>(initialAccessories);

  // Active highlighted product card
  const [activeProductId, setActiveProductId] = useState<string>("macbook-pro");
  const [savedIds, setSavedIds] = useState<string[]>([]);

  // Price Alerts state
  const [priceAlerts, setPriceAlerts] = useState<PriceAlert[]>([]);
  const [isPriceAlertOpen, setIsPriceAlertOpen] = useState(false);
  const [selectedAlertLaptop, setSelectedAlertLaptop] = useState<Laptop | null>(
    null
  );
  const [toastNotification, setToastNotification] = useState<string | null>(
    null
  );

  // Open Price Alert Dialog
  const handleOpenPriceAlert = (laptop: Laptop) => {
    setSelectedAlertLaptop(laptop);
    setIsPriceAlertOpen(true);
  };

  // Save/Update alert
  const handleSaveAlert = (productId: string, targetPrice: number) => {
    setPriceAlerts((prev) => {
      const existingIdx = prev.findIndex((a) => a.productId === productId);
      if (existingIdx > -1) {
        return prev.map((a, idx) =>
          idx === existingIdx
            ? {
                ...a,
                targetPrice,
                active: true,
                createdAt: new Date().toISOString(),
              }
            : a
        );
      }
      return [
        ...prev,
        {
          productId,
          targetPrice,
          active: true,
          createdAt: new Date().toISOString(),
        },
      ];
    });

    const timeNow = new Date().toLocaleTimeString("vi-VN", {
      hour: "2-digit",
      minute: "2-digit",
    });
    const product = products.find((p) => p.id === productId);
    if (product) {
      setLogs((prev) => [
        {
          time: timeNow,
          message: `🔔 Price alert set for ${product.name} at $${targetPrice.toLocaleString()}`,
          status: "done",
        },
        ...prev,
      ]);
    }
  };

  // Delete/Cancel alert
  const handleDeleteAlert = (productId: string) => {
    setPriceAlerts((prev) => prev.filter((a) => a.productId !== productId));

    const timeNow = new Date().toLocaleTimeString("vi-VN", {
      hour: "2-digit",
      minute: "2-digit",
    });
    const product = products.find((p) => p.id === productId);
    if (product) {
      setLogs((prev) => [
        {
          time: timeNow,
          message: `🔕 Cancelled price alert for ${product.name}`,
          status: "done",
        },
        ...prev,
      ]);
    }
  };

  // Retail accounts state
  const [retailAccounts, setRetailAccounts] = useState([
    {
      id: "amazon",
      name: "Amazon US Premium",
      connected: false,
      discountRate: 0.05,
    },
    {
      id: "bestbuy",
      name: "Best Buy Elite",
      connected: false,
      discountRate: 0.03,
    },
    {
      id: "apple",
      name: "Apple Corporate Store",
      connected: false,
      discountRate: 0.08,
    },
  ]);

  // Modals visibility
  const [isCompareOpen, setIsCompareOpen] = useState(false);
  const [isAccessoriesOpen, setIsAccessoriesOpen] = useState(false);
  const [isReasoningOpen, setIsReasoningOpen] = useState(false);
  const [isRetailOpen, setIsRetailOpen] = useState(false);
  const [isCheckoutOpen, setIsCheckoutOpen] = useState(false);

  // Auto-connect to retail tab actions
  useEffect(() => {
    if (activeTab === "retail") {
      setIsRetailOpen(true);
      setActiveTab("sessions");
    }
  }, [activeTab]);

  // Calculate current active retail discount rate
  const connectedDiscount = retailAccounts
    .filter((a) => a.connected)
    .reduce((max, curr) => Math.max(max, curr.discountRate), 0);

  // Trigger notifications when price alert targets are reached
  useEffect(() => {
    priceAlerts.forEach((alert) => {
      if (!alert.active) return;

      const laptop = products.find((p) => p.id === alert.productId);
      if (!laptop) return;

      const currentDiscountedPrice = Math.round(
        laptop.price * (1 - connectedDiscount)
      );

      if (currentDiscountedPrice <= alert.targetPrice) {
        setToastNotification(
          `Price for ${laptop.name} dropped to $${currentDiscountedPrice.toLocaleString()} (reached target $${alert.targetPrice.toLocaleString()}!)`
        );

        const timeNow = new Date().toLocaleTimeString("vi-VN", {
          hour: "2-digit",
          minute: "2-digit",
        });
        setLogs((prevLogs) => {
          const logMsg = `🔔 PRICE ALERT REACHED: ${laptop.name} is $${currentDiscountedPrice.toLocaleString()} (Target: $${alert.targetPrice.toLocaleString()})`;
          if (prevLogs.some((l) => l.message === logMsg)) return prevLogs;
          return [
            {
              time: timeNow,
              message: logMsg,
              status: "done",
            },
            ...prevLogs,
          ];
        });
      }
    });
  }, [retailAccounts, priceAlerts, products, connectedDiscount]);

  const handleToggleSave = (id: string, e: React.MouseEvent) => {
    e.stopPropagation();
    if (savedIds.includes(id)) {
      setSavedIds(savedIds.filter((item) => item !== id));
    } else {
      setSavedIds([...savedIds, id]);
    }
  };

  const decisionMutation = useMutation({
    mutationFn: mockDecisionApi,
    onMutate: (queryText) => {
      setUserIntent(queryText);
      setLogs([
        {
          time: "Now",
          message: "Parsing new hardware requirements...",
          status: "running",
        },
        {
          time: "Pending",
          message: "Mapping ideal graphics specifications...",
          status: "pending",
        },
        {
          time: "Pending",
          message: "Connecting to Gemini data center for cross-checking...",
          status: "pending",
        },
      ]);
    },
    onSuccess: (data) => {
      setSessionTitle(data.sessionTitle);
      setProducts(data.products);
      setLogs(data.logs);
      setAgents(data.agents);
      setTrustScore(data.trustScore);
      setTrustFactors(data.trustFactors);
      setGraphNodes(data.graphNodes);
      setReasoning(data.reasoningExplanation);
      setAccessories(data.accessories);

      if (data.products && data.products.length > 0) {
        setActiveProductId(data.products[0].id);
      }
    },
    onError: (err, queryText) => {
      console.error("Decision mutation error", err);
    },
  });

  const handleQueryEvaluation = (queryText: string) => {
    decisionMutation.mutate(queryText);
  };

  const handleReplay = () => {
    handleQueryEvaluation(userIntent);
  };

  const handleToggleConnectRetail = (id: string, username?: string) => {
    setRetailAccounts(
      retailAccounts.map((acc) => {
        if (acc.id === id) {
          return {
            ...acc,
            connected: !acc.connected,
            username: username || "",
            loyaltyPoints: !acc.connected
              ? Math.floor(Math.random() * 800) + 200
              : undefined,
          };
        }
        return acc;
      })
    );
  };

  const handleNewSession = () => {
    handleQueryEvaluation(
      "High refresh rate OLED screen, thin & light, long battery for mobile coding."
    );
  };

  const activeProduct =
    products.find((p) => p.id === activeProductId) || products[0];

  return (
    <div className="relative flex h-screen w-full overflow-hidden bg-[#09090B]">
      {/* Background radial overlays */}
      <div className="shader-bg" />

      {/* 3-zone core workspace layout */}
      <Sidebar
        activeTab={activeTab}
        setActiveTab={setActiveTab}
        onNewSession={handleNewSession}
        savedCount={savedIds.length}
      />

      {/* Main Core Section */}
      <div className="flex h-full flex-1 overflow-hidden">
        {activeTab === "saved" ? (
          /* SAVED COLLECTIONS VIEW */
          <div className="z-10 flex flex-1 flex-col overflow-y-auto p-8 select-none">
            <div className="mb-6">
              <h2 className="text-on-surface flex items-center gap-2 text-2xl font-extrabold tracking-tight">
                <Bookmark size={24} className="text-[#4F7CFF]" />
                Saved Collection (Pinned hardware)
              </h2>
              <p className="text-on-surface-variant mt-1 text-xs font-medium">
                Manage optimal hardware products saved during decision sessions
              </p>
            </div>

            {savedIds.length === 0 ? (
              <div className="border-outline-variant/15 bg-surface-low/30 mx-auto flex w-full max-w-2xl flex-1 flex-col items-center justify-center rounded-2xl border border-dashed p-8 text-center">
                <Bookmark
                  size={48}
                  className="text-on-surface-variant/40 mb-3"
                />
                <h4 className="text-on-surface text-sm font-bold">
                  No products saved yet
                </h4>
                <p className="text-on-surface-variant mt-1.5 max-w-xs text-xs leading-relaxed">
                  Click any device in the Spatial Workspace and select "Save to
                  collection" to pin it here.
                </p>
              </div>
            ) : (
              <div className="grid max-w-5xl grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
                {products
                  .filter((p) => savedIds.includes(p.id))
                  .map((p) => (
                    <div
                      key={p.id}
                      className="glass-card border-outline-variant/15 relative flex flex-col gap-3 rounded-xl border p-4"
                    >
                      <img
                        src={p.image}
                        alt={p.name}
                        className="border-outline-variant/10 h-32 w-full rounded-lg border object-cover"
                      />
                      <div>
                        <h4 className="font-display text-on-surface text-sm font-bold">
                          {p.name}
                        </h4>
                        <p className="mt-1 font-mono text-xs font-semibold text-[#4F7CFF]">
                          ${p.price.toLocaleString("en-US")}
                        </p>
                      </div>
                      <p className="text-on-surface-variant font-sans text-xs leading-relaxed">
                        {p.matchExplanation}
                      </p>
                      <button
                        onClick={(e) => handleToggleSave(p.id, e)}
                        className="mt-auto cursor-pointer rounded border border-red-500/20 bg-red-500/10 px-3 py-1.5 font-mono text-[10px] text-red-400 transition-colors hover:bg-red-500/20"
                      >
                        Remove
                      </button>
                    </div>
                  ))}
              </div>
            )}
          </div>
        ) : (
          /* MAIN SPATIAL DECISION OS WORKSPACE */
          <>
            <AuditTrail
              logs={logs}
              userIntent={userIntent}
              onInjectConstraint={handleQueryEvaluation}
              isLoading={decisionMutation.isPending}
              onReplay={handleReplay}
            />

            <SpatialWorkspace
              products={products}
              activeProductId={activeProductId}
              setActiveProductId={setActiveProductId}
              savedIds={savedIds}
              onToggleSave={handleToggleSave}
              trustScore={trustScore}
              trustFactors={trustFactors}
              graphNodes={graphNodes}
              onChipClick={handleQueryEvaluation}
              onCompareAll={() => setIsCompareOpen(true)}
              onExplainReasoning={() => setIsReasoningOpen(true)}
              onAccessories={() => setIsAccessoriesOpen(true)}
              onCheckout={() => setIsCheckoutOpen(true)}
              discountRate={connectedDiscount}
              priceAlerts={priceAlerts}
              onOpenPriceAlert={handleOpenPriceAlert}
            />
          </>
        )}

        <AgentHub agents={agents} />
      </div>

      {/* IMMERSIVE LOADING SCREEN OVERLAY */}
      {decisionMutation.isPending && (
        <div className="fixed inset-0 z-[60] flex flex-col items-center justify-center bg-black/80 p-4 backdrop-blur-xl select-none">
          <div className="glass-card flex w-full max-w-sm flex-col items-center gap-4 rounded-2xl border border-[#4F7CFF]/30 p-8 text-center shadow-[0_0_50px_rgba(79,124,255,0.2)]">
            <div className="relative flex items-center justify-center">
              <Loader2 size={44} className="animate-spin text-[#4F7CFF]" />
              <Sparkles
                size={18}
                className="absolute animate-pulse text-[#4F7CFF]"
              />
            </div>
            <div>
              <h3 className="font-display text-on-surface text-sm font-black tracking-wide uppercase">
                SHOPWISE AI Decision Engine
              </h3>
              <p className="text-on-surface-variant mt-2 font-sans text-xs leading-relaxed">
                Simulating multi-agent collaboration flow, measuring thermal
                data & cross-checking optimal configurations...
              </p>
            </div>
            {/* Visual indicators */}
            <div className="bg-surface-lowest relative mt-2 h-1.5 w-full overflow-hidden rounded-full">
              <div className="h-full w-full animate-pulse bg-gradient-to-r from-[#4F7CFF] to-cyan-400" />
            </div>
          </div>
        </div>
      )}

      {/* Modals Mounting */}
      <ComparisonModal
        isOpen={isCompareOpen}
        onClose={() => setIsCompareOpen(false)}
        products={products}
      />

      <AccessoriesModal
        isOpen={isAccessoriesOpen}
        onClose={() => setIsAccessoriesOpen(false)}
        accessories={accessories}
      />

      <ReasoningModal
        isOpen={isReasoningOpen}
        onClose={() => setIsReasoningOpen(false)}
        reasoning={reasoning}
      />

      <RetailModal
        isOpen={isRetailOpen}
        onClose={() => setIsRetailOpen(false)}
        accounts={retailAccounts}
        onToggleConnect={handleToggleConnectRetail}
      />

      <CheckoutModal
        isOpen={isCheckoutOpen}
        onClose={() => setIsCheckoutOpen(false)}
        product={activeProduct}
        discountRate={connectedDiscount}
      />

      <PriceAlertModal
        isOpen={isPriceAlertOpen}
        onClose={() => setIsPriceAlertOpen(false)}
        product={selectedAlertLaptop}
        currentDiscountedPrice={
          selectedAlertLaptop
            ? Math.round(selectedAlertLaptop.price * (1 - connectedDiscount))
            : 0
        }
        activeAlert={
          selectedAlertLaptop
            ? priceAlerts.find((a) => a.productId === selectedAlertLaptop.id) ||
              null
            : null
        }
        onSaveAlert={handleSaveAlert}
        onDeleteAlert={handleDeleteAlert}
      />

      {/* Dynamic price alert toast notification */}
      {toastNotification && (
        <div className="fixed top-4 right-4 z-[70] w-full max-w-md select-none">
          <div className="bg-surface-highest/95 text-on-surface flex items-start gap-3 rounded-xl border-2 border-[#4F7CFF] p-4 shadow-[0_10px_30px_rgba(79,124,255,0.3)] backdrop-blur">
            <div className="shrink-0 animate-bounce rounded-lg bg-[#4F7CFF]/15 p-2 text-[#4F7CFF]">
              <Bell size={18} />
            </div>
            <div className="flex-1">
              <h4 className="font-display text-xs font-black tracking-wider text-[#4F7CFF] uppercase">
                PRICE ALERT REACHED
              </h4>
              <p className="text-on-surface mt-1 text-xs leading-relaxed font-medium">
                {toastNotification}
              </p>
            </div>
            <button
              onClick={() => setToastNotification(null)}
              className="text-on-surface-variant hover:text-on-surface hover:bg-surface-low shrink-0 cursor-pointer self-start rounded px-1.5 py-1 font-mono text-xs font-bold transition-colors"
            >
              Close
            </button>
          </div>
        </div>
      )}
    </div>
  );
}
