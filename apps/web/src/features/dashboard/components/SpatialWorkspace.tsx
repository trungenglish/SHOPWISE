import React, { useEffect, useState } from "react";
import {
  ZoomIn,
  ZoomOut,
  CheckCircle2,
  Star,
  Sparkles,
  AlertTriangle,
  GitCompare,
  HelpCircle,
} from "lucide-react";
import { Laptop, TrustFactors, PriceAlert } from "../types";
import {
  ReactFlow,
  ReactFlowProvider,
  useReactFlow,
  Background,
  useNodesState,
  useEdgesState,
  Node,
  Edge,
} from "@xyflow/react";
import "@xyflow/react/dist/style.css";
import { UserGoalNode, PriorityNode, LaptopNode } from "./FlowNodes";

const nodeTypes = {
  userGoal: UserGoalNode,
  priority: PriorityNode,
  laptopNode: LaptopNode,
};

interface SpatialWorkspaceProps {
  readonly products: Laptop[];
  readonly activeProductId: string;
  readonly setActiveProductId: (id: string) => void;
  readonly savedIds: string[];
  readonly onToggleSave: (id: string, e: React.MouseEvent) => void;
  readonly trustScore: number;
  readonly trustFactors: TrustFactors;
  readonly graphNodes: string[];
  readonly onChipClick: (suggestion: string) => void;
  readonly onCompareAll: () => void;
  readonly onExplainReasoning: () => void;
  readonly onAccessories: () => void;
  readonly onCheckout: () => void;
  readonly discountRate: number;
  readonly priceAlerts: PriceAlert[];
  readonly onOpenPriceAlert: (laptop: Laptop) => void;
}

