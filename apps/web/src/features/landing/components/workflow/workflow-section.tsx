import { useRef } from "react";
import { useGSAP } from "@gsap/react";
import { gsap } from "@/lib/gsap-config";
import { useReducedMotion } from "@/features/landing/hooks/use-reduced-motion";
import { WORKFLOW_STEPS } from "../../data/workflow-steps";
import { WorkflowTimeline } from "./workflow-timeline";

export function WorkflowSection() {
  const sectionRef = useRef<HTMLDivElement>(null);
  const prefersReducedMotion = useReducedMotion();

  useGSAP(
    () => {
      if (prefersReducedMotion) {
        return;
      }

      gsap.from(".gsap-workflow-reveal", {
        scrollTrigger: {
          trigger: sectionRef.current,
          start: "top 75%",
          toggleActions: "play none none none",
        },
        opacity: 0,
        y: 30,
        stagger: 0.1,
        duration: 0.6,
        ease: "power2.out",
      });
    },
    { scope: sectionRef }
  );

  return (
    <section
      id="intelligence"
      ref={sectionRef}
      className="bg-background border-border/20 border-b py-20 lg:py-32"
    >
      <div className="container mx-auto px-4 md:px-6">
        {/* Header Block */}
        <div className="mb-16 flex flex-col items-center space-y-4 text-center lg:mb-24">
          <span className="gsap-workflow-reveal font-heading text-xs font-semibold tracking-wider text-[var(--landing-gradient-accent)] uppercase">
            How ShopWise Works
          </span>
          <h2 className="gsap-workflow-reveal font-heading text-foreground max-w-2xl text-3xl leading-[1.15] font-bold tracking-tight sm:text-4xl lg:text-5xl">
            How ShopWise Works
          </h2>
          <p className="gsap-workflow-reveal text-muted-foreground max-w-lg text-sm sm:text-base">
            See how ShopWise handles user requests, retrieves product
            specifications, and delivers verified recommendations.
          </p>
        </div>

        {/* Workflow Timeline Component */}
        <WorkflowTimeline steps={WORKFLOW_STEPS} />
      </div>
    </section>
  );
}
