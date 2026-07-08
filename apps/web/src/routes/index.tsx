import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
  component: LandingPageComponent,
});

function LandingPageComponent() {
  return (
    <div>
      <h1>Landing Page</h1>
    </div>
  );
}
