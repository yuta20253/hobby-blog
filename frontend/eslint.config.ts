import js from "@eslint/js";
import pluginVue from "eslint-plugin-vue";
import vueTsEslintConfig from "@vue/eslint-config-typescript";
import eslintConfigPrettier from "eslint-config-prettier";

export default [
  js.configs.recommended,

  ...pluginVue.configs["flat/recommended"],

  ...vueTsEslintConfig(),

  eslintConfigPrettier,

  {
    rules: {
      "vue/multi-word-component-names": "off",
      "@typescript-eslint/no-unused-vars": "warn",
      "no-console": "off",
    },
  },
];
