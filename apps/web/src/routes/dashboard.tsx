import { createFileRoute } from "@tanstack/react-router";
import React, { useState, useEffect } from "react";
import { ReactFlowProvider } from "@xyflow/react";
import {
  type AgentEnvelope,
  runAgentTurnStream,
  sendInteractionStream,
} from "@/api/decision-memory";
import type { AccessoryRecommendation } from "@/api/accessories";
import DynamicUIRenderer from "@/components/dynamic-uirenderer";
import type { InteractionRequest } from "@shopwise/protocols";
import Sidebar from "@/features/dashboard/components/sidebar";
import AuditTrail from "@/features/dashboard/components/audit-trail";
import AgentHub from "@/features/dashboard/components/agent-hub";
import TrustCenterHeader from "@/features/dashboard/components/trust-center-header";
import SpatialWorkspace from "@/features/dashboard/components/spatial-workspace";
import ComparisonModal from "@/features/dashboard/components/comparison-modal";
import AccessoriesModal from "@/features/dashboard/components/accessories-modal";
import ReasoningModal from "@/features/dashboard/components/reasoning-modal";
import RetailModal from "@/features/dashboard/components/retail-modal";
import CheckoutModal from "@/features/dashboard/components/checkout-modal";
import PriceAlertModal from "@/features/dashboard/components/price-alert-modal";
import { OrderHistoryPage } from "@/features/orders/components/order-history-page";
import {
  Laptop,
  AuditLog,
  AgentStatus,
  TrustFactors,
  PriceAlert,
} from "@/features/dashboard/types";
import { mapRecommendationProduct } from "@/features/dashboard/agent-mapping";
import { Bookmark, Bell, Menu } from "lucide-react";
import EmptyWorkspace from "@/features/dashboard/components/empty-workspace";

export const Route = createFileRoute("/dashboard")({
  component: DashboardPage,
});

