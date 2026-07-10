import { useEffect, useState } from "react";

type UseScrollSpyOptions = {
  sectionIds: string[];
  /** Distance from the top of the viewport that counts as "active" (px). Default: 100 */
  offset?: number;
};

/**
 * Tracks which page section is currently visible in the viewport using
 * IntersectionObserver. Returns the `id` of the active section or `null`
 * when none match.
 *
 * Used by LandingNavbar to highlight the nav link for the visible section.
 */
export function useScrollSpy({
  sectionIds,
  offset = 100,
}: UseScrollSpyOptions): string | null {
  const [activeId, setActiveId] = useState<string | null>(null);

  useEffect(() => {
    if (sectionIds.length === 0) return;

    const rootMarginTop = `-${offset}px`;

    const observer = new IntersectionObserver(
      (entries) => {
        for (const entry of entries) {
          if (entry.isIntersecting) {
            setActiveId(entry.target.id);
          }
        }
      },
      {
        rootMargin: `${rootMarginTop} 0px -50% 0px`,
        threshold: 0,
      },
    );

    for (const id of sectionIds) {
      const element = document.getElementById(id);
      if (element) {
        observer.observe(element);
      }
    }

    return () => {
      observer.disconnect();
    };
  }, [sectionIds, offset]);

  return activeId;
}
