import { createFileRoute, Link } from "@tanstack/react-router";
import { LandingLayout } from "@/components/layouts/landing-layout";
import { Button } from "@shopwise/ui/components/button";
import { ArrowRight } from "lucide-react";

export const Route = createFileRoute("/")({
  component: LandingPageComponent,
});

function LandingPageComponent() {
  return (
    <LandingLayout>
      <section className="w-full py-12 md:py-24 lg:py-32 xl:py-48">
        <div className="container mx-auto px-4 md:px-6">
          <div className="flex flex-col items-center space-y-4 text-center">
            <div className="space-y-2">
              <h1 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl lg:text-6xl/none">
                Welcome to Shopwise
              </h1>
              <p className="mx-auto max-w-[700px] text-muted-foreground md:text-xl">
                The smart way to manage your shopping and track your warranties. 
                Keep everything organized in one place.
              </p>
            </div>
            <div className="space-x-4">
              <Button asChild size="lg">
                <Link to="/auth">
                  Get Started <ArrowRight className="ml-2 h-4 w-4" />
                </Link>
              </Button>
              <Button asChild variant="outline" size="lg">
                <Link to="/">Learn more</Link>
              </Button>
            </div>
          </div>
        </div>
      </section>
      
      {/* Features Section placeholder */}
      <section className="w-full py-12 md:py-24 lg:py-32 bg-muted/40">
        <div className="container mx-auto px-4 md:px-6">
          <div className="grid gap-10 sm:grid-cols-2 lg:grid-cols-3">
            <div className="flex flex-col items-center space-y-4 text-center">
              <div className="h-16 w-16 rounded-full bg-primary/10 flex items-center justify-center">
                <span className="text-2xl font-bold text-primary">1</span>
              </div>
              <h3 className="text-xl font-bold">Smart Tracking</h3>
              <p className="text-muted-foreground">Keep track of all your purchases and their warranty status automatically.</p>
            </div>
            <div className="flex flex-col items-center space-y-4 text-center">
              <div className="h-16 w-16 rounded-full bg-primary/10 flex items-center justify-center">
                <span className="text-2xl font-bold text-primary">2</span>
              </div>
              <h3 className="text-xl font-bold">Easy Management</h3>
              <p className="text-muted-foreground">Organize receipts, manuals, and important documents in one secure place.</p>
            </div>
            <div className="flex flex-col items-center space-y-4 text-center">
              <div className="h-16 w-16 rounded-full bg-primary/10 flex items-center justify-center">
                <span className="text-2xl font-bold text-primary">3</span>
              </div>
              <h3 className="text-xl font-bold">Never Forget</h3>
              <p className="text-muted-foreground">Get timely notifications before your warranties and return periods expire.</p>
            </div>
          </div>
        </div>
      </section>
    </LandingLayout>
  );
}