const SpatialWorkspace = ({
  products,
  activeProductId,
  setActiveProductId,
  savedIds,
  onToggleSave,
  trustScore,
  trustFactors,
  graphNodes = ["GPU Priority", "Thermals"],
  onChipClick,
  onCompareAll,
  onExplainReasoning,
  onAccessories,
  onCheckout,
  discountRate,
  priceAlerts,
  onOpenPriceAlert,
}: SpatialWorkspaceProps) => {

  const { zoomIn, zoomOut, fitView } = useReactFlow();

  const [hoveredProductId, setHoveredProductId] = useState<string | null>(null);

  const [nodes, setNodes, onNodesChange] = useNodesState<Node>([]);
  const [edges, setEdges, onEdgesChange] = useEdgesState<Edge>([]);

  useEffect(() => {
    // Generate the Goal title based on intent/context
    const updatedNodes = [
      {
        id: "goal",
        type: "userGoal",
        position: { x: 30, y: 160 },
        data: { label: "Hardware Optimization" },
      },
      {
        id: "p1",
        type: "priority",
        position: { x: 260, y: 70 },
        data: { label: graphNodes[0] || "GPU Priority" },
      },
      {
        id: "p2",
        type: "priority",
        position: { x: 260, y: 250 },
        data: { label: graphNodes[1] || "Thermals" },
      },
      ...products.map((laptop, index) => ({
        id: laptop.id,
        type: "laptopNode",
        position: { x: 470, y: index * 170 },
        data: {
          laptop,
          isActive: laptop.id === activeProductId,
          discountRate,
          priceAlerts,
          onOpenPriceAlert,
          onSelect: () => setActiveProductId(laptop.id),
          onHover: () => setHoveredProductId(laptop.id),
          onHoverLeave: () => setHoveredProductId(null),
          onToggleSave,
          isSaved: savedIds.includes(laptop.id),
          isHovered: laptop.id === hoveredProductId,
        },
      })),
    ];

    const updatedEdges = [
      {
        id: "e-goal-p1",
        source: "goal",
        sourceHandle: "a",
        target: "p1",
        targetHandle: "in",
        animated: true,
        style: { stroke: "#4F7CFF", strokeWidth: 2 },
      },
      {
        id: "e-goal-p2",
        source: "goal",
        sourceHandle: "a",
        target: "p2",
        targetHandle: "in",
        animated: true,
        style: { stroke: "#4F7CFF", strokeWidth: 2 },
      },
      ...products.map((laptop) => {
        const isL1Active = laptop.id === activeProductId;
        return {
          id: `e-p1-${laptop.id}`,
          source: "p1",
          sourceHandle: "out",
          target: laptop.id,
          animated: isL1Active,
          style: {
            stroke: isL1Active ? "#4F7CFF" : "rgba(255, 255, 255, 0.15)",
            strokeWidth: isL1Active ? 2 : 1,
          },
        };
      }),
      ...products.map((laptop) => {
        const isL2Active = laptop.id === activeProductId;
        return {
          id: `e-p2-${laptop.id}`,
          source: "p2",
          sourceHandle: "out",
          target: laptop.id,
          animated: isL2Active,
          style: {
            stroke: isL2Active ? "#4F7CFF" : "rgba(255, 255, 255, 0.15)",
            strokeWidth: isL2Active ? 2 : 1,
          },
        };
      }),
    ];

    setNodes(updatedNodes);
    setEdges(updatedEdges);
  }, [products, activeProductId, discountRate, priceAlerts, graphNodes, setActiveProductId, onOpenPriceAlert, setNodes, setEdges, onToggleSave, savedIds, hoveredProductId]);

  const activeProduct =
    products.find((p) => p.id === activeProductId) || products[0];
  const displayProduct = hoveredProductId
    ? products.find((p) => p.id === hoveredProductId)
    : activeProduct;
  const isSaved = displayProduct ? savedIds.includes(displayProduct.id) : false;

  return (
    <main className="relative flex h-full flex-1 flex-col overflow-hidden bg-transparent p-4 select-none">
      {/* Title block */}
      <div className="relative z-10 mb-4 flex items-start justify-between">
        <div>
          <h2 className="font-display text-on-surface flex items-center gap-2 text-2xl font-extrabold tracking-tight">
            Spatial Workspace
          </h2>
          <p className="text-on-surface-variant mt-1 text-xs font-medium">
            Visual hardware mapping space & compatible configuration comparison
          </p>
        </div>

        {/* Zoom controls */}
        {/* <div className="bg-surface-lowest/60 border-outline-variant/15 flex gap-1.5 rounded-lg border p-1 backdrop-blur">
          <button
            onClick={() => zoomIn()}
            className="bg-surface-low hover:bg-surface-highest text-on-surface-variant hover:text-primary cursor-pointer rounded p-1.5 transition-colors"
            title="Zoom In"
          >
            <ZoomIn size={15} />
          </button>
          <button
            onClick={() => zoomOut()}
            className="bg-surface-low hover:bg-surface-highest text-on-surface-variant hover:text-primary cursor-pointer rounded p-1.5 transition-colors"
            title="Zoom Out"
          >
            <ZoomOut size={15} />
          </button>
          <button
            onClick={() => fitView()}
            className="bg-surface-low hover:bg-surface-highest text-on-surface-variant hover:text-primary cursor-pointer rounded px-2 py-1 font-mono text-[10px] font-bold transition-colors"
            title="Fit View"
          >
            FIT
          </button>
        </div> */}
      </div>

      {/* Main Canvas Area */}
      <div className="relative h-full min-h-[480px] w-full flex-1 overflow-hidden rounded-2xl">
        {/* React Flow Editor */}
        <ReactFlow
          nodes={nodes}
          edges={edges}
          onNodesChange={onNodesChange}
          onEdgesChange={onEdgesChange}
          nodeTypes={nodeTypes}
          fitView
          className="animate-fade-in bg-transparent"
          minZoom={0.2}
          maxZoom={1.5}
          proOptions={{ hideAttribution: true }}
        >
          <Background color="rgba(79, 124, 255, 0.08)" gap={16} size={1} />
        </ReactFlow>

        {/* Contextual Specs Side Panel */}
        {displayProduct ? <div className="glass-card border-outline-variant/15 absolute top-4 left-4 z-20 flex w-[300px] flex-col gap-4 rounded-2xl border p-5 shadow-[0_15px_40px_-10px_rgba(79,124,255,0.15)] transition-all">
            <div className="border-outline-variant/15 flex items-start justify-between border-b pb-3">
              <span className="font-display text-[11px] font-black tracking-wider text-[#4F7CFF]">
                {hoveredProductId ? "PREVIEWING DETAILS" : "SELECTED PRODUCT"}
              </span>
              <span className="font-mono text-sm font-bold text-emerald-400">
                {(displayProduct.price * (1 - discountRate)).toLocaleString(
                  "vi-VN"
                )} ₫
              </span>
            </div>

            <h3 className="text-on-surface text-lg leading-tight font-bold">
              {displayProduct.name}
            </h3>

            <div className="text-on-surface-variant mt-2 flex flex-col gap-1.5 font-mono text-xs">
              <div className="border-outline-variant/5 flex justify-between border-b pb-1">
                <span className="opacity-70">CPU:</span>
                <span className="max-w-[180px] truncate text-right text-white">
                  {displayProduct.specs?.cpu || "Standard"}
                </span>
              </div>
              <div className="border-outline-variant/5 flex justify-between border-b pb-1">
                <span className="opacity-70">RAM:</span>
                <span className="max-w-[180px] truncate text-right text-white">
                  {displayProduct.specs?.ram || "16GB"}
                </span>
              </div>
              <div className="border-outline-variant/5 flex justify-between border-b pb-1">
                <span className="opacity-70">GPU:</span>
                <span className="max-w-[180px] truncate text-right text-white">
                  {displayProduct.specs?.gpu || "Integrated"}
                </span>
              </div>
              <div className="border-outline-variant/5 flex justify-between border-b pb-1">
                <span className="opacity-70">Screen:</span>
                <span className="max-w-[180px] truncate text-right text-white">
                  {displayProduct.specs?.screen || "Standard"}
                </span>
              </div>
              <div className="border-outline-variant/5 flex justify-between border-b pb-1">
                <span className="opacity-70">Cooling:</span>
                <span className="max-w-[180px] truncate text-right text-white">
                  {displayProduct.specs?.cooling || "Standard"}
                </span>
              </div>
              <div className="flex justify-between pb-1">
                <span className="opacity-70">Warranty:</span>
                <span className="max-w-[180px] truncate text-right text-white">
                  {displayProduct.specs?.warranty || "Standard 1-Year"}
                </span>
              </div>
            </div>

            {/* Match Parameter Bars */}
            <div className="border-outline-variant/15 mt-1 flex flex-col gap-2.5 border-t pt-3">
              <span className="text-on-surface-variant mb-1 text-[10px] font-bold tracking-wider uppercase">
                Performance Match
              </span>
              <div className="flex items-center justify-between text-[10px]">
                <span className="text-on-surface-variant w-16">AI Perf</span>
                <div className="bg-surface-low mx-2 h-1.5 flex-1 overflow-hidden rounded-full">
                  <div
                    className="h-full rounded-full bg-gradient-to-r from-[#4F7CFF] to-indigo-400 transition-all duration-500"
                    style={{ width: `${displayProduct.aiPerf}%` }}
                  />
                </div>
                <span className="w-7 text-right font-mono text-white">
                  {displayProduct.aiPerf}%
                </span>
              </div>
              <div className="flex items-center justify-between text-[10px]">
                <span className="text-on-surface-variant w-16">3D Render</span>
                <div className="bg-surface-low mx-2 h-1.5 flex-1 overflow-hidden rounded-full">
                  <div
                    className="h-full rounded-full bg-gradient-to-r from-[#4F7CFF] to-indigo-400 transition-all duration-500"
                    style={{ width: `${displayProduct.rendering}%` }}
                  />
                </div>
                <span className="w-7 text-right font-mono text-white">
                  {displayProduct.rendering}%
                </span>
              </div>
              <div className="flex items-center justify-between text-[10px]">
                <span className="text-on-surface-variant w-16">Thermals</span>
                <div className="bg-surface-low mx-2 h-1.5 flex-1 overflow-hidden rounded-full">
                  <div
                    className="h-full rounded-full bg-gradient-to-r from-[#4F7CFF] to-indigo-400 transition-all duration-500"
                    style={{ width: `${displayProduct.thermals}%` }}
                  />
                </div>
                <span className="w-7 text-right font-mono text-white">
                  {displayProduct.thermals}%
                </span>
              </div>
            </div>

            {/* Pin action */}
            <button
              onClick={(e) => onToggleSave(displayProduct.id, e)}
              className="mt-4 flex w-full cursor-pointer items-center justify-center gap-1.5 rounded-lg border border-[#4F7CFF]/30 bg-[#4F7CFF]/15 px-3 py-2 text-[11px] font-bold text-[#4F7CFF] transition-colors hover:bg-[#4F7CFF]/25"
            >
              {isSaved ? "Unpin from Board" : "Pin to Board"}
            </button>
          </div> : null}


      </div>

      {/* Persistent Floating Action Dock */}
      <div className="bg-surface-highest/80 border-outline-variant/40 z-20 mx-auto mb-2 flex w-full max-w-2xl items-center justify-between gap-5 rounded-full border px-6 py-3 shadow-[0_10px_30px_rgba(0,0,0,0.5)] backdrop-blur-xl">
        <button
          onClick={onCompareAll}
          className="text-on-surface flex cursor-pointer items-center gap-2 font-mono text-[12px] font-bold transition-colors hover:text-[#4F7CFF]"
        >
          <GitCompare size={15} /> Compare All
        </button>
        <button
          onClick={onExplainReasoning}
          className="text-on-surface flex cursor-pointer items-center gap-2 font-mono text-[12px] font-bold transition-colors hover:text-[#4F7CFF]"
        >
          <Sparkles size={15} className="text-[#4F7CFF]" /> Explain Reasoning
        </button>
        <button
          onClick={onAccessories}
          className="text-on-surface flex cursor-pointer items-center gap-2 font-mono text-[12px] font-bold transition-colors hover:text-[#4F7CFF]"
        >
          <HelpCircle size={15} /> Compatible Accessories
        </button>
        <button
          onClick={onCheckout}
          className="cursor-pointer rounded-full bg-[#4F7CFF] px-5 py-2 font-mono text-[12px] font-black text-white shadow-[0_0_15px_rgba(79,124,255,0.4)] transition-all hover:bg-[#4F7CFF]/90"
        >
          Checkout
        </button>
      </div>
    </main>
  );
}

export default SpatialWorkspace;
