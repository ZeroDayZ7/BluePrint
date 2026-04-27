// lib/types/types.ts

export const BUTTON_VARIANTS = {
  PRIMARY: "primary",
  SECONDARY: "secondary",
  DANGER: "danger",
  GHOST: "ghost",
  ACTION: "action",
  ACTION_DANGER: "actionDanger",
  SYSTEM: "system",
  SYSTEM_DANGER: "systemDanger",
} as const;

export const BUTTON_SIZES = {
  SM: "sm",
  MD: "md",
  LG: "lg",
  ICON: "icon",
  NONE: "none",
} as const;

export type ButtonVariant =
  (typeof BUTTON_VARIANTS)[keyof typeof BUTTON_VARIANTS];

export type ButtonSize = (typeof BUTTON_SIZES)[keyof typeof BUTTON_SIZES];
