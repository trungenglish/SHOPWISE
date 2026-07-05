import { cn } from "heroui-native";
import { type PropsWithChildren } from "react";
import {
  ScrollView,
  View,
  type ScrollViewProps,
  type ViewProps,
} from "react-native";
import Animated, { type AnimatedProps } from "react-native-reanimated";
import { useSafeAreaInsets } from "react-native-safe-area-context";

const AnimatedView = Animated.createAnimatedComponent(View);

type Props = AnimatedProps<ViewProps> & {
  readonly className?: string;
  readonly isScrollable?: boolean;
  readonly scrollViewProps?: Omit<ScrollViewProps, "contentContainerStyle">;
};

export const Container = ({
  children,
  className,
  isScrollable = true,
  scrollViewProps,
  ...props
}: PropsWithChildren<Props>) => {
  const insets = useSafeAreaInsets();

  return (
    <AnimatedView
      className={cn("bg-background flex-1", className)}
      style={{
        paddingBottom: insets.bottom,
      }}
      {...props}
    >
      {isScrollable ? (
        <ScrollView
          contentContainerStyle={{ flexGrow: 1 }}
          keyboardShouldPersistTaps="handled"
          contentInsetAdjustmentBehavior="automatic"
          {...scrollViewProps}
        >
          {children}
        </ScrollView>
      ) : (
        <View className="flex-1">{children}</View>
      )}
    </AnimatedView>
  );
};
