import { theme } from "./ThemeProvider"

type Theme = typeof theme

declare module "styled-components" {
  interface DefaultTheme extends Theme {}
}
