import { useRef } from "react";
import { useGSAP } from "@gsap/react";
import { gsap } from "@/lib/gsap-config";
import { useReducedMotion } from "@/features/landing/hooks/use-reduced-motion";
import { WorkflowStep } from "../../data/workflow-steps";

type WorkflowTimelineProps = {
  steps: WorkflowStep[];
};

export function WorkflowTimeline({ steps }: WorkflowTimelineProps) {
  const containerRef = useRef<HTMLDivElement>(null);
  const prefersReducedMotion = useReducedMotion();

  useGSAP(
    () => {
      if (prefersReducedMotion) {
        return;
      }

      // GSAP ScrollTrigger timeline to draw line and stagger steps
      const tl = gsap.timeline({
        scrollTrigger: {
          trigger: containerRef.current,
          start: "top 70%",
          toggleActions: "play none none none",
        },
      });

      tl.from(".gsap-connector-line", {
        scaleX: 0,
        transformOrigin: "left center",
        duration: 0.8,
        ease: "power2.inOut",
      })
        .from(
          ".gsap-connector-line-vertical",
          {
            scaleY: 0,
            transformOrigin: "center top",
            duration: 0.8,
            ease: "power2.inOut",
          },
          "<"
        ) // runs at same time
        .from(
          ".gsap-step-badge",
          {
            scale: 0,
            opacity: 0,
            stagger: 0.15,
            duration: 0.4,
            ease: "back.out(1.7)",
          },
          "-=0.4"
        )
        .from(
          ".gsap-step-content",
          {
            opacity: 0,
            y: 20,
            stagger: 0.15,
            duration: 0.5,
            ease: "power2.out",
          },
          "-=0.4"
        );
    },
    { scope: containerRef }
  );

  return (
    <div ref={containerRef} className="relative w-full">
      {/* Connector line (Horizontal for desktop, >= 1024px) */}
      <div
        aria-hidden="true"
        className="gsap-connector-line absolute top-6 right-8 left-8 -z-10 hidden h-0.5 bg-gradient-to-r from-[var(--landing-gradient-accent)]/20 via-[var(--landing-gradient-accent)]/40 to-[var(--landing-gradient-accent)]/10 lg:block"
      />

      {/* Connector line (Vertical for mobile/tablet, < 1024px) */}
      <div
        aria-hidden="true"
        className="gsap-connector-line-vertical absolute top-8 bottom-8 left-6 -z-10 w-0.5 bg-gradient-to-b from-[var(--landing-gradient-accent)]/20 via-[var(--landing-gradient-accent)]/40 to-[var(--landing-gradient-accent)]/10 lg:hidden"
      />

      <ol className="m-0 grid list-none grid-cols-1 gap-10 p-0 lg:grid-cols-7 lg:gap-4">
        {steps.map((step) => (
          <li
            key={step.step}
            className="flex flex-row items-start space-x-5 lg:flex-col lg:space-y-4 lg:space-x-0"
          >
            {/* Number Circle Badge */}
            <span className="gsap-step-badge text-foreground font-heading relative z-10 flex h-12 w-12 shrink-0 items-center justify-center rounded-full border border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] text-base font-bold shadow-md transition-colors duration-300 hover:border-[var(--landing-gradient-accent)]">
              {step.step}
            </span>

            {/* Text Content Block */}
            <div className="gsap-step-content flex flex-col space-y-1 pt-1.5 lg:pt-0">
              <h3 className="text-foreground font-heading text-sm leading-tight font-semibold">
                {step.title}
              </h3>
              <p className="text-muted-foreground text-[11px] leading-relaxed">
                {step.description}
              </p>
            </div>
          </li>
        ))}
      </ol>
    </div>
  );
}
