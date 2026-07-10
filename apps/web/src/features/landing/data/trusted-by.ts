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

export const TRUSTED_BY_LOGOS: TrustedByLogo[] = [
  {
    id: "acme-corp",
    name: "Acme Corp",
    src: "/src/assets/logos/acme-corp.svg",
    width: 120,
    height: 32,
  },
  {
    id: "nexus-group",
    name: "Nexus Group",
    src: "/src/assets/logos/nexus-group.svg",
    width: 110,
    height: 32,
  },
  {
    id: "vertex-industries",
    name: "Vertex Industries",
    src: "/src/assets/logos/vertex-industries.svg",
    width: 130,
    height: 32,
  },
  {
    id: "orion-retail",
    name: "Orion Retail",
    src: "/src/assets/logos/orion-retail.svg",
    width: 115,
    height: 32,
  },
  {
    id: "pinnacle-logistics",
    name: "Pinnacle Logistics",
    src: "/src/assets/logos/pinnacle-logistics.svg",
    width: 140,
    height: 32,
  },
  {
    id: "zenith-procurement",
    name: "Zenith Procurement",
    src: "/src/assets/logos/zenith-procurement.svg",
    width: 125,
    height: 32,
  },
];
