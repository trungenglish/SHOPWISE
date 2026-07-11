import { Ionicons, MaterialIcons } from "@expo/vector-icons";
import { Link } from "expo-router";
import { Drawer } from "expo-router/drawer";
import { useThemeColor } from "heroui-native";
import type { ComponentProps } from "react";
import React, { useCallback } from "react";
import { Pressable, Text, type ColorValue } from "react-native";

import { ThemeToggle } from "@/components/theme-toggle";

type DrawerLabelProps = {
  readonly color: ColorValue;
  readonly focused: boolean;
  readonly label: string;
};

const DrawerLabel = ({ color, focused, label }: DrawerLabelProps) => {
  const themeColorForeground = useThemeColor("foreground");

  return (
    <Text style={{ color: focused ? color : themeColorForeground }}>
      {label}
    </Text>
  );
};

type DrawerIoniconProps = {
  readonly color: ColorValue;
  readonly focused: boolean;
  readonly name: ComponentProps<typeof Ionicons>["name"];
  readonly size: number;
};

const DrawerIonicon = ({ color, focused, name, size }: DrawerIoniconProps) => {
  const themeColorForeground = useThemeColor("foreground");

  return (
    <Ionicons
      color={focused ? color : themeColorForeground}
      name={name}
      size={size}
    />
  );
};

type DrawerMaterialIconProps = {
  readonly color: ColorValue;
  readonly focused: boolean;
  readonly name: ComponentProps<typeof MaterialIcons>["name"];
  readonly size: number;
};

const DrawerMaterialIcon = ({
  color,
  focused,
  name,
  size,
}: DrawerMaterialIconProps) => {
  const themeColorForeground = useThemeColor("foreground");

  return (
    <MaterialIcons
      color={focused ? color : themeColorForeground}
      name={name}
      size={size}
    />
  );
};

const ModalHeaderRight = () => {
  const themeColorForeground = useThemeColor("foreground");

  return (
    <Link asChild href="/modal">
      <Pressable className="mr-4">
        <Ionicons color={themeColorForeground} name="add-outline" size={24} />
      </Pressable>
    </Link>
  );
};

const homeDrawerLabel = (
  props: Pick<DrawerLabelProps, "color" | "focused">
) => <DrawerLabel {...props} label="Home" />;

const homeDrawerIcon = (props: Omit<DrawerIoniconProps, "name">) => (
  <DrawerIonicon {...props} name="home-outline" />
);

const tabsDrawerLabel = (
  props: Pick<DrawerLabelProps, "color" | "focused">
) => <DrawerLabel {...props} label="Tabs" />;

const tabsDrawerIcon = (props: Omit<DrawerMaterialIconProps, "name">) => (
  <DrawerMaterialIcon {...props} name="border-bottom" />
);

const DrawerLayout = () => {
  const themeColorForeground = useThemeColor("foreground");
  const themeColorBackground = useThemeColor("background");

  const renderThemeToggle = useCallback(() => <ThemeToggle />, []);

  return (
    <Drawer
      screenOptions={{
        drawerStyle: { backgroundColor: themeColorBackground },
        headerRight: renderThemeToggle,
        headerStyle: { backgroundColor: themeColorBackground },
        headerTintColor: themeColorForeground,
        headerTitleStyle: {
          color: themeColorForeground,
          fontWeight: "600",
        },
      }}
    >
      <Drawer.Screen
        name="index"
        options={{
          drawerIcon: homeDrawerIcon,
          drawerLabel: homeDrawerLabel,
          headerTitle: "Home",
        }}
      />
      <Drawer.Screen
        name="(tabs)"
        options={{
          drawerIcon: tabsDrawerIcon,
          drawerLabel: tabsDrawerLabel,
          headerRight: ModalHeaderRight,
          headerTitle: "Tabs",
        }}
      />
    </Drawer>
  );
};

export default DrawerLayout;
