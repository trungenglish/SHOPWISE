export type TrustedByLogo = {
  id: string;
  /** Company name — used as the img alt attribute (accessibility requirement). */
  name: string;
  /** Path to the monochrome SVG asset under src/assets/logos/. */
  src: string;
  /** Intrinsic width in px — prevents CLS while image loads. */
  width: number;
  /** Intrinsic height in px. */
  height: number;
};

export const TRUSTED_BY_LOGOS: TrustedByLogo[] = [];
