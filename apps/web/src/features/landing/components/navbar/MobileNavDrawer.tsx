import { Sheet, SheetContent } from "@shopwise/ui/components/sheet";
import { Button } from "@shopwise/ui/components/button";
import { Link } from "@tanstack/react-router";

type MobileNavDrawerProps = {
  isOpen: boolean;
  onClose: () => void;
};

export function MobileNavDrawer({ isOpen, onClose }: MobileNavDrawerProps) {
  const navLinks = [
    { label: "Features", href: "#features" },
    { label: "Workflow", href: "#workflow" },
    { label: "Pricing", href: "#pricing" },
    { label: "FAQ", href: "#faq" },
  ];

  return (
    <Sheet open={isOpen} onOpenChange={(open) => { if (!open) onClose(); }}>
      <SheetContent side="right" className="flex flex-col p-6 w-80">
        <div className="flex flex-col gap-6 mt-8 flex-grow">
          <nav id="mobile-nav" className="flex flex-col gap-4 text-base font-medium">
            {navLinks.map((link) => (
              <a
                key={link.href}
                href={link.href}
                onClick={onClose}
                className="text-muted-foreground hover:text-foreground transition-colors py-2 border-b border-border/20"
              >
                {link.label}
              </a>
            ))}
          </nav>
        </div>
        <div className="flex flex-col gap-3 mt-auto">
          <Button asChild variant="ghost" className="w-full">
            <Link to="/auth" onClick={onClose}>Sign In</Link>
          </Button>
          <Button asChild className="w-full">
            <Link to="/auth" onClick={onClose}>Get Started</Link>
          </Button>
        </div>
      </SheetContent>
    </Sheet>
  );
}
