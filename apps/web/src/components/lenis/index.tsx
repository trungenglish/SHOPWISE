import type { ReactNode } from "react";
import { ReactLenis } from "lenis/react";

type LenisProps = {
  children: ReactNode;
};

export function Lenis({ children }: LenisProps) {
  return (
    <ReactLenis
      root
      options={{
        duration: 1.2,
        easing: (t) => Math.min(1, 1.001 - 2 ** (-10 * t)),
        orientation: "vertical",
        smoothWheel: true,
      }}
    >
      {children}
    </ReactLenis>
  );
}