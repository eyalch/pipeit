import {
  createMuiTheme,
  MuiThemeProvider,
  StylesProvider,
} from "@material-ui/core"
import React from "react"
import { ThemeProvider as StyledThemeProvider } from "styled-components/macro"

export const theme = createMuiTheme({
  palette: {
    type: "light",
    primary: { main: "#000" },
  },
  shape: { borderRadius: 24 },
  typography: {
    fontFamily: "Inconsolata, monospace",
    h1: { fontSize: "2.25rem" },
  },
})
theme.overrides = {
  MuiButton: {
    root: {
      padding: theme.spacing(1.5, 2),
    },
  },
}

const ThemeProvider: React.FC = ({ children }) => (
  <MuiThemeProvider theme={theme}>
    <StyledThemeProvider theme={theme}>
      <StylesProvider injectFirst>{children}</StylesProvider>
    </StyledThemeProvider>
  </MuiThemeProvider>
)

export default ThemeProvider
