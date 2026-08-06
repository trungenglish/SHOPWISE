import { useRef } from "react";
import {
  AlertCircle,
  CheckCircle2,
  Search,
  MessageSquare,
  ShieldAlert,
} from "lucide-react";
import { useGSAP } from "@gsap/react";
import { gsap } from "@/lib/gsap-config";
import { useReducedMotion } from "@/features/landing/hooks/use-reduced-motion";

export function ProblemStatementSection() {
  const containerRef = useRef<HTMLDivElement>(null);
  const prefersReducedMotion = useReducedMotion();

  useGSAP(
    () => {
      if (prefersReducedMotion) {
        return;
      }

      gsap.from(".gsap-problem-reveal", {
        scrollTrigger: {
          trigger: containerRef.current,
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
    { scope: containerRef }
  );

  const painPoints = [
    "Too many products makes choosing overwhelming for customers.",
    "Hard to compare specifications across different brands and models.",
    "Live stock uncertainty causes cart abandonment and customer drop-off.",
    "Promotions and campaign discounts are difficult to find and apply.",
    "Human support is unavailable during peak hours or late nights.",
    "Customers leave the site without purchasing due to unanswered questions.",
  ];

  const outcomes = [
    "AI answers instantly, helping shoppers make immediate choices.",
    "Smart recommendations tailored to budget and technical requirements.",
    "Live inventory queries fetch real-time warehouse data instantly.",
    "Product comparison highlights key differences side-by-side.",
    "Promotion lookup automatically finds the best coupons and vouchers.",
    "Guided checkout directs the customer directly to purchase completion.",
  ];

  return (
    <section
      id="problem"
      ref={containerRef}
      className="bg-background border-border/20 relative border-b py-20 lg:py-32"
    >
      <div className="container mx-auto px-4 md:px-6">
        {/* Section Header */}
        <div className="gsap-problem-reveal mx-auto mb-16 max-w-3xl text-center">
          <span className="font-heading text-xs font-semibold tracking-wider text-[var(--landing-gradient-accent)] uppercase">
            The Product Discovery Problem
          </span>
          <h2 className="font-heading text-foreground mt-3 text-3xl font-bold tracking-tight sm:text-4xl lg:text-5xl">
            The Product Discovery Problem
          </h2>
          <p className="text-muted-foreground mt-4 text-base sm:text-lg">
            Online shoppers face overwhelming options and stale inventory. AI
            Sales Agent bridges the gap between searching and purchasing.
          </p>
        </div>

        {/* Two Column Comparison */}
        <div className="grid grid-cols-1 gap-12 lg:grid-cols-2 lg:gap-16">
          {/* Legacy Side */}
          <div
            role="group"
            aria-labelledby="legacy-title"
            className="gsap-problem-reveal border-border/10 bg-muted/5 flex flex-col rounded-xl border p-6 lg:p-8"
          >
            <h3
              id="legacy-title"
              className="font-heading text-muted-foreground/80 flex items-center gap-2 text-lg font-bold sm:text-xl"
            >
              <AlertCircle className="text-muted-foreground/60 h-5 w-5" />
              How customers buy today
            </h3>
            <p className="text-muted-foreground/60 mt-2 text-sm">
              The legacy search and browse process leads to shopper fatigue and
              dropped carts.
            </p>
            <ul className="mt-6 flex-grow space-y-4">
              {painPoints.map((pain, idx) => (
                <li
                  key={idx}
                  className="text-muted-foreground/70 flex items-start gap-3 text-sm"
                >
                  <span className="bg-muted-foreground/30 mt-1 flex h-1.5 w-1.5 shrink-0 rounded-full" />
                  <span>{pain}</span>
                </li>
              ))}
            </ul>
          </div>

          {/* ShopWise Side */}
          <div
            role="group"
            aria-labelledby="solution-title"
            className="gsap-problem-reveal relative flex flex-col rounded-xl border border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] p-6 shadow-2xl backdrop-blur-xs lg:p-8"
          >
            {/* Edge Glow */}
            <div className="pointer-events-none absolute -inset-px -z-10 rounded-xl bg-gradient-to-r from-[var(--landing-gradient-accent)]/20 to-indigo-500/20 opacity-50 blur-xs" />
            <h3
              id="solution-title"
              className="font-heading text-foreground flex items-center gap-2 text-lg font-bold sm:text-xl"
            >
              <CheckCircle2 className="h-5 w-5 text-[var(--landing-gradient-accent)]" />
              How ShopWise works
            </h3>
            <p className="text-muted-foreground mt-2 text-sm">
              A conversational sales assistant that converts queries into
              e-commerce checkouts.
            </p>
            <ul className="mt-6 flex-grow space-y-4">
              {outcomes.map((outcome, idx) => (
                <li
                  key={idx}
                  className="text-foreground/90 flex items-start gap-3 text-sm"
                >
                  <span className="mt-1 flex h-1.5 w-1.5 shrink-0 rounded-full bg-[var(--landing-gradient-accent)]" />
                  <span>{outcome}</span>
                </li>
              ))}
            </ul>
          </div>
        </div>

        {/* Differentiation Block */}
        <div className="gsap-problem-reveal border-border/10 bg-muted/5 mt-16 rounded-xl border p-6 lg:p-8">
          <div className="grid grid-cols-1 items-start gap-6 lg:grid-cols-12 lg:gap-8">
            <div className="flex items-center gap-3 lg:col-span-4">
              <ShieldAlert className="text-muted-foreground/60 h-6 w-6 shrink-0" />
              <h4 className="font-heading text-foreground text-base font-bold sm:text-lg">
                Why ShopWise is Different
              </h4>
            </div>
            <div className="lg:col-span-8">
              <p className="text-muted-foreground text-sm leading-relaxed">
                ShopWise is{" "}
                <strong className="text-foreground">
                  not another search engine
                </strong>
                ,{" "}
                <strong className="text-foreground">
                  not a generic chatbot
                </strong>
                , and{" "}
                <strong className="text-foreground">
                  not a procurement marketplace
                </strong>
                . It is an AI Sales Agent — a persistent, conversational
                commerce environment that maintains shopping context, product
                specifications, inventory state, and cart history across the
                entire customer journey.
              </p>
              <div className="text-muted-foreground/60 mt-4 flex flex-wrap gap-4 text-xs">
                <span className="flex items-center gap-1">
                  <Search className="h-3 w-3" /> No standard search queries
                </span>
                <span className="flex items-center gap-1">
                  <MessageSquare className="h-3 w-3" /> No generic conversations
                </span>
                <span className="flex items-center gap-1">
                  <CheckCircle2 className="h-3 w-3" /> Live catalog sync
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
