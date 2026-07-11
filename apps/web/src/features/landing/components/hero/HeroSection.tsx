import { useRef } from "react";
import { Link } from "@tanstack/react-router";
import { Badge } from "@shopwise/ui/components/badge";
import { Button } from "@shopwise/ui/components/button";
import { useGSAP } from "@gsap/react";
import { gsap } from "@/lib/gsap-config";
import { useReducedMotion } from "@/features/landing/hooks/useReducedMotion";
import { HeroVisual } from "./HeroVisual";

export function HeroSection() {
  const containerRef = useRef<HTMLDivElement>(null);
  const prefersReducedMotion = useReducedMotion();

  useGSAP(
    () => {
      if (prefersReducedMotion) {
        return;
      }

      const timeline = gsap.timeline({ defaults: { ease: "power2.out" } });

      timeline
        .from(".gsap-eyebrow", { opacity: 0, y: 20, duration: 0.5 })
        .from(".gsap-headline", { opacity: 0, y: 24, duration: 0.6 }, "-=0.3")
        .from(".gsap-subtitle", { opacity: 0, y: 20, duration: 0.5 }, "-=0.4")
        .from(".gsap-actions", { opacity: 0, y: 16, duration: 0.5 }, "-=0.4")
        .from(
          ".gsap-visual",
          { opacity: 0, scale: 0.95, duration: 0.7, ease: "power1.out" },
          "-=0.5"
        );
    },
    { scope: containerRef }
  );

  return (
    <section
      id="hero"
      ref={containerRef}
      style={{
        background:
          "linear-gradient(180deg, var(--landing-gradient-start) 0%, var(--landing-gradient-mid) 100%)",
      }}
      className="relative flex min-h-screen items-center overflow-hidden pt-28 pb-16 lg:pt-36"
    >
      {/* Background Grid Pattern Overlay */}
      <div className="pointer-events-none absolute inset-0 bg-[linear-gradient(to_right,oklch(100%_0_0deg_/_3%)_1px,transparent_1px),linear-gradient(to_bottom,oklch(100%_0_0deg_/_3%)_1px,transparent_1px)] [mask-image:radial-gradient(ellipse_60%_50%_at_50%_0%,#000_70%,transparent_100%)] bg-[size:4rem_4rem]" />

      <div className="relative z-10 container mx-auto px-4 md:px-6">
        <div className="grid grid-cols-1 items-center gap-12 lg:grid-cols-12 lg:gap-8">
          {/* Left Column: Content (55% space roughly on lg: screen) */}
          <div className="flex flex-col items-start space-y-6 lg:col-span-7">
            <div className="gsap-animate gsap-eyebrow">
              <Badge
                variant="outline"
                className="border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] px-3 py-1 text-[11px] backdrop-blur-xs select-none"
              >
                AI Decision Intelligence for Enterprise Procurement
              </Badge>
            </div>

            <h1 className="gsap-animate gsap-headline font-heading text-foreground max-w-xl text-4xl leading-[1.1] font-bold tracking-tight sm:text-5xl lg:text-6xl">
              From Fragmented Research to{" "}
              <span className="bg-gradient-to-r from-[var(--landing-gradient-accent)] to-indigo-300 bg-clip-text text-transparent">
                Confident Decisions.
              </span>
            </h1>

            <p className="gsap-animate gsap-subtitle text-muted-foreground max-w-lg text-base leading-relaxed sm:text-lg">
              ShopWise unifies supplier comparison, technical specification
              evaluation, and explainable AI reasoning into a single workspace
              for procurement teams.
            </p>

            <div className="gsap-animate gsap-actions flex flex-wrap gap-4 pt-2">
              <Button
                asChild
                size="lg"
                className="shadow-[var(--landing-glow)] shadow-lg transition-all hover:shadow-[var(--landing-glow)] hover:shadow-xl"
              >
                <Link to="/auth">Launch Decision Workspace</Link>
              </Button>
              <Button
                asChild
                variant="outline"
                size="lg"
                className="hover:bg-muted/40 border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] transition-colors"
              >
                <a href="#problem">See How ShopWise Reasons</a>
              </Button>
            </div>
          </div>

          {/* Right Column: Visual (45% space roughly) */}
          <div className="gsap-animate gsap-visual flex justify-center lg:col-span-5 lg:justify-end">
            <HeroVisual />
          </div>
        </div>
      </div>
    </section>
  );
}
