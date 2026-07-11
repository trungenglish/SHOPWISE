import { Cpu, Truck } from "lucide-react";
import { useRef } from "react";
import { useGSAP } from "@gsap/react";
import { gsap } from "@/lib/gsap-config";
import { useReducedMotion } from "@/features/landing/hooks/useReducedMotion";

export function HeroVisual() {
  const containerRef = useRef<HTMLDivElement>(null);
  const prefersReducedMotion = useReducedMotion();

  useGSAP(
    () => {
      if (prefersReducedMotion) {
        return;
      }

      const tl = gsap.timeline({ defaults: { ease: "power2.out" } });

      tl.from(".gsap-radar", { scale: 0.8, opacity: 0, duration: 0.8 })
        .from(".gsap-floating-node", { opacity: 0, y: 15, stagger: 0.15, duration: 0.5 }, "-=0.3")
        .from(".gsap-audit-trail", { opacity: 0, x: 20, duration: 0.6 }, "-=0.4");
    },
    { scope: containerRef }
  );

  return (
    <div
      ref={containerRef}
      aria-hidden="true"
      className="relative flex h-[350px] sm:h-[400px] lg:h-[500px] w-full items-center justify-center overflow-hidden"
    >
      <style>{`
        @keyframes float-slow {
          0%, 100% { transform: translateY(0px); }
          50% { transform: translateY(-10px); }
        }
        @keyframes float-reverse {
          0%, 100% { transform: translateY(0px); }
          50% { transform: translateY(8px); }
        }
        @keyframes sonar-pulse {
          0% { transform: scale(0.9); opacity: 0.2; }
          50% { transform: scale(1.1); opacity: 0.4; }
          100% { transform: scale(0.9); opacity: 0.2; }
        }
        .anim-float {
          animation: float-slow 6s ease-in-out infinite;
        }
        .anim-float-delayed {
          animation: float-reverse 7s ease-in-out infinite;
        }
        .anim-sonar {
          animation: sonar-pulse 4s ease-in-out infinite;
        }
        @media (prefers-reduced-motion: reduce) {
          .anim-float, .anim-float-delayed, .anim-sonar {
            animation: none !important;
            transform: none !important;
          }
        }
      `}</style>

      {/* Grid Canvas Background (Thin lines overlay) */}
      <div className="absolute inset-0 bg-[linear-gradient(to_right,oklch(100%_0_0deg_/_3%)_1px,transparent_1px),linear-gradient(to_bottom,oklch(100%_0_0deg_/_3%)_1px,transparent_1px)] bg-[size:1.5rem_1.5rem] opacity-70" />

      {/* Background glow shadow */}
      <div className="pointer-events-none absolute h-[300px] w-[300px] rounded-full bg-blue-500/10 blur-3xl" />

      {/* Dynamic Sonar / Radar Overlay */}
      <div className="gsap-radar relative flex h-52 w-52 items-center justify-center">
        {/* Pulsing Outer Ring */}
        <div className="anim-sonar absolute inset-0 rounded-full border border-blue-500/15" />
        <div className="absolute h-40 w-40 rounded-full border border-blue-500/10" />

        {/* Central Core Circle */}
        <div className="relative z-10 flex h-32 w-32 flex-col items-center justify-center rounded-full border border-blue-500/30 bg-background/80 shadow-2xl backdrop-blur-xs">
          <span className="text-foreground text-3xl font-extrabold tracking-tight">94%</span>
          <span className="text-muted-foreground mt-0.5 text-[9px] font-bold tracking-widest uppercase">
            Confidence
          </span>
          <div className="mt-2.5 h-1 w-14 rounded-full bg-blue-500/20">
            <div className="h-full w-[94%] rounded-full bg-blue-500" />
          </div>
        </div>
      </div>

      {/* Top Left Node: LAPTOPS Verified */}
      <div className="gsap-floating-node anim-float absolute top-12 left-6 sm:left-12 z-20 flex items-center space-x-2.5 rounded-full border border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] px-3.5 py-2 shadow-lg backdrop-blur-md">
        <div className="flex h-7 w-7 items-center justify-center rounded-full bg-blue-500/10 text-blue-400">
          <Cpu className="h-4 w-4" />
        </div>
        <div className="flex flex-col text-left">
          <span className="text-foreground font-heading text-[10px] font-bold tracking-wider uppercase">
            Laptops
          </span>
          <span className="text-muted-foreground text-[9px]">Verified</span>
        </div>
      </div>

      {/* Top Right Node: LIVE STOCK Verified */}
      <div className="gsap-floating-node anim-float-delayed absolute top-8 right-6 sm:right-16 z-20 flex items-center space-x-2.5 rounded-full border border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] px-3.5 py-2 shadow-lg backdrop-blur-md">
        <div className="flex h-7 w-7 items-center justify-center rounded-full bg-emerald-500/10 text-emerald-400">
          <Truck className="h-4 w-4" />
        </div>
        <div className="flex flex-col text-left">
          <span className="text-foreground font-heading text-[10px] font-bold tracking-wider uppercase">
            Live Stock
          </span>
          <span className="text-muted-foreground text-[9px]">Verified</span>
        </div>
      </div>

      {/* Bottom Left Card: ASUS TUF A15 */}
      <div className="gsap-floating-node anim-float-delayed absolute bottom-10 left-6 sm:left-10 z-20 flex w-36 flex-col justify-between rounded-xl border border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] p-3.5 shadow-xl backdrop-blur-md">
        <div className="text-left">
          <div className="text-foreground text-xs font-bold leading-none">ASUS TUF A15</div>
          <div className="mt-2 flex items-center gap-1.5">
            <span className="h-1.5 w-1.5 rounded-full bg-emerald-500 animate-pulse" />
            <span className="text-emerald-400 text-[9px] font-semibold">Optimizing...</span>
          </div>
        </div>
      </div>

      {/* Bottom Right Card: Lenovo LOQ */}
      <div className="gsap-floating-node anim-float absolute bottom-12 right-28 sm:right-36 z-20 flex w-40 flex-col justify-between rounded-xl border border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] p-3.5 shadow-xl backdrop-blur-md">
        <div className="text-left">
          <div className="text-foreground text-xs font-bold leading-none">Lenovo LOQ</div>
          <div className="mt-2 flex items-center gap-1.5">
            <span className="h-1.5 w-1.5 rounded-full bg-blue-500" />
            <span className="text-blue-400 text-[9px] font-semibold">Priority Match</span>
          </div>
        </div>
      </div>

      {/* Floating Right Side Panel: LIVE AUDIT TRAIL */}
      <div className="gsap-audit-trail absolute right-2 top-24 z-10 hidden w-44 rounded-lg border border-[var(--landing-glass-border)] bg-background/70 p-3.5 shadow-xl backdrop-blur-xs sm:block">
        <div className="border-b border-border/10 mb-2 pb-1.5 text-left">
          <span className="text-muted-foreground/60 font-heading text-[8px] font-bold tracking-wider uppercase">
            Live Analysis
          </span>
        </div>
        <div className="space-y-1.5 text-left font-mono text-[8px] leading-tight text-blue-400/90">
          <p className="truncate">&gt; Customer request received.</p>
          <p className="truncate">&gt; Analyzing laptop criteria...</p>
          <p className="truncate">&gt; Retrieving Phong Vu stock...</p>
          <p className="truncate">&gt; Cross-referencing specs.</p>
          <p className="truncate">&gt; Promotions applied.</p>
          <p className="truncate text-blue-400 font-bold">&gt; ASUS TUF recommended (94%).</p>
        </div>
      </div>

      {/* Decorative SVG Dashed Connectors */}
      <svg
        viewBox="0 0 400 400"
        className="absolute inset-0 -z-10 h-full w-full stroke-blue-500/20"
        fill="none"
        strokeWidth="1"
        strokeDasharray="4 4"
      >
        {/* Connector to Top Left (HARDWARE SP) */}
        <path d="M120,80 L200,200" />
        {/* Connector to Top Right (LOGISTICS) */}
        <path d="M280,70 L200,200" />
        {/* Connector to Bottom Left (Option A) */}
        <path d="M100,320 L200,200" />
        {/* Connector to Bottom Right (Option B) */}
        <path d="M240,320 L200,200" />
      </svg>
    </div>
  );
}
