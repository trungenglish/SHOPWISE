import type { QueryClient } from "@tanstack/react-query";
import { Toaster } from "@shopwise/ui/components/sonner";
import { TooltipProvider } from "@shopwise/ui/components/tooltip";
import { env } from "@shopwise/env/web";
import {
  HeadContent,
  Outlet,
  createRootRouteWithContext,
} from "@tanstack/react-router";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";

import { ThemeProvider } from "@/components/theme-provider";

import "../index.css";
import Loader from "@/components/loader";
import { NotFoundError } from "@/features/errors/not-found-error";
import { GeneralError } from "@/features/errors/general-error";
import { CheckoutIncentiveProvider } from "@/features/checkout-incentive/hooks/useCheckoutIncentive";
import { CheckoutIncentivePanel } from "@/features/checkout-incentive/components/CheckoutIncentivePanel";

export interface RouterAppContext {
  queryClient: QueryClient;
}

const RootComponent = () => {
  return (
    <TooltipProvider>
      <HeadContent />
      <ThemeProvider
        attribute="class"
        defaultTheme="dark"
        forcedTheme="dark"
        enableSystem={false}
        disableTransitionOnChange
      >
        <CheckoutIncentiveProvider>
          <Outlet />
          <CheckoutIncentivePanel />
          <Toaster richColors duration={5000} />
          {env.VITE_NODE_ENV === "development" && (
          <>
            <ReactQueryDevtools buttonPosition="bottom-left" />
            <TanStackRouterDevtools position="bottom-right" />
          </>
        )}
        </CheckoutIncentiveProvider>
      </ThemeProvider>
    </TooltipProvider>
  );
};

export const Route = createRootRouteWithContext<RouterAppContext>()({
  component: RootComponent,
  head: () => ({
    meta: [
      {
        title: "Shopwise",
      },
      {
        name: "description",
        content: "shopwise is a web application",
      },
    ],
    links: [
      {
        rel: "icon",
        href: "/favicon.ico",
      },
    ],
  }),
  pendingComponent: () => <Loader />,
  notFoundComponent: NotFoundError,
  errorComponent: GeneralError,
});
