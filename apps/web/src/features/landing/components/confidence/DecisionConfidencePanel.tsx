import { useRef } from "react";
import { useGSAP } from "@gsap/react";
import { gsap } from "@/lib/gsap-config";
import { useReducedMotion } from "@/features/landing/hooks/useReducedMotion";
import { Database, ShieldCheck, Tag, Cpu } from "lucide-react";

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
      icon: Database,
      title: "Live Stock",
      value: "100% In-Stock",
      subText: "Dynamic inventory queries directly from warehouse catalog",
      color: "text-emerald-400",
      glow: "shadow-[0_0_15px_rgba(52,211,153,0.1)]",
    },
    {
      icon: ShieldCheck,
      title: "Latest Pricing",
      value: "Verified Price",
      subText: "Accurate pricing matching active e-commerce catalogs",
      color: "text-[var(--landing-gradient-accent)]",
      glow: "shadow-[0_0_15px_var(--landing-glow)]",
    },
    {
      icon: Cpu,
      title: "Product Specs",
      value: "Verified Specs",
      subText: "Technically vetted CPU, RAM, GPU, and screen sizes",
      color: "text-emerald-400",
      glow: "shadow-[0_0_15px_rgba(52,211,153,0.1)]",
    },
    {
      icon: Tag,
      title: "Active Promotions",
      value: "Active Promos",
      subText: "Automated checks for coupon codes and campaign deals",
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
            className={`gsap-conf-card rounded-xl border border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] p-5 backdrop-blur-md transition-all duration-300 hover:border-foreground/20 ${m.glow}`}
          >
            <div className="flex items-center space-x-3">
              <div className={`rounded-lg border border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] p-2.5 ${m.color}`}>
                <Icon className="h-5 w-5" />
              </div>
              <h4 className="text-foreground/80 font-heading text-xs font-semibold uppercase tracking-wider">
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
