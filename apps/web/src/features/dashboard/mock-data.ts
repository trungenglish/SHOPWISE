import { Laptop, AuditLog, AgentStatus } from "./types";

export const initialProducts: Laptop[] = [
  {
    id: "rog-strix-g16-5057",
    name: "ROG Strix G16 (RTX 5057)",
    price: 38500000,
    image:
      "https://images.unsplash.com/photo-1603302576837-37561b2e2302?w=600&q=80",
    matchScore: 85,
    specs: {
      gpu: "RTX 5057, DLSS 4",
      ram: "16-32GB DDR5",
      cooling: "Solid standard ROG thermal",
      cpu: "Ryzen 9 8940HX / Core Ultra 9 275HX",
      screen: '16" FHD, 165Hz',
      warranty: "1-Year Premium Support",
    },
    aiPerf: 80,
    rendering: 85,
    thermals: 85,
    matchExplanation:
      "Best Value: Perfect for the budget-conscious AAA gamer wanting solid 1440p performance.",
  },
  {
    id: "rog-strix-g16-5080",
    name: "ROG Strix G16 (RTX 5070 Ti/5080)",
    price: 50000000,
    image:
      "https://images.unsplash.com/photo-1593640408182-31c70c8268f5?w=600&q=80",
    matchScore: 98,
    specs: {
      gpu: "RTX 5070 Ti-5080, up to 175W TGP",
      ram: "32GB DDR5",
      cooling: "Reinforced for sustained 175W TGP",
      cpu: "Core Ultra 9 290HX Plus",
      screen: '16" 2.5K, up to 300Hz, 100% DCI-P3',
      warranty: "2-Year Manufacturer",
    },
    aiPerf: 95,
    rendering: 98,
    thermals: 92,
    matchExplanation:
      "Best Balance: Ideal for gamers wanting max FPS and highest settings with sustained performance.",
  },
  {
    id: "rog-zephyrus-g16-5070ti",
    name: "ROG Zephyrus G16 (RTX 5070 Ti, OLED)",
    price: 55000000,
    image:
      "https://images.unsplash.com/photo-1593642632823-8f785ba67e45?w=600&q=80",
    matchScore: 92,
    specs: {
      gpu: "RTX 5070 Ti, DLSS 4",
      ram: "32GB DDR5",
      cooling: "Slim chassis, stays ~70-76°C",
      cpu: "Core Ultra 9 285H/386H",
      screen: '16" 2.5K OLED, 240Hz, 0.2ms',
      warranty: "2-Year Manufacturer",
    },
    aiPerf: 90,
    rendering: 94,
    thermals: 80,
    matchExplanation:
      "Best Portable Premium: The ultimate choice for a gamer who travels and needs extreme portability without compromising on a gorgeous OLED display.",
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
    price: 39975000,
    category: "DISPLAY",
    reason: "Absolute color accuracy for 3D creative design.",
    image: "https://images.unsplash.com/photo-1517336714731-489689fd1ca8?w=500&q=80",
  },
  {
    name: "Razer Thunderbolt 4 Dock Chroma",
    price: 8225000,
    category: "ADAPTERS",
    reason:
      "Maximum expansion of high-speed ports for heavy external storage devices.",
    image: "https://images.unsplash.com/photo-1614624532983-4ce03382d63d?w=500&q=80",
  },
  {
    name: "IETS GT500 Powerful Turbo Cooling Pad",
    price: 1475000,
    category: "COOLING",
    reason:
      "High-speed centrifugal cooling fan helps maintain stable temperatures during long rendering sessions.",
    image: "https://images.unsplash.com/photo-1587202372775-e229f172b9d7?w=500&q=80",
  },
];

export const initialReasoning = `
### Overview
We put together the best laptop for AAA, gaming under 25 tr VND. 
Our main focus was picking a strong graphics card so your games look amazing and run smoothly, plus a great cooling system so your PC stays nice and cool during long gaming sessions!
### Core Candidates
ROG Strix G16 (RTX 5057)
•	Best Value: The absolute perfect pick if you are a budget-conscious gamer wanting to play big AAA titles.
•	Smooth Gameplay: Uses a powerful RTX 5057 graphics card and DLSS 4 technology to run heavy games incredibly well.
•	Fast Action: Features a speedy 165Hz screen that is great for quick reflexes and fast-paced games.
•	Stays Cool: The reliable ROG cooling system makes sure your laptop does not overheat during marathon gaming sessions.
Razer Blade 16
•	Ultimate Power: The top choice for pure, unmatched gaming performance.
•	Top-Tier Hardware: Packed with a massive NVIDIA RTX 4090 graphics card and special Vapor Chamber cooling.
•	Stunning Realism: Has absolutely no competition when it comes to maxed-out ray tracing for ultra-realistic lighting.
•	Good to Know: Just keep in mind that the fans can get pretty loud when you are pushing it with demanding games!
Dell XPS 16
•	Sleek & Stylish: A wonderful balance of beautiful design and solid gaming performance.
•	Gorgeous Visuals: The 4K+ OLED screen makes the colors and worlds in your games look absolutely stunning.
•	Quiet Performer: The RTX 4070 graphics card runs nice and cool while you play.
•	Good to Know: It might struggle just a tiny bit if you try to completely max out the graphics settings on the very heaviest AAA games.
### What the Specs mean? 
GPU & VRAM (The Artist): The GPU is the artist drawing the game. VRAM is the artist's desk. A bigger desk means they can use more colors and details at once. A better artist means beautiful, smooth graphics.
 Display / Refresh Rate (The Flipbook): The screen is like a flipbook. The refresh rate is how fast the pages flip. More flips per second mean perfectly smooth movement instead of a jumpy slideshow.
 CPU (The Brain): This is the brain giving out orders and doing all the background math. More "cores" just means having extra helper brains, so the computer stays fast even when a lot is happening at once.
RAM (The Workbench): This is the computer's workbench. It keeps what you're using right now close by so it can grab it instantly. A bigger workbench means you can run a game and other apps at the same time without stuttering.
Cooling / Thermals (The Sweat): Computers get hot when they work hard. Good fans keep them cool, which stops the computer from having to slow itself down just to avoid overheating.
### Conclusion & Trade-offs
If you want the absolute best laptop for playing big AAA games without breaking the bank, the ROG Strix G16 is the clear winner!
Here is why it is the perfect choice for gaming:
•	Strong Brain (CPU): It has a super powerful processor that makes sure your games run fast and smoothly, even when there is a lot of action happening on the screen.
•	Lots of Memory (RAM): With plenty of system memory, your laptop can easily juggle running your game, chatting with friends on Discord, and keeping a few browser tabs open all at the same time without slowing down.
•	Great Graphics Memory (VRAM): The graphics card has enough video memory to easily load huge, beautiful game worlds with high-quality textures, making everything look incredibly realistic!
`;
