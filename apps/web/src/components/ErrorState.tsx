export interface ErrorStateProps {
  title: string;
  message: string;
  onDismiss: () => void;
}

export function ErrorState({ title, message, onDismiss }: ErrorStateProps) {
  return (
    <div className="fixed inset-0 bg-background/80 backdrop-blur-sm z-50 flex items-center justify-center">
      <div className="bg-card text-card-foreground border rounded-lg shadow-lg p-6 max-w-md w-full">
        <h3 className="text-lg font-bold text-destructive mb-2">{title}</h3>
        <p className="mb-4">{message}</p>
        <button
          onClick={onDismiss}
          className="bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-4 py-2 rounded-md font-medium w-full"
        >
          Dismiss
        </button>
      </div>
    </div>
  );
}
