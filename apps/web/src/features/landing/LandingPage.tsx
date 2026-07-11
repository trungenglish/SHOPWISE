import { ReactLenis, useLenis } from "lenis/react";
import { ScrollTrigger } from "@/lib/gsap-config";
import { LandingNavbar } from "./components/navbar/LandingNavbar";
import { HeroSection } from "./components/hero/HeroSection";
import { TrustedBySection } from "./components/trusted-by/TrustedBySection";
import { ProblemStatementSection } from "./components/problem/ProblemStatementSection";
import { AiDecisionSection } from "./components/ai-viz/AiDecisionSection";
import { FeaturesSection } from "./components/features/FeaturesSection";
import { WorkflowSection } from "./components/workflow/WorkflowSection";
import { DecisionConfidenceSection } from "./components/confidence/DecisionConfidenceSection";
import { FinalCtaSection } from "./components/final-cta/FinalCtaSection";
import { LandingFooter } from "./components/footer/LandingFooter";

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
