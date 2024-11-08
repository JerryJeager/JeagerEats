import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        background: "var(--background)",
        foreground: "var(--foreground)",
        primary: "#FF6600"
      },
      backgroundImage: {
        hero: "linear-gradient(90deg, rgba(255,102,0,1) 0%, rgba(255,165,0,1) 100%)"
      }
    },
  },
  plugins: [],
};
export default config;
