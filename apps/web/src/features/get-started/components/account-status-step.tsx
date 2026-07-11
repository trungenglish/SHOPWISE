import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@shopwise/ui/components/card";
import { Button } from "@shopwise/ui/components/button";
import type { FlowState } from "./get-started-flow";
import { ArrowRight, UserPlus, LogIn, ShoppingBag } from "lucide-react";

interface AccountStatusStepProps {
  onSelect: (state: FlowState) => void;
}

export function AccountStatusStep({ onSelect }: AccountStatusStepProps) {
  return (
    <Card className="animate-in fade-in slide-in-from-bottom-2 w-full duration-200 shadow-xl border-outline-variant/20 bg-surface-low">
      <CardHeader className="text-center p-8 pb-8">
        <div className="mx-auto mb-4 flex h-12 w-12 items-center justify-center rounded-xl bg-primary">
          <ShoppingBag className="text-primary-foreground h-6 w-6" />
        </div>
        <CardTitle className="text-foreground text-2xl font-bold tracking-tight">
          Welcome to ShopWise
        </CardTitle>
        <CardDescription className="text-muted-foreground mt-2 text-sm">
          AI-powered shopping assistant for smarter product decisions.
        </CardDescription>
      </CardHeader>
      <CardContent className="flex flex-col gap-4 px-8 pb-8">
        <Button
          variant="outline"
          size="lg"
          className="h-16 justify-start gap-4 text-left font-normal hover:bg-primary/5 hover:border-primary/30 transition-all"
          onClick={() => onSelect("signin")}
        >
          <div className="bg-primary/10 text-primary flex h-8 w-8 items-center justify-center rounded-full">
            <LogIn className="h-4 w-4" />
          </div>
          <div className="flex-1">
            <div className="text-foreground font-medium">Yes, sign in</div>
            <div className="text-muted-foreground text-xs">
              I already have a ShopWise profile
            </div>
          </div>
          <ArrowRight className="text-muted-foreground h-4 w-4" />
        </Button>

        <Button
          variant="outline"
          size="lg"
          className="h-16 justify-start gap-4 text-left font-normal hover:bg-primary/5 hover:border-primary/30 transition-all"
          onClick={() => onSelect("onboarding")}
        >
          <div className="bg-primary/10 text-primary flex h-8 w-8 items-center justify-center rounded-full">
            <UserPlus className="h-4 w-4" />
          </div>
          <div className="flex-1">
            <div className="text-foreground font-medium">
              No, create my profile
            </div>
            <div className="text-muted-foreground text-xs">
              I am new to ShopWise
            </div>
          </div>
          <ArrowRight className="text-muted-foreground h-4 w-4" />
        </Button>
      </CardContent>
    </Card>
  );
}
