import { createFileRoute } from "@tanstack/react-router";
import { GetStartedFlow } from "@/features/get-started/components/get-started-flow";

export const Route = createFileRoute("/get-started")({
  component: GetStartedFlow,
});
