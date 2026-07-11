import { useRef } from "react";
import { useGSAP } from "@gsap/react";
import { gsap } from "@/lib/gsap-config";
import { useReducedMotion } from "@/features/landing/hooks/useReducedMotion";
import { ShieldCheck, TrendingUp, History, UserCheck } from "lucide-react";

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

      tl.from(".gsap-conf-card", {
        opacity: 0,
        scale: 0.95,
        y: 20,
        stagger: 0.1,
        duration: 0.5,
        ease: "power2.out",
      });
    },
    { scope: panelRef }
  );

  const metrics = [
    {
      icon: ShieldCheck,
      title: "Evidence Status",
      value: "100% Verified",
      subText: "All specification sheets & quotes matching source claims",
      color: "text-emerald-400",
      glow: "shadow-[0_0_15px_rgba(52,211,153,0.1)]",
    },
    {
      icon: TrendingUp,
      title: "Confidence Level",
      value: "High Confidence (96%)",
      subText: "Scored against cost, compliance, and spec criteria",
      color: "text-[var(--landing-gradient-accent)]",
      glow: "shadow-[0_0_15px_var(--landing-glow)]",
    },
    {
      icon: UserCheck,
      title: "Risk Status",
      value: "Low Risk",
      subText: "No financial anomalies or supply chain warnings",
      color: "text-emerald-400",
      glow: "shadow-[0_0_15px_rgba(52,211,153,0.1)]",
    },
    {
      icon: History,
      title: "Audit Trail",
      value: "Traceable History",
      subText: "Decisions version-controlled with immutable logs",
      color: "text-indigo-400",
      glow: "shadow-[0_0_15px_rgba(129,140,248,0.1)]",
    },
  ];

  return (
    <div
      ref={panelRef}
      aria-hidden="true"
      className="grid w-full grid-cols-1 gap-4 sm:grid-cols-2 lg:max-w-lg"
    >
      {metrics.map((m, idx) => {
        const Icon = m.icon;
        return (
          <div
            key={idx}
            className={`gsap-conf-card hover:border-foreground/20 rounded-xl border border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] p-5 backdrop-blur-md transition-all duration-300 ${m.glow}`}
          >
            <div className="flex items-center space-x-3">
              <div
                className={`rounded-lg border border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] p-2.5 ${m.color}`}
              >
                <Icon className="h-5 w-5" />
              </div>
              <h4 className="text-foreground/80 font-heading text-xs font-semibold tracking-wider uppercase">
                {m.title}
              </h4>
            </div>
            <div className="mt-4">
              <span className="text-foreground text-lg font-bold sm:text-xl">
                {m.value}
              </span>
              <p className="text-muted-foreground mt-1 text-[11px] leading-normal">
                {m.subText}
              </p>
            </div>
          </div>
        );
      })}
    </div>
  );
}
