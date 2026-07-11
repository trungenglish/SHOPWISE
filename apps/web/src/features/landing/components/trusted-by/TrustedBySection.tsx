import { TRUSTED_BY_LOGOS } from "../../data/trusted-by";
import { LogoMarquee } from "./LogoMarquee";

export function TrustedBySection() {
  if (!TRUSTED_BY_LOGOS || TRUSTED_BY_LOGOS.length === 0) {
    return null;
  }

  return (
    <section
      id="trusted-by"
      className="bg-background border-border/20 overflow-hidden border-b py-10"
    >
      <div className="container mx-auto flex flex-col items-center px-4 md:px-6">
        <p className="text-muted-foreground/60 font-heading mb-6 text-center text-xs font-semibold tracking-wider uppercase select-none">
          Trusted by leading enterprises worldwide
        </p>

        {/* Logo Marquee & Grid */}
        <div className="w-full">
          <LogoMarquee logos={TRUSTED_BY_LOGOS} />
        </div>
      </div>
    </section>
  );
}
