import { useRef } from "react";
import { useGSAP } from "@gsap/react";
import { gsap } from "@/lib/gsap-config";
import { useReducedMotion } from "@/features/landing/hooks/useReducedMotion";
import { FEATURES } from "../../data/features";
import { FeatureCard } from "./FeatureCard";

export function FeaturesSection() {
  const sectionRef = useRef<HTMLDivElement>(null);
  const prefersReducedMotion = useReducedMotion();

  useGSAP(
    () => {
      if (prefersReducedMotion) {
        return;
      }

      gsap.from(".gsap-feature-reveal", {
        scrollTrigger: {
          trigger: sectionRef.current,
          start: "top 80%",
          toggleActions: "play none none none",
        },
        opacity: 0,
        y: 30,
        stagger: 0.08,
        duration: 0.6,
        ease: "power2.out",
      });
    },
    { scope: sectionRef }
  );

  return (
    <section
      id="features"
      ref={sectionRef}
      className="bg-background border-border/20 border-b py-20 lg:py-32"
    >
      <div className="container mx-auto px-4 md:px-6">
        {/* Header Block */}
        <div className="mb-12 flex flex-col items-center space-y-4 text-center lg:mb-20">
          <span className="gsap-feature-reveal font-heading text-xs font-semibold tracking-wider text-[var(--landing-gradient-accent)] uppercase">
            Meet the Workspace
          </span>
          <h2 className="gsap-feature-reveal font-heading text-foreground max-w-2xl text-3xl leading-[1.15] font-bold tracking-tight sm:text-4xl lg:text-5xl">
            A Unified Workspace for Procurement Decisions
          </h2>
          <p className="gsap-feature-reveal text-muted-foreground max-w-lg text-sm sm:text-base">
            ShopWise centralizes the tools required for enterprise product and
            supplier decisions, from initial criteria definition to final human
            sign-off.
          </p>
        </div>

        {/* Features Card Grid */}
        <ul className="gsap-feature-reveal grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
          {FEATURES.map((feature) => (
            <FeatureCard key={feature.id} item={feature} />
          ))}
        </ul>
      </div>
    </section>
  );
}
