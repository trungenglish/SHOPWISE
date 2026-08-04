import { createFileRoute } from "@tanstack/react-router";
import ChatWorkspace from "@/components/chat-workspace";

export const Route = createFileRoute("/chat")({
  component: ChatComponent,
});

function ChatComponent() {
  return (
    <div className="p-8">
      <h1 className="mb-4 text-2xl font-bold">AI Shopping Conversation</h1>
      <ChatWorkspace />
    </div>
  );
}
