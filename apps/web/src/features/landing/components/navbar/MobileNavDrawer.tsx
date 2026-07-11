import { Sheet, SheetContent } from "@shopwise/ui/components/sheet";
import { Button } from "@shopwise/ui/components/button";
import { Link } from "@tanstack/react-router";

type MobileNavDrawerProps = {
  isOpen: boolean;
  onClose: () => void;
};

export function MobileNavDrawer({ isOpen, onClose }: MobileNavDrawerProps) {
  const navLinks = [
    { label: "Problem", href: "#problem" },
    { label: "Workspace", href: "#workspace" },
    { label: "Intelligence", href: "#intelligence" },
    { label: "Recommendation", href: "#recommendation" },
  ];

  return (
    <Sheet
      open={isOpen}
      onOpenChange={(open) => {
        if (!open) onClose();
      }}
    >
      <SheetContent side="right" className="flex w-80 flex-col p-6">
        <div className="mt-8 flex flex-grow flex-col gap-6">
          <nav
            id="mobile-nav"
            className="flex flex-col gap-4 text-base font-medium"
          >
            {navLinks.map((link) => (
              <a
                key={link.href}
                href={link.href}
                onClick={onClose}
                className="text-muted-foreground hover:text-foreground border-border/20 border-b py-2 transition-colors"
              >
                {link.label}
              </a>
            ))}
          </nav>
        </div>
        <div className="mt-auto flex flex-col gap-3">
          <Button asChild variant="ghost" className="w-full">
            <Link to="/auth" onClick={onClose}>
              Sign In
            </Link>
          </Button>
          <Button asChild className="w-full">
            <Link to="/auth" onClick={onClose}>
              Get Started
            </Link>
          </Button>
        </div>
      </SheetContent>
    </Sheet>
  );
}
