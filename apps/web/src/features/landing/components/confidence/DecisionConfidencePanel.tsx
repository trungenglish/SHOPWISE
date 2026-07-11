import { useRef } from "react";
import { useGSAP } from "@gsap/react";
import { gsap } from "@/lib/gsap-config";
import { useReducedMotion } from "@/features/landing/hooks/useReducedMotion";
import { Check, Info, ArrowRight } from "lucide-react";

export function DecisionConfidencePanel() {
  const panelRef = useRef<HTMLDivElement>(null);
  const prefersReducedMotion = useReducedMotion();

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

      tl.from(".gsap-rec-card", {
        opacity: 0,
        scale: 0.95,
        y: 20,
        duration: 0.6,
        ease: "power2.out",
      })
        .from(
          ".gsap-rec-item",
          { opacity: 0, y: 10, stagger: 0.1, duration: 0.4 },
          "-=0.3"
        );
    },
    { scope: panelRef }
  );

  const whyPoints = [
    "RTX 4060 GPU",
    "Better cooling system",
    "24GB RAM config",
    "In Stock (Phong Vu Hanoi)",
    "Eligible for 10% Coupon Promo",
    "Best value in price range",
  ];

  return (
    <div
      ref={panelRef}
      aria-hidden="true"
      className="gsap-rec-card relative w-full max-w-md rounded-xl border border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] p-6 shadow-2xl backdrop-blur-md text-left select-none"
    >
      {/* Decorative background glow behind the panel */}
      <div className="pointer-events-none absolute -inset-2 -z-10 rounded-2xl bg-gradient-to-r from-[var(--landing-gradient-accent)]/10 to-indigo-500/10 opacity-50 blur-xl" />

      {/* Section 1: Customer Request */}
      <div className="gsap-rec-item mb-5 border-b border-border/10 pb-4">
        <span className="text-muted-foreground/60 font-heading text-[9px] font-bold tracking-wider uppercase block mb-2.5">
          Customer Request
        </span>
        <div className="flex flex-wrap gap-2">
          <span className="rounded-full border border-border/10 bg-muted/5 px-2.5 py-1 text-[11px] font-medium text-foreground">
            Gaming laptop
          </span>
          <span className="rounded-full border border-blue-500/20 bg-blue-500/10 px-2.5 py-1 text-[11px] font-medium text-blue-400">
            RTX 4060
          </span>
          <span className="rounded-full border border-border/10 bg-muted/5 px-2.5 py-1 text-[11px] font-medium text-foreground">
            Under 30M VND
          </span>
        </div>
      </div>

      {/* Section 2: Recommended Product */}
      <div className="gsap-rec-item mb-5 border-b border-border/10 pb-4">
        <span className="text-muted-foreground/60 font-heading text-[9px] font-bold tracking-wider uppercase block mb-1.5">
          Recommended Product
        </span>
        <div className="flex items-center justify-between">
          <h4 className="text-foreground text-lg font-black tracking-tight">
            Lenovo LOQ 15IRX9
          </h4>
          <span className="rounded bg-[var(--landing-gradient-accent)]/10 px-2 py-0.5 text-[10px] font-extrabold text-[var(--landing-gradient-accent)] border border-[var(--landing-gradient-accent)]/20 animate-pulse">
            Best Match
          </span>
        </div>
      </div>

      {/* Section 3: Why this recommendation */}
      <div className="gsap-rec-item mb-6">
        <span className="text-muted-foreground/60 font-heading text-[9px] font-bold tracking-wider uppercase block mb-3">
          Why this recommendation
        </span>
        <div className="grid grid-cols-2 gap-x-4 gap-y-2.5">
          {whyPoints.map((point, idx) => (
            <div key={idx} className="flex items-center space-x-2 text-xs">
              <div className="flex h-4.5 w-4.5 shrink-0 items-center justify-center rounded-full bg-emerald-500/10 text-emerald-400">
                <Check className="h-3 w-3" />
              </div>
              <span className="text-foreground/90 truncate">{point}</span>
            </div>
          ))}
        </div>
      </div>

      {/* Section 4: Actions */}
      <div className="gsap-rec-item flex flex-col gap-2.5 pt-2">
        <button
          type="button"
          tabIndex={-1}
          className="w-full rounded-lg bg-[var(--landing-gradient-accent)] hover:opacity-90 py-2.5 text-center text-xs font-bold text-background transition-opacity flex items-center justify-center gap-1.5 cursor-default"
        >
          View Details
          <ArrowRight className="h-3.5 w-3.5" />
        </button>
        <button
          type="button"
          tabIndex={-1}
          className="w-full rounded-lg border border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] hover:bg-muted/10 py-2.5 text-center text-xs font-semibold text-foreground transition-colors flex items-center justify-center gap-1.5 cursor-default"
        >
          <Info className="h-3.5 w-3.5" />
          Compare Similar Products
        </button>
      </div>
    </div>
  );
}
