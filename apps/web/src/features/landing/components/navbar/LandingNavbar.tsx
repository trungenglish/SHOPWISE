import { useState, useEffect } from "react";
import { Link } from "@tanstack/react-router";
import { Button } from "@shopwise/ui/components/button";
import { Menu } from "lucide-react";
import { useScrollSpy } from "@/features/landing/hooks/useScrollSpy";
import { MobileNavDrawer } from "./MobileNavDrawer";

export function LandingNavbar() {
  const [isScrolled, setIsScrolled] = useState(false);
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

  const activeSection = useScrollSpy({
    sectionIds: ["problem", "workspace", "intelligence", "recommendation"],
    offset: 100,
  });

  useEffect(() => {
    const handleScroll = () => {
      setIsScrolled(window.scrollY > 100);
    };

    window.addEventListener("scroll", handleScroll, { passive: true });
    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, []);

  const navLinks = [
    { label: "Problem", id: "problem", href: "#problem" },
    { label: "Workspace", id: "workspace", href: "#workspace" },
    { label: "Intelligence", id: "intelligence", href: "#intelligence" },
    { label: "Recommendation", id: "recommendation", href: "#recommendation" },
  ];

  return (
    <header
      className={`fixed top-0 right-0 left-0 z-40 transition-all duration-300 ${
        isScrolled
          ? "bg-background/80 border-border/40 border-b py-3 backdrop-blur-md"
          : "bg-transparent py-5"
      }`}
    >
      <div className="container mx-auto flex items-center justify-between px-4 md:px-6">
        {/* Left: ShopWise Wordmark */}
        <Link to="/" className="flex items-center space-x-2 select-none">
          <div className="text-background flex h-6 w-6 items-center justify-center rounded-md bg-blue-500">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="14"
              height="14"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              strokeWidth="3"
              strokeLinecap="round"
              strokeLinejoin="round"
            >
              <circle cx="12" cy="12" r="3" fill="currentColor" />
              <path d="M12 2v4M12 18v4M2 12h4M18 12h4" />
            </svg>
          </div>
          <span className="font-heading text-foreground text-sm font-black tracking-widest uppercase">
            SHOPWISE
          </span>
        </Link>

        {/* Center: Desktop Navigation Links */}
        <nav
          aria-label="Main navigation"
          className="hidden items-center space-x-8 md:flex"
        >
          {navLinks.map((link) => {
            const isActive = activeSection === link.id;
            return (
              <a
                key={link.id}
                href={link.href}
                className={`hover:text-foreground relative py-1 text-sm font-medium transition-colors ${
                  isActive ? "text-foreground" : "text-muted-foreground"
                }`}
              >
                {link.label}
                {isActive && (
                  <span className="absolute right-0 bottom-0 left-0 h-0.5 rounded-full bg-[var(--landing-gradient-accent)]" />
                )}
              </a>
            );
          })}
        </nav>

        {/* Right: Actions */}
        <div className="hidden items-center space-x-4 md:flex">
          <Button asChild variant="ghost" size="sm">
            <Link to="/auth">Sign In</Link>
          </Button>
          <Button asChild size="sm">
            <Link to="/auth">Get Started</Link>
          </Button>
        </div>

        {/* Mobile: Hamburger Button */}
        <Button
          variant="ghost"
          size="icon-sm"
          className="md:hidden"
          onClick={() => setIsMobileMenuOpen(true)}
          aria-expanded={isMobileMenuOpen}
          aria-controls="mobile-nav"
        >
          <Menu className="h-5 w-5" />
          <span className="sr-only">Open main menu</span>
        </Button>
      </div>

      {/* Mobile Drawer */}
      <MobileNavDrawer
        isOpen={isMobileMenuOpen}
        onClose={() => setIsMobileMenuOpen(false)}
      />
    </header>
  );
}
