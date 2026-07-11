import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/error/something-went-wrong")({
  component: RouteComponent,
});

function RouteComponent() {
  return <div>Hello "/error/something-went-wrong"!</div>;
}
