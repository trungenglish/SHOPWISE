import { createFileRoute } from "@tanstack/react-router";
import { LandingPage } from "@/features/landing/LandingPage";

export const Route = createFileRoute("/")({
  component: LandingPageComponent,
  head: () => ({
    meta: [
      {
        title: "ShopWise — AI-Powered Procurement Intelligence",
      },
      {
        name: "description",
        content:
          "ShopWise is an AI-powered shopping decision platform for enterprise procurement teams.",
      },
    ],
  }),
});

function LandingPageComponent() {
  return <LandingPage />;
}
