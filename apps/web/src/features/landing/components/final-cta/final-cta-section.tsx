import { useRef } from "react";
import { Link } from "@tanstack/react-router";
import { Button } from "@shopwise/ui/components/button";
import { useGSAP } from "@gsap/react";
import { gsap } from "@/lib/gsap-config";
import { useReducedMotion } from "@/features/landing/hooks/use-reduced-motion";

export function FinalCtaSection() {
  const sectionRef = useRef<HTMLDivElement>(null);
  const prefersReducedMotion = useReducedMotion();

  useGSAP(
    () => {
      if (prefersReducedMotion) {
        return;
      }

      gsap.from(".gsap-cta-reveal", {
        scrollTrigger: {
          trigger: sectionRef.current,
          start: "top 80%",
          toggleActions: "play none none none",
        },
        opacity: 0,
        scale: 0.96,
        duration: 0.7,
        stagger: 0.1,
        ease: "power2.out",
      });
    },
    { scope: sectionRef }
  );

  return (
    <section
      id="contact"
      ref={sectionRef}
      style={{
        background:
          "radial-gradient(circle at center, var(--landing-gradient-mid) 0%, var(--landing-gradient-start) 100%)",
      }}
      className="border-border/20 relative overflow-hidden border-b py-24 lg:py-36"
    >
      <style>{`
        @keyframes shadow-pulse {
          0% { box-shadow: 0 0 0 0 oklch(55% 0.2 290deg / 40%); }
          70% { box-shadow: 0 0 0 10px oklch(55% 0.2 290deg / 0%); }
          100% { box-shadow: 0 0 0 0 oklch(55% 0.2 290deg / 0%); }
        }
        .cta-pulse-button {
          animation: shadow-pulse 2s infinite;
        }
        @media (prefers-reduced-motion: reduce) {
          .cta-pulse-button {
            animation: none !important;
          }
        }
      `}</style>

      {/* Subtle glowing orb in background */}
      <div className="pointer-events-none absolute top-1/2 left-1/2 -z-10 h-[500px] w-[500px] -translate-x-1/2 -translate-y-1/2 rounded-full bg-[var(--landing-gradient-accent)]/5 blur-3xl" />

      <div className="relative z-10 container mx-auto flex flex-col items-center px-4 text-center md:px-6">
        {/* Centered Copy */}
        <span className="gsap-cta-reveal font-heading mb-4 text-xs font-semibold tracking-wider text-[var(--landing-gradient-accent)] uppercase">
          AI Sales Agent
        </span>

        <h2 className="gsap-cta-reveal font-heading text-foreground mb-6 max-w-2xl text-3xl leading-[1.15] font-bold tracking-tight sm:text-4xl lg:text-5xl">
          Ready to Transform Online Shopping?
        </h2>

        <p className="gsap-cta-reveal text-muted-foreground mb-10 max-w-lg text-sm leading-relaxed sm:text-base">
          Bring conversational AI to your e-commerce platform and help customers
          discover the right products faster.
        </p>

        {/* Call to Actions */}
        <div className="gsap-cta-reveal flex flex-wrap justify-center gap-4">
          <Button
            asChild
            size="lg"
            className="cta-pulse-button transition-transform duration-200 hover:scale-[1.02] active:scale-[0.98]"
          >
            <Link to="/auth">Try AI Sales Agent</Link>
          </Button>
          <Button
            asChild
            variant="outline"
            size="lg"
            className="hover:bg-muted/40 border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] transition-colors"
          >
            <a href="mailto:sales@shopwise.io" rel="noopener noreferrer">
              Request Demo
            </a>
          </Button>
        </div>
      </div>
    </section>
  );
}
