import { Card, CardContent } from "@shopwise/ui/components/card";
import { FeatureItem } from "../../data/features";

type FeatureCardProps = {
  item: FeatureItem;
};

export function FeatureCard({ item }: FeatureCardProps) {
  const Icon = item.icon;

  return (
    <li className="list-none">
      <Card className="h-full rounded-xl border border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] transition-all duration-300 hover:border-[var(--landing-gradient-accent)] hover:shadow-[0_0_20px_var(--landing-glow)]">
        <CardContent className="flex flex-col items-start p-6">
          <div className="mb-4 rounded-lg border border-[var(--landing-glass-border)] bg-[var(--landing-glass-bg)] p-3 text-[var(--landing-gradient-accent)]">
            <Icon className="h-6 w-6" aria-hidden="true" />
          </div>
          <h3 className="text-foreground font-heading mb-2 text-base font-semibold">
            {item.title}
          </h3>
          <p className="text-muted-foreground text-xs leading-relaxed">
            {item.description}
          </p>
        </CardContent>
      </Card>
    </li>
  );
}