function DashboardPage() {
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);
  const [isInitialState, setIsInitialState] = useState(true);
  const [isSimulating, setIsSimulating] = useState(false);
  const [sessionId, setSessionId] = useState<string | null>(null);
  const [activeEnvelope, setActiveEnvelope] = useState<AgentEnvelope | null>(
    null
  );
  const [activeTab, setActiveTab] = useState<string>("sessions");
  const [userIntent, setUserIntent] = useState<string>("");
  const [products, setProducts] = useState<Laptop[]>([]);
  const [logs, setLogs] = useState<AuditLog[]>([]);
  const [agents, setAgents] = useState<AgentStatus[]>([]);
  const trustScore = 0;
  const trustFactors = {} as TrustFactors;
  const [graphNodes, setGraphNodes] = useState<string[]>([]);
  const [reasoning, setReasoning] = useState<string>("");
  const [checkoutAccessories, setCheckoutAccessories] = useState<
    AccessoryRecommendation[]
  >([]);

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

  const applyAgentEnvelope = (envelope: AgentEnvelope, time: string) => {
    setActiveEnvelope(envelope);
    setLogs((previous) => [
      ...previous,
      { time, message: `Agent: ${envelope.message}`, status: "done" },
    ]);
    if (
      envelope.type !== "recommendation" &&
      envelope.type !== "comparison" &&
      envelope.type !== "checkout_ready"
    ) {
      return;
    }
    const recommendedProducts = envelope.decision.products.map(
      mapRecommendationProduct
    );
    setCheckoutAccessories([]);
    setProducts(recommendedProducts);
    setReasoning(envelope.decision.reasoning);
    setAgents([
      {
        id: "shopping-agent",
        name: "Shopping Agent",
        progress: 100,
        statusMessage: "Recommendation generated from the live catalog",
        type: "shopping",
      },
    ]);
    setGraphNodes(["Catalog recommendation"]);
    setActiveProductId(recommendedProducts[0]?.id ?? "");
    if (envelope.type === "comparison") {
      setIsCompareOpen(true);
    }
    if (envelope.type === "checkout_ready") {
      setIsCheckoutOpen(true);
    }
  };

  const handleQueryEvaluation = async (queryText: string) => {
    const trimmedMessage = queryText.trim();
    if (isSimulating || trimmedMessage.length === 0) {
      return;
    }

    const timeNow = new Date().toLocaleTimeString("vi-VN", {
      hour: "2-digit",
      minute: "2-digit",
    });
    setIsInitialState(false);
    setIsSimulating(true);
    setUserIntent(trimmedMessage);
    setLogs((previous) => [
      ...previous,
      { time: timeNow, message: `User: ${trimmedMessage}`, status: "done" },
    ]);

    try {
      const turn = await runAgentTurnStream(sessionId, trimmedMessage);
      setSessionId(turn.sessionId);
      applyAgentEnvelope(turn.envelope, timeNow);
    } catch {
      const errorMessage = "Agent request failed. Please retry.";
      setToastNotification(errorMessage);
      setLogs((previous) => [
        ...previous,
        { time: timeNow, message: errorMessage, status: "pending" },
      ]);
    } finally {
      setIsSimulating(false);
    }
  };

  const handleInteraction = async (request: InteractionRequest) => {
    if (sessionId === null || isSimulating) {
      return;
    }
    const timeNow = new Date().toLocaleTimeString("vi-VN", {
      hour: "2-digit",
      minute: "2-digit",
    });
    setIsSimulating(true);
    try {
      const envelope = await sendInteractionStream(sessionId, request);
      applyAgentEnvelope(envelope, timeNow);
    } catch {
      setToastNotification("Agent request failed. Please retry.");
    } finally {
      setIsSimulating(false);
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
    setSessionId(null);
    setActiveEnvelope(null);
    setIsInitialState(true);
    setUserIntent("");
    setProducts([]);
    setLogs([]);
    setAgents([]);
    setGraphNodes([]);
    setReasoning("");
    setCheckoutAccessories([]);
    setActiveProductId("");
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
            className="border-outline-variant/20 bg-surface-high text-on-surface-variant hover:bg-surface-highest hover:text-on-surface absolute top-6 left-5 z-50 flex h-10 w-10 cursor-pointer items-center justify-center rounded-xl border shadow-lg transition-colors"
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
                  Manage optimal hardware products saved during decision
                  sessions
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
                    Click any device in the Spatial Workspace and select "Save
                    to collection" to pin it here.
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
            <OrderHistoryPage
              onContinueShopping={() => setActiveTab("sessions")}
            />
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
                  hasMenuButton={!isSidebarOpen}
                  hideComposer={activeEnvelope?.type === "question"}
                  suggestions={[
                    "Ngân sách dưới 25 triệu VND",
                    "Ưu tiên hiệu năng và tản nhiệt",
                    "Ưu tiên mỏng nhẹ và pin lâu",
                  ]}
                />
              )}

              {isInitialState ? (
                <EmptyWorkspace onSubmit={handleQueryEvaluation} />
              ) : activeEnvelope?.type === "question" &&
                products.length === 0 ? (
                <main className="flex flex-1 items-center justify-center overflow-y-auto p-8">
                  <DynamicUIRenderer
                    envelope={activeEnvelope}
                    onInteraction={handleInteraction}
                  />
                </main>
              ) : products.length === 0 ? (
                <main className="flex flex-1 items-center justify-center p-8 text-center text-white/60">
                  Tell the agent your constraints to build a recommendation
                  workspace.
                </main>
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
          products={products}
          onAddToCheckout={(accessory) => {
            setCheckoutAccessories((current) =>
              current.some(
                (item) => item.retailer_offer_id === accessory.retailer_offer_id
              )
                ? current
                : [...current, accessory]
            );
            setToastNotification(
              `${accessory.name} đã được thêm vào checkout.`
            );
          }}
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
          accessories={checkoutAccessories}
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
              ? priceAlerts.find(
                  (a) => a.productId === selectedAlertLaptop.id
                ) || null
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
