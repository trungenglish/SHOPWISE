import { ReactLenis, useLenis } from "lenis/react";
import { ScrollTrigger } from "@/lib/gsap-config";
import { LandingNavbar } from "./components/navbar/landing-navbar";
import { HeroSection } from "./components/hero/hero-section";
import { TrustedBySection } from "./components/trusted-by/trusted-by-section";
import { ProblemStatementSection } from "./components/problem/problem-statement-section";
import { AiDecisionSection } from "./components/ai-viz/ai-decision-section";
import { FeaturesSection } from "./components/features/features-section";
import { WorkflowSection } from "./components/workflow/workflow-section";
import { DecisionConfidenceSection } from "./components/confidence/decision-confidence-section";
import { FinalCtaSection } from "./components/final-cta/final-cta-section";
import { LandingFooter } from "./components/footer/landing-footer";

function ScrollTriggerSync() {
  useLenis(() => {
    ScrollTrigger.update();
  });
  return null;
}

export function LandingPage() {
  return (
    <ReactLenis root>
      <ScrollTriggerSync />
      <div className="flex min-h-screen flex-col">
        {/* Navigation Bar */}
        <LandingNavbar />

        <main className="flex-grow">
          {/* Hero Section */}
          <HeroSection />

          {/* Trusted By Brand Logos */}
          <TrustedBySection />

          {/* Problem Statement Section */}
          <ProblemStatementSection />

          {/* AI Decision Tool Visualisation */}
          <AiDecisionSection />

          {/* Product Features Grid */}
          <FeaturesSection />

          {/* How It Works Workflow Steps */}
          <WorkflowSection />

          {/* Decision Confidence Section */}
          <DecisionConfidenceSection />

          {/* Final Call To Action */}
          <FinalCtaSection />
        </main>

        {/* Brand Footer */}
        <LandingFooter />
      </div>
    </ReactLenis>
  );
}
