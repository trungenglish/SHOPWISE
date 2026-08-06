import { useRef } from "react";
import { Link } from "@tanstack/react-router";
import { Badge } from "@shopwise/ui/components/badge";
import { Button } from "@shopwise/ui/components/button";
import { useGSAP } from "@gsap/react";
import { gsap } from "@/lib/gsap-config";
import { useReducedMotion } from "@/features/landing/hooks/use-reduced-motion";
import { HeroVisual } from "./hero-visual";

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
        .from(".gsap-stats", { opacity: 0, y: 15, duration: 0.5 }, "-=0.4")
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
                className="flex items-center gap-1.5 border-blue-500/30 bg-blue-500/10 px-3 py-1 text-[11px] font-semibold text-blue-400 backdrop-blur-xs select-none"
              >
                <span className="h-1.5 w-1.5 animate-pulse rounded-full bg-blue-400" />
                AI SALES AGENT FOR E-COMMERCE
              </Badge>
            </div>

            <h1 className="gsap-animate gsap-headline font-heading text-foreground max-w-xl text-4xl leading-[1.1] font-bold tracking-tight sm:text-5xl lg:text-6xl">
              Shopping Should Feel <br />
              <span className="bg-gradient-to-r from-blue-400 via-indigo-400 to-indigo-300 bg-clip-text text-transparent">
                Like Talking to an Expert.
              </span>
            </h1>

            <p className="gsap-animate gsap-subtitle text-muted-foreground max-w-lg text-sm leading-relaxed sm:text-base">
              ShopWise helps customers discover products, compare options, check
              live inventory and promotions, and purchase with confidence
              through conversational AI.
            </p>

            {/* Stats Cards Row */}
            <div className="gsap-animate gsap-stats grid w-full grid-cols-2 gap-3 pt-2 sm:grid-cols-4">
              <div className="border-border/10 bg-muted/5 rounded-lg border p-3">
                <div className="text-foreground text-xl font-bold sm:text-2xl">
                  1+
                </div>
                <div className="text-muted-foreground text-[9px] font-semibold tracking-wider uppercase">
                  BRANDS (Phong Vu)
                </div>
              </div>
              <div className="border-border/10 bg-muted/5 rounded-lg border p-3">
                <div className="text-foreground text-xl font-bold sm:text-2xl">
                  Live
                </div>
                <div className="text-muted-foreground text-[9px] font-semibold tracking-wider uppercase">
                  INVENTORY
                </div>
              </div>
            </div>

            <div className="gsap-animate gsap-actions flex flex-wrap gap-4 pt-2">
              <Button
                asChild
                size="lg"
                className="flex items-center gap-2 shadow-[var(--landing-glow)] shadow-lg transition-all hover:shadow-[var(--landing-glow)] hover:shadow-xl"
              >
                <Link to="/auth">
                  Try AI Sales Agent
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="16"
                    height="16"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="2"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    className="h-4 w-4"
                  >
                    <rect width="18" height="18" x="3" y="3" rx="2" />
                    <path d="M3 9h18" />
                    <path d="M9 21V9" />
                  </svg>
                </Link>
              </Button>
              <Button
                asChild
                variant="outline"
                size="lg"
                className="hover:bg-muted/40 flex items-center gap-2 border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] transition-colors"
              >
                <a href="#problem">
                  See AI in Action
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="16"
                    height="16"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="2"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    className="h-4 w-4"
                  >
                    <path d="m12 3-1.912 5.813a2 2 0 0 1-1.275 1.275L3 12l5.813 1.912a2 2 0 0 1 1.275 1.275L12 21l1.912-5.813a2 2 0 0 1 1.275-1.275L21 12l-5.813-1.912a2 2 0 0 1-1.275-1.275Z" />
                  </svg>
                </a>
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
