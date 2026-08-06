import { TrustedByLogo } from "../../data/trusted-by";

type LogoMarqueeProps = {
  logos: TrustedByLogo[];
  pauseOnHover?: boolean;
};

export function LogoMarquee({ logos, pauseOnHover = true }: LogoMarqueeProps) {
  // Duplicate logos for seamless looping marquee
  const duplicatedLogos = [...logos, ...logos];

  return (
    <div
      role="region"
      aria-label="Trusted by — company logos"
      className="relative w-full overflow-hidden py-4 select-none"
    >
      <style>{`
        @keyframes marquee {
          0% {
            transform: translateX(0);
          }
          100% {
            transform: translateX(-50%);
          }
        }
        .marquee-track {
          display: flex;
          width: max-content;
          animation: marquee 30s linear infinite;
        }
        ${
          pauseOnHover
            ? `.marquee-track:hover {
                animation-play-state: paused;
              }`
            : ""
        }
        @media (max-width: 640px) {
          .marquee-track {
            animation: none !important;
            transform: none !important;
            justify-content: center;
            width: 100%;
            flex-wrap: wrap;
            gap: 1.5rem 2.5rem;
          }
        }
        @media (prefers-reduced-motion: reduce) {
          .marquee-track {
            animation: none !important;
            transform: none !important;
            justify-content: center;
            width: 100%;
            flex-wrap: wrap;
            gap: 2rem;
          }
        }
      `}</style>

      <div className="marquee-track flex items-center gap-16">
        {duplicatedLogos.map((logo, index) => (
          <div
            key={`${logo.id}-${index}`}
            className={`text-muted-foreground/30 hover:text-muted-foreground/60 flex items-center justify-center transition-colors duration-200 ${
              index >= logos.length ? "hidden sm:flex" : ""
            }`}
          >
            <img
              src={logo.src}
              alt={logo.name}
              width={logo.width}
              height={logo.height}
              loading="lazy"
              className="h-7 w-auto object-contain"
            />
          </div>
        ))}
      </div>
    </div>
  );
}
