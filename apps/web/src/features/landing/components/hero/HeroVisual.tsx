export function HeroVisual() {
  return (
    <div
      aria-hidden="true"
      className="relative flex h-full min-h-[300px] w-full items-center justify-center lg:min-h-[500px]"
    >
      <style>{`
        @keyframes float {
          0%, 100% {
            transform: translateY(0px) rotate(0deg);
          }
          50% {
            transform: translateY(-15px) rotate(1deg);
          }
        }
        @keyframes float-delayed {
          0%, 100% {
            transform: translateY(0px) rotate(0deg);
          }
          50% {
            transform: translateY(12px) rotate(-1.5deg);
          }
        }
        @keyframes pulse-slow {
          0%, 100% {
            opacity: 0.15;
          }
          50% {
            opacity: 0.45;
          }
        }
        .hero-visual-card {
          animation: float 6s ease-in-out infinite;
        }
        .hero-visual-card-delayed {
          animation: float-delayed 8s ease-in-out infinite;
        }
        .pulse-element {
          animation: pulse-slow 4s ease-in-out infinite;
        }
        @media (prefers-reduced-motion: reduce) {
          .hero-visual-card,
          .hero-visual-card-delayed,
          .pulse-element {
            animation: none !important;
            transform: none !important;
          }
        }
      `}</style>

      {/* Background glow effects */}
      <div className="absolute inset-0 bg-gradient-to-tr from-[var(--landing-gradient-start)] via-transparent to-[var(--landing-gradient-accent)]/10 opacity-60 blur-3xl" />

      {/* Decorative Network Nodes SVG */}
      <svg
        viewBox="0 0 400 400"
        className="absolute h-[80%] w-[80%] text-[var(--landing-gradient-accent)]/20"
        fill="none"
        stroke="currentColor"
        strokeWidth="1"
      >
        <circle
          cx="200"
          cy="200"
          r="120"
          className="pulse-element"
          strokeDasharray="4 4"
        />
        <circle cx="200" cy="200" r="80" opacity="0.5" />
        <path
          d="M80,200 L320,200 M200,80 L200,320"
          strokeDasharray="2 2"
          opacity="0.3"
        />
        <circle cx="200" cy="80" r="4" fill="currentColor" />
        <circle cx="200" cy="320" r="4" fill="currentColor" />
        <circle cx="80" cy="200" r="4" fill="currentColor" />
        <circle cx="320" cy="200" r="4" fill="currentColor" />
      </svg>

      {/* Floating Glassmorphism Decision Cards */}
      <div className="hero-visual-card relative z-10 flex h-36 w-64 flex-col justify-between rounded-xl border border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] p-5 shadow-2xl backdrop-blur-md">
        <div className="flex items-center justify-between">
          <span className="text-muted-foreground font-heading text-[10px] tracking-wider uppercase">
            Decision Node
          </span>
          <span className="h-2 w-2 rounded-full bg-emerald-500" />
        </div>
        <div>
          <div className="text-foreground text-sm font-semibold">
            AI Procurement Model
          </div>
          <div className="text-muted-foreground mt-1 text-[11px]">
            Analyzing Category Q3 Spend...
          </div>
        </div>
        <div className="bg-muted/30 h-1.5 w-full overflow-hidden rounded-full">
          <div className="h-full w-3/4 rounded-full bg-[var(--landing-gradient-accent)]" />
        </div>
      </div>

      <div className="hero-visual-card-delayed absolute -right-4 -bottom-4 z-20 flex h-28 w-48 flex-col justify-between rounded-xl border border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] p-4 shadow-2xl backdrop-blur-md lg:-right-8 lg:bottom-12">
        <div className="flex items-center justify-between">
          <span className="text-muted-foreground font-heading text-[10px] tracking-wider uppercase">
            Vendor score
          </span>
          <span className="text-foreground text-xs font-semibold">98.4%</span>
        </div>
        <div className="text-foreground text-xs font-semibold">
          Apex Services LLC
        </div>
        <div className="flex gap-1">
          <span className="rounded bg-emerald-500/10 px-1.5 py-0.5 text-[8px] font-medium text-emerald-400">
            Optimal
          </span>
          <span className="rounded bg-[var(--landing-gradient-accent)]/10 px-1.5 py-0.5 text-[8px] font-medium text-[var(--landing-gradient-accent)]">
            Low Risk
          </span>
        </div>
      </div>
    </div>
  );
}
