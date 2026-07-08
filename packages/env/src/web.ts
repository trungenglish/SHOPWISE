import { createEnv } from "@t3-oss/env-core";
import { z } from "zod";

export const env = createEnv({
  clientPrefix: "VITE_",
  client: {
    VITE_NODE_ENV: z.enum(["development", "production"]),
    VITE_SERVER_URL: z.url(),
  },
  // eslint-disable-next-line @typescript-eslint/no-unsafe-type-assertion, @typescript-eslint/no-explicit-any, @typescript-eslint/no-unsafe-member-access
  runtimeEnv: (import.meta as any).env,
  emptyStringAsUndefined: true,
});
