import { gsap } from "gsap";
import { ScrollTrigger } from "gsap/ScrollTrigger";

gsap.registerPlugin(ScrollTrigger);

// Ensure Lenis is compatible with ScrollTrigger
ScrollTrigger.normalizeScroll(false);

export { gsap, ScrollTrigger };
