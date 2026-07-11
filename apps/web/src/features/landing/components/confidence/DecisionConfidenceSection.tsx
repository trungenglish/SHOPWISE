import { useRef } from "react";
import { Link } from "@tanstack/react-router";
import { Button } from "@shopwise/ui/components/button";
import { useGSAP } from "@gsap/react";
import { gsap } from "@/lib/gsap-config";
import { useReducedMotion } from "@/features/landing/hooks/useReducedMotion";
import { DecisionConfidencePanel } from "./DecisionConfidencePanel";

export function DecisionConfidenceSection() {
  const sectionRef = useRef<HTMLDivElement>(null);
  const prefersReducedMotion = useReducedMotion();

  useGSAP(
    () => {
      if (prefersReducedMotion) {
        return;
      }

      gsap.from(".gsap-confidence-reveal", {
        scrollTrigger: {
          trigger: sectionRef.current,
          start: "top 75%",
          toggleActions: "play none none none",
        },
        opacity: 0,
        y: 40,
        stagger: 0.15,
        duration: 0.6,
        ease: "power2.out",
      });
    },
    { scope: sectionRef }
  );

  return (
    <section
      id="recommendation"
      ref={sectionRef}
      className="bg-background border-border/20 relative border-b py-20 lg:py-32"
    >
      <div className="container mx-auto px-4 md:px-6">
        <div className="grid grid-cols-1 items-center gap-12 lg:grid-cols-12 lg:gap-16">
          {/* Left Column: Copy Section */}
          <div className="flex flex-col items-start space-y-6 lg:col-span-6">
            <span className="gsap-confidence-reveal font-heading text-xs font-semibold tracking-wider text-[var(--landing-gradient-accent)] uppercase">
              AI Shopping Outcome
            </span>

            <h2 className="gsap-confidence-reveal font-heading text-foreground max-w-xl text-3xl leading-[1.15] font-bold tracking-tight sm:text-4xl lg:text-5xl">
              Find the Right{" "}
              <span className="bg-gradient-to-r from-[var(--landing-gradient-accent)] to-indigo-300 bg-clip-text text-transparent">
                Product Faster.
              </span>
            </h2>

            <p className="gsap-confidence-reveal text-muted-foreground max-w-lg text-base leading-relaxed sm:text-lg">
              See how ShopWise combines product specifications, live inventory,
              pricing, promotions, and customer requirements into one
              explainable recommendation.
            </p>

            <div className="gsap-confidence-reveal flex flex-wrap gap-4 pt-2">
              <Button
                asChild
                size="lg"
                className="shadow-[var(--landing-glow)] shadow-lg transition-all hover:shadow-[var(--landing-glow)] hover:shadow-xl"
              >
                <Link to="/auth">Try AI Sales Agent</Link>
              </Button>
              <Button
                asChild
                variant="outline"
                size="lg"
                className="hover:bg-muted/40 border border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] transition-colors"
              >
                <a href="#workspace">See Product Comparison</a>
              </Button>
            </div>
          </div>

          {/* Right Column: Visual Mockup */}
          <div className="flex justify-center lg:col-span-6 lg:justify-end">
            <span className="sr-only">
              Simulated demonstration of the ShopWise AI shopping recommendation
              outcome: customer requests a gaming laptop with RTX 4060 under 30
              million VND. Lenovo LOQ 15IRX9 is recommended as the best match
              because it includes an RTX 4060 GPU, better cooling system, 24GB
              RAM, is in stock, and is eligible for a 10% coupon promo.
            </span>
            <DecisionConfidencePanel />
          </div>
        </div>
      </div>
    </section>
  );
}
