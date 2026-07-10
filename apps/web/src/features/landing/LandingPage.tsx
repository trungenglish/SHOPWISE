import { ReactLenis } from "lenis/react";
import { LandingNavbar } from "./components/navbar/LandingNavbar";
import { HeroSection } from "./components/hero/HeroSection";
import { TrustedBySection } from "./components/trusted-by/TrustedBySection";
import { AiDecisionSection } from "./components/ai-viz/AiDecisionSection";
import { FeaturesSection } from "./components/features/FeaturesSection";
import { WorkflowSection } from "./components/workflow/WorkflowSection";
import { TestimonialsSection } from "./components/testimonials/TestimonialsSection";
import { PricingSection } from "./components/pricing/PricingSection";
import { FaqSection } from "./components/faq/FaqSection";
import { FinalCtaSection } from "./components/final-cta/FinalCtaSection";
import { LandingFooter } from "./components/footer/LandingFooter";

export function LandingPage() {
  return (
    <ReactLenis root>
      <div className="flex min-h-screen flex-col">
        {/* Navigation Bar */}
        <LandingNavbar />

        <main className="flex-grow">
          {/* Hero Section */}
          <HeroSection />

          {/* Trusted By Brand Logos */}
          <TrustedBySection />

          {/* AI Decision Tool Visualisation */}
          <AiDecisionSection />

          {/* Product Features Grid */}
          <FeaturesSection />

          {/* How It Works Workflow Steps */}
          <WorkflowSection />

          {/* Customer Testimonials Carousel */}
          <TestimonialsSection />

          {/* Transparent Pricing Table */}
          <PricingSection />

          {/* Frequently Asked Questions */}
          <FaqSection />

          {/* Final Call To Action */}
          <FinalCtaSection />
        </main>

        {/* Global Landing Footer */}
        <LandingFooter />
      </div>
    </ReactLenis>
  );
}
