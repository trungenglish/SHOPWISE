import tsParser from "@typescript-eslint/parser";
import eslintConfigPrettier from "eslint-config-prettier/flat";
import { defineConfig } from "eslint/config";
import core from "ultracite/eslint/core";
import react from "ultracite/eslint/react";
import tanstack from "ultracite/eslint/tanstack";

export default defineConfig([
  {
    ignores: [
      "**/*.json",
      ".vscode",
      ".github",
      ".turbo",
      ".agents",
      ".claude",
      ".gemini",
      ".ultracite",
      "prettier.config.mjs",
      "apps/native/**",
      "apps/web/.alchemy/**",
      "apps/web/vitest.config.ts",
    ],
  },
  {
    extends: [core, react, tanstack],
  },
  eslintConfigPrettier,
  {
    files: ["**/*.tsx"],
    languageOptions: {
      parser: tsParser,
      parserOptions: {
        ecmaFeatures: {
          jsx: true,
        },
        projectService: true,
      },
    },
  },
  {
    rules: {
      "sonarjs/file-header": "off",
    },
  },
  {
    files: ["eslint.config.mjs"],
    rules: {
      "import-x/no-rename-default": "off",
    },
  },
  {
    files: ["packages/env/**/*.ts"],
    rules: {
      "@typescript-eslint/naming-convention": "off",
      "@typescript-eslint/no-unsafe-assignment": "off",
      "sort-keys": "off",
      "unicorn/prevent-abbreviations": "off",
    },
  },
  {
    files: ["packages/infra/**/*.ts"],
    rules: {
      "@typescript-eslint/naming-convention": "off",
      "@typescript-eslint/no-non-null-assertion": "off",
      "n/no-top-level-await": "off",
      "new-cap": "off",
      "no-console": "off",
      "sort-keys": "off",
    },
  },
  {
    files: ["apps/server/scripts/**/*.mjs"],
    rules: {
      "import-x/no-nodejs-modules": "off",
      "n/no-process-exit": "off",
      "n/no-sync": "off",
      "no-console": "off",
      "no-magic-numbers": "off",
      "sonarjs/no-os-command-from-path": "off",
      "sort-keys": "off",
      "unicorn/import-style": "off",
      "unicorn/no-process-exit": "off",
      "unicorn/no-unreadable-array-destructuring": "off",
      "unicorn/prevent-abbreviations": "off",
    },
  },
]);
