module.exports = {
  corePlugins: {
    preflight: false,
  },
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
  darkMode: "class",
  theme: {
    screens: {
      mobile: { max: "640px" },
      tablet: { max: "1080px" },
    },
    extend: {
      colors: {
        grey: { white7: "#F4F4F4" },
        status: {
          green: "#10B070",
          grey: "#605E5C",
        },
      },
    },
  },
  plugins: [require("@savvywombat/tailwindcss-grid-areas")],
};
