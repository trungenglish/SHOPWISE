import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@shopwise/ui/components/card";
import { Button } from "@shopwise/ui/components/button";
import type { FlowState } from "./get-started-flow";
import { ArrowRight, UserPlus, LogIn } from "lucide-react";

interface AccountStatusStepProps {
  onSelect: (state: FlowState) => void;
}

export function AccountStatusStep({ onSelect }: AccountStatusStepProps) {
  return (
    <Card className="animate-in fade-in zoom-in-95 w-full duration-300">
      <CardHeader className="text-center">
        <CardTitle className="text-2xl font-bold tracking-tight">
          Welcome to ShopWise
        </CardTitle>
        <CardDescription>
          Do you already have an account with us?
        </CardDescription>
      </CardHeader>
      <CardContent className="flex flex-col gap-4">
        <Button
          variant="outline"
          size="lg"
          className="h-14 justify-start gap-4 text-left font-normal"
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
          className="h-14 justify-start gap-4 text-left font-normal"
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
