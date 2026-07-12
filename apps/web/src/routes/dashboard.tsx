import { createFileRoute } from "@tanstack/react-router";
import React, { useState, useEffect } from "react";
import { ReactFlowProvider } from "@xyflow/react";
import { useMutation } from "@tanstack/react-query";
import Sidebar from "@/features/dashboard/components/Sidebar";
import AuditTrail from "@/features/dashboard/components/AuditTrail";
import AgentHub from "@/features/dashboard/components/AgentHub";
import TrustCenterHeader from "@/features/dashboard/components/TrustCenterHeader";
import SpatialWorkspace from "@/features/dashboard/components/SpatialWorkspace";
import ComparisonModal from "@/features/dashboard/components/ComparisonModal";
import AccessoriesModal from "@/features/dashboard/components/AccessoriesModal";
import ReasoningModal from "@/features/dashboard/components/ReasoningModal";
import RetailModal from "@/features/dashboard/components/RetailModal";
import CheckoutModal from "@/features/dashboard/components/CheckoutModal";
import PriceAlertModal from "@/features/dashboard/components/PriceAlertModal";
import { OrderHistoryPage } from "@/features/orders/components/OrderHistoryPage";
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
import { Sparkles, Loader2, Bookmark, Bell, Menu } from "lucide-react";
import EmptyWorkspace from "@/features/dashboard/components/EmptyWorkspace";

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
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);
  const [isInitialState, setIsInitialState] = useState(true);
  const [isSimulating, setIsSimulating] = useState(false);
  const [isSimulationComplete, setIsSimulationComplete] = useState(false);
  const [conversationStep, setConversationStep] = useState<string>("initial");
  const [activeTab, setActiveTab] = useState<string>("sessions");
  const [userIntent, setUserIntent] = useState<string>("");
  const [sessionTitle, setSessionTitle] = useState<string>("");
  const [products, setProducts] = useState<Laptop[]>([]);
  const [logs, setLogs] = useState<AuditLog[]>([]);
  const [agents, setAgents] = useState<AgentStatus[]>([]);
  const [trustScore, setTrustScore] = useState<number>(0);
  const [trustFactors, setTrustFactors] = useState<TrustFactors>(
    {} as TrustFactors
  );
  const [graphNodes, setGraphNodes] = useState<string[]>([]);
  const [reasoning, setReasoning] = useState<string>("");
  const [accessories, setAccessories] = useState<any[]>([]);

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
          message: `🔔 Price alert set for ${product.name} at ${targetPrice.toLocaleString("vi-VN")} ₫`,
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
          `Price for ${laptop.name} dropped to ${currentDiscountedPrice.toLocaleString("vi-VN")} ₫ (reached target ${alert.targetPrice.toLocaleString("vi-VN")} ₫!)`
        );

        const timeNow = new Date().toLocaleTimeString("vi-VN", {
          hour: "2-digit",
          minute: "2-digit",
        });
        setLogs((prevLogs) => {
          const logMsg = `🔔 PRICE ALERT REACHED: ${laptop.name} is ${currentDiscountedPrice.toLocaleString("vi-VN")} ₫ (Target: ${alert.targetPrice.toLocaleString("vi-VN")} ₫)`;
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

  const handleToggleSave = (id: string, e?: React.MouseEvent) => {
    if (e && e.stopPropagation) {
      e.stopPropagation();
    }
    if (savedIds.includes(id)) {
      setSavedIds(savedIds.filter((item) => item !== id));
    } else {
      setSavedIds([...savedIds, id]);
    }
  };

  // const [isSimulating, setIsSimulating] = useState(false);

  const handleQueryEvaluation = (queryText: string) => {
    if (isSimulating) return;

    if (isInitialState) {
      setIsInitialState(false);
      setUserIntent(queryText);
      setConversationStep("asked_aaa");

      const timeNow = new Date().toLocaleTimeString("vi-VN", {
        hour: "2-digit",
        minute: "2-digit",
      });
      setLogs([
        {
          time: timeNow,
          message: `User: ${queryText}`,
          status: "done",
        },
        {
          time: timeNow,
          message: "Agent: Do you play AAA titles (e.g., Cyberpunk, Call of Duty)?",
          status: "done",
        },
      ]);
      return;
    }

    if (conversationStep === "asked_aaa") {
      const timeNow = new Date().toLocaleTimeString("vi-VN", {
        hour: "2-digit",
        minute: "2-digit",
      });
      setLogs((prev) => [
        ...prev,
        {
          time: timeNow,
          message: `User: ${queryText}`,
          status: "done",
        },
        {
          time: timeNow,
          message: "Agent: Are you planning to stream or play competitively for money?",
          status: "done",
        }
      ]);
      setConversationStep("asked_money");
      if (queryText.includes("Yes")) setUserIntent(prev => prev + " | Plays AAA");
      else setUserIntent(prev => prev + " | Esports/Indie");
      return;
    }

    if (conversationStep === "asked_money") {
      const timeNow = new Date().toLocaleTimeString("vi-VN", {
        hour: "2-digit",
        minute: "2-digit",
      });
      setLogs((prev) => [
        ...prev,
        {
          time: timeNow,
          message: `User: ${queryText}`,
          status: "done",
        },
        {
          time: timeNow,
          message: "Agent: What's your budget?",
          status: "done",
        }
      ]);
      setConversationStep("asked_budget");
      if (queryText.includes("Yes")) setUserIntent(prev => prev + " | Earning money");
      else setUserIntent(prev => prev + " | Just for fun");
      return;
    }

    if (conversationStep === "asked_budget") {
      const timeNow = new Date().toLocaleTimeString("vi-VN", {
        hour: "2-digit",
        minute: "2-digit",
      });
      setLogs((prev) => [
        ...prev,
        {
          time: timeNow,
          message: `User: ${queryText}`,
          status: "done",
        },
      ]);
      setConversationStep("completed");
      setUserIntent(prev => prev + ` | Budget: ${queryText}`);

      setIsSimulating(true);

      // Clear state for progressive load
      setAgents([]);
      setProducts([]);
      setTrustFactors({} as TrustFactors);
      setAccessories([]);
      setGraphNodes([]);

      // Step 1: Initial Parsing
      setLogs((prev) => [
        ...prev,
        {
          time: "Now",
          message: (
            <span className="flex items-center gap-2">
              <Loader2 size={14} className="animate-spin text-[#4F7CFF]" />
              Parsing final hardware requirements...
            </span>
          ),
          status: "running",
        },
      ]);

      // Step 2: Agent initialization
      setTimeout(() => {
        setAgents(initialAgents.map((a) => ({ ...a, progress: 10 })));
        setLogs((prev) => {
          const newLogs = [...prev];
          newLogs[newLogs.length - 1] = {
            ...newLogs[newLogs.length - 1],
            status: "done",
            message: "Parsed final hardware requirements",
          };
          return [
            ...newLogs,
            {
              time: "Now",
              message: (
                <span className="flex items-center gap-2">
                  <Loader2 size={14} className="animate-spin text-[#4F7CFF]" />
                  Mapping ideal graphics specifications...
                </span>
              ),
              status: "running",
            },
          ];
        });
        setGraphNodes(["GPU Priority", "Thermals"]);
      }, 4000);

      // Step 3: API cross-check
      setTimeout(() => {
        setAgents(initialAgents.map((a) => ({ ...a, progress: 60 })));
        setLogs((prev) => {
          const newLogs = [...prev];
          newLogs[newLogs.length - 1] = {
            ...newLogs[newLogs.length - 1],
            status: "done",
            message: "Mapped ideal graphics specifications",
          };
          return [
            ...newLogs,
            {
              time: "Now",
              message: (
                <span className="flex items-center gap-2">
                  <Loader2 size={14} className="animate-spin text-[#4F7CFF]" />
                  Connecting to retail data centers...
                </span>
              ),
              status: "running",
            },
          ];
        });
      }, 8000);

      // Step 4: Final Resolve
      setTimeout(() => {
        setSessionTitle("AI Custom Decision");

        // Customize products based on intent
        let finalProducts = [...initialProducts];
        if (queryText.includes("Yes")) {
          // If earning money, push high reliability options (just reordering mock data as an example)
          finalProducts = [initialProducts[0], initialProducts[2], initialProducts[1]];
        }

        setProducts(
          finalProducts.map((p) => ({
            ...p,
            matchScore: Math.floor(Math.random() * 15) + 84,
            matchExplanation: `Recommended based on your specific requirements for AAA titles and earning money.`,
          }))
        );

        setLogs((prev) => {
          const newLogs = [...prev];
          newLogs[newLogs.length - 1] = {
            ...newLogs[newLogs.length - 1],
            status: "done",
            message: "Connected to retail data centers",
          };
          return newLogs;
        });

        setAgents(initialAgents);
        setTrustScore(95);
        setTrustFactors({
          benchmarkSources: 10,
          reviewCoverage: 90,
          retailConsensus: "Extremely well matched",
          confidenceEvolution: [30, 45, 60, 80, 95],
        });
        setReasoning(initialReasoning);
        setAccessories(initialAccessories);

        if (finalProducts.length > 0) {
          setActiveProductId(finalProducts[0].id);
          setGraphNodes(["Completed"]);
        }
        setIsSimulating(false);
        setIsSimulationComplete(true);
      }, 12000);
    }
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
    products.find((p) => p.id === activeProductId) || products[0] || null;

  return (
    <ReactFlowProvider>
      <div className="relative flex h-screen w-full overflow-hidden bg-[#09090B]">
      {/* Background radial overlays */}
      <div className="shader-bg" />

      {/* Global Header */}
      <TrustCenterHeader
        trustScore={trustScore}
        trustFactors={trustFactors}
        isSimulating={isSimulating}
      />

      {/* Sidebar Toggle Button (if hidden) */}
      {!isSidebarOpen && (
        <button
          onClick={() => setIsSidebarOpen(true)}
          className="border-outline-variant/20 bg-surface-high text-on-surface-variant hover:bg-surface-highest hover:text-on-surface absolute top-4 left-4 z-50 flex h-10 w-10 cursor-pointer items-center justify-center rounded-xl border shadow-lg transition-colors"
        >
          <Menu size={20} />
        </button>
      )}

      {/* 3-zone core workspace layout */}
      {isSidebarOpen && (
        <Sidebar
          activeTab={activeTab}
          setActiveTab={setActiveTab}
          onNewSession={handleNewSession}
          savedCount={savedIds.length}
          onClose={() => setIsSidebarOpen(false)}
        />
      )}

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
                          {p.price.toLocaleString("vi-VN")} ₫
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
        ) : activeTab === "orders" ? (
          <OrderHistoryPage onContinueShopping={() => setActiveTab("sessions")} />
        ) : (
          /* MAIN SPATIAL DECISION OS WORKSPACE */
          <>
            {!isInitialState && (
              <AuditTrail
                logs={logs}
                userIntent={userIntent}
                onInjectConstraint={handleQueryEvaluation}
                isLoading={isSimulating}
                onReplay={handleReplay}
                suggestions={
                  conversationStep === "asked_aaa"
                    ? ["Yes", "No (esports/indie)"]
                    : conversationStep === "asked_money"
                      ? ["streaming/esports", "just for fun"]
                      : conversationStep === "asked_budget"
                        ? ["Under 25tr VND", "25tr - 30tr VND", "30tr - 40tr VND", "40tr - 50tr VND", "50tr - 60tr VND", "Over 60tr VND"]
                        : [
                          `Compare with ${products[1]?.name || "Razer Blade"}`,
                          `Cooling: Vapor Chamber vs Dual Fans`,
                          `Test AI Performance with TensorRT`,
                        ]
                }
              />
            )}

            {isInitialState ? (
              <EmptyWorkspace onSubmit={handleQueryEvaluation} />
            ) : (
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
            )}
          </>
        )}

        {!isInitialState && <AgentHub agents={agents} />}
      </div>

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
        onSuccess={() => setActiveTab("orders")}
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
    </ReactFlowProvider>
  );
}
