import { useRef } from "react";
import { useGSAP } from "@gsap/react";
import { gsap } from "@/lib/gsap-config";
import { useReducedMotion } from "@/features/landing/hooks/useReducedMotion";
import { ShieldCheck, ShieldAlert, CheckSquare } from "lucide-react";

export function AiDecisionPanel() {
  const panelRef = useRef<HTMLDivElement>(null);
  const prefersReducedMotion = useReducedMotion();

  const products = [
    {
      name: "ASUS TUF A15",
      score: "96%",
      width: "96%",
      color: "bg-[var(--landing-gradient-accent)]",
      matches: "12/12 matched",
      stock: "In Stock",
      promo: "10% Off",
      promoColor: "text-emerald-400",
      hasAlert: false,
    },
    {
      name: "Lenovo LOQ",
      score: "88%",
      width: "88%",
      color: "bg-indigo-400",
      matches: "10/12 matched",
      stock: "In Stock",
      promo: "Free Gift",
      promoColor: "text-emerald-400",
      hasAlert: false,
    },
    {
      name: "Acer Nitro V",
      score: "74%",
      width: "74%",
      color: "bg-indigo-500",
      matches: "8/12 matched",
      stock: "1 Left",
      promo: "No Promo",
      promoColor: "text-amber-400",
      hasAlert: true,
    },
  ];

  useGSAP(
    () => {
      if (prefersReducedMotion) {
        return;
      }

      const tl = gsap.timeline({
        scrollTrigger: {
          trigger: panelRef.current,
          start: "top 75%",
          toggleActions: "play none none none",
        },
      });

      tl.from(".gsap-panel-title", { opacity: 0, y: 15, duration: 0.4 })
        .from(
          ".gsap-panel-input",
          { opacity: 0, y: 15, duration: 0.4 },
          "-=0.2"
        )
        .from(".gsap-panel-loading", { opacity: 0, duration: 0.3 }, "-=0.1")
        .from(
          ".gsap-vendor-row",
          { opacity: 0, y: 15, stagger: 0.15, duration: 0.5 },
          "-=0.1"
        )
        .from(
          ".gsap-score-bar",
          {
            scaleX: 0,
            transformOrigin: "left center",
            stagger: 0.15,
            duration: 0.6,
            ease: "power2.out",
          },
          "-=0.45"
        )
        .from(
          ".gsap-vendor-meta",
          { opacity: 0, y: 5, stagger: 0.15, duration: 0.4 },
          "-=0.3"
        );
    },
    { scope: panelRef }
  );

  return (
    <div
      ref={panelRef}
      aria-hidden="true"
      className="relative w-full max-w-md rounded-xl border border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] p-6 shadow-2xl backdrop-blur-md"
    >
      {/* Decorative background glow behind the panel */}
      <div className="pointer-events-none absolute -inset-2 -z-10 rounded-2xl bg-gradient-to-r from-[var(--landing-gradient-accent)]/10 to-indigo-500/10 opacity-50 blur-xl" />

      {/* Header bar simulating AI console */}
      <div className="gsap-panel-title border-border/20 mb-4 flex items-center justify-between border-b pb-4 select-none">
        <div className="flex items-center space-x-2">
          <div className="flex space-x-1">
            <span className="h-2.5 w-2.5 rounded-full bg-rose-500/80" />
            <span className="h-2.5 w-2.5 rounded-full bg-amber-500/80" />
            <span className="h-2.5 w-2.5 rounded-full bg-emerald-500/80" />
          </div>
          <span className="text-muted-foreground/60 font-heading text-[10px] font-bold tracking-wider uppercase">
            ShopWise AI Shopping Assistant
          </span>
        </div>
        <span className="text-muted-foreground/40 font-mono text-[10px]">
          v2.1.0
        </span>
      </div>

      {/* Query Search Input mockup */}
      <div className="gsap-panel-input bg-background/50 border-border/20 mb-4 flex items-center justify-between rounded-lg border p-3">
        <span className="text-foreground font-heading text-xs font-medium">
          Compare: Gaming laptop under 30 million VND
        </span>
        <div className="gsap-panel-loading flex items-center space-x-2">
          <style>{`
            @keyframes pulse-ring {
              0% { transform: scale(0.95); opacity: 0.5; }
              50% { transform: scale(1.15); opacity: 1; }
              100% { transform: scale(0.95); opacity: 0.5; }
            }
            .pulse-ring-active {
              animation: pulse-ring 2s infinite ease-in-out;
            }
            @media (prefers-reduced-motion: reduce) {
              .pulse-ring-active {
                animation: none !important;
              }
            }
          `}</style>
          <span className="pulse-ring-active h-2 w-2 rounded-full bg-emerald-500" />
          <span className="font-heading text-[9px] font-semibold tracking-wider text-emerald-400 uppercase">
            Live Analysis
          </span>
        </div>
      </div>

      {/* Scored Results List */}
      <div className="space-y-5">
        {products.map((p) => (
          <div
            key={p.name}
            className="gsap-vendor-row flex flex-col space-y-2 border-b border-border/10 pb-4 last:border-0 last:pb-0"
          >
            {/* Header: Name and Score */}
            <div className="flex items-center justify-between text-xs">
              <span className="text-foreground font-semibold">
                {p.name}
              </span>
              <span className="text-muted-foreground font-mono font-semibold">
                {p.score}
              </span>
            </div>

            {/* Score Bar */}
            <div className="bg-muted/20 h-2 w-full overflow-hidden rounded-full">
              <div
                style={{ width: p.width }}
                className={`gsap-score-bar ${p.color} h-full rounded-full`}
              />
            </div>

            {/* Sub-row: Metadata (Matches, Stock, Promo) */}
            <div className="gsap-vendor-meta flex items-center justify-between text-[10px] text-muted-foreground/80">
              <span className="flex items-center gap-1">
                <CheckSquare className="h-3 w-3 text-indigo-400 shrink-0" />
                {p.matches}
              </span>
              <span className="flex items-center gap-1">
                {p.hasAlert ? (
                  <ShieldAlert className="h-3 w-3 text-amber-400 shrink-0" />
                ) : (
                  <ShieldCheck className="h-3 w-3 text-emerald-400 shrink-0" />
                )}
                {p.stock}
              </span>
              <span className={`font-semibold ${p.promoColor}`}>
                {p.promo}
              </span>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
