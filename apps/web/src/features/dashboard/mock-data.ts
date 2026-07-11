import { Laptop, AuditLog, AgentStatus } from "./types";

export const initialProducts: Laptop[] = [
  {
    id: "macbook-pro",
    name: 'MacBook Pro 16" (M3 Max)',
    price: 3499,
    image:
      "https://images.unsplash.com/photo-1517336714731-489689fd1ca8?w=600&q=80",
    matchScore: 98,
    specs: {
      gpu: "40-core GPU",
      ram: "64GB Unified",
      cooling: "Advanced Dual Fan System",
      cpu: "M3 Max 16-Core",
      screen: '16.2" Liquid Retina XDR',
    },
    aiPerf: 95,
    rendering: 98,
    thermals: 90,
    matchExplanation:
      "The absolute optimal choice for local AI tasks and 3D rendering performance thanks to ultra-high 400GB/s unified memory bandwidth.",
  },
  {
    id: "razer-blade",
    name: "Razer Blade 16",
    price: 3599.99,
    image:
      "https://images.unsplash.com/photo-1603302576837-37561b2e2302?w=600&q=80",
    matchScore: 94,
    specs: {
      gpu: "NVIDIA RTX 4090",
      ram: "32GB DDR5 Dual Channel",
      cooling: "Vapor Chamber Active Cooling",
      cpu: "Intel Core i9-14900HX",
      screen: '16" Dual UHD+/FHD+ Mini-LED',
    },
    aiPerf: 99,
    rendering: 95,
    thermals: 85,
    matchExplanation:
      "Equipped with RTX 4090 hardware with max TGP for outstanding real-time Ray Tracing and hardware rendering performance.",
  },
  {
    id: "dell-xps",
    name: "Dell XPS 16",
    price: 3249,
    image:
      "https://images.unsplash.com/photo-1593642632823-8f785ba67e45?w=600&q=80",
    matchScore: 88,
    specs: {
      gpu: "NVIDIA RTX 4070",
      ram: "32GB LPDDR5x Dual Channel",
      cooling: "Dual Fans with GORE Thermal",
      cpu: "Intel Core Ultra 9 185H",
      screen: '16.3" 4K+ OLED Touch',
    },
    aiPerf: 80,
    rendering: 82,
    thermals: 75,
    matchExplanation:
      "The perfect combination of elegant minimalist design, vibrant OLED display, and high durability for mobile content creators.",
  },
];

export const initialLogs: AuditLog[] = [
  {
    time: "09:14",
    message: "Analyze user hardware requirements and constraints",
    status: "done",
  },
  {
    time: "09:14",
    message: "Map optimal specifications (GPU & Cool)",
    status: "done",
  },
  {
    time: "09:15",
    message: "Cross-check inventory data from 12 official retailers",
    status: "done",
  },
  {
    time: "09:16",
    message: "Conducting real-world thermal performance evaluation...",
    status: "running",
  },
];

export const initialAgents: AgentStatus[] = [
  {
    id: "shopping-expert",
    name: "Shopping Expert",
    progress: 100,
    statusMessage:
      "> Optimal search -> Sending detailed query to Inventory Agent",
    type: "shopping",
  },
  {
    id: "inventory-agent",
    name: "Inventory Agent",
    progress: 85,
    statusMessage: "> Verified matching with 12/12 official online stores",
    type: "inventory",
  },
  {
    id: "warranty-agent",
    name: "Warranty Agent",
    progress: 40,
    statusMessage:
      "> Reviewing gold warranty policies & accidental damage support...",
    type: "warranty",
  },
  {
    id: "finance-agent",
    name: "Finance Agent",
    progress: 0,
    statusMessage:
      "> Awaiting inventory confirmation from agents to optimize pricing...",
    type: "finance",
  },
];

export const initialAccessories = [
  {
    name: 'Apple Studio Display 27"',
    price: 1599,
    category: "DISPLAY",
    reason: "Absolute color accuracy for 3D creative design.",
  },
  {
    name: "Razer Thunderbolt 4 Dock Chroma",
    price: 329,
    category: "ADAPTERS",
    reason:
      "Maximum expansion of high-speed ports for heavy external storage devices.",
  },
  {
    name: "IETS GT500 Powerful Turbo Cooling Pad",
    price: 59,
    category: "COOLING",
    reason:
      "High-speed centrifugal cooling fan helps maintain stable temperatures during long rendering sessions.",
  },
];

export const initialReasoning = `
# HARDWARE DECISION MODEL ANALYSIS

### Overview
We conducted a multi-dimensional analysis for the **Intensive 3D rendering and local AI development** configuration. The key requirement lies in the graphics card's memory bandwidth to load large AI models (LLMs/Diffusion models) and an optimal cooling solution against thermal throttling.

### Evaluation of Core Candidates

* **MacBook Pro 16" (M3 Max)**: 
  Outstanding product thanks to **Unified Memory** technology up to 64GB. This allows loading large language models (LLM) exceeding the standard 16GB VRAM on mobile PCs. The dual thermal system operates extremely quietly even under continuous high rendering loads.
  
* **Razer Blade 16**:
  Leading in pure graphics processing thanks to the **NVIDIA RTX 4090** card and **Vapor Chamber** cooling. This device has no rival in real-time ray tracing and training small AI models using next-generation CUDA Tensor Cores. However, fan noise under heavy load is quite noticeable.
  
* **Dell XPS 16**:
  A balanced choice of trendiness and mobile performance. The **4K+ OLED** display provides an excellent creative space with extreme color accuracy. The RTX 4070 runs cool but is bandwidth-limited when processing massive AI datasets.

### Conclusion & Trade-offs
If prioritizing testing large AI models, the **MacBook Pro** is the absolute optimal choice. If your graphics project requires maximum DirectX 12 or Unreal Engine compatibility, choose the **Razer Blade 16**.
`;
