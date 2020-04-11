import React from 'react'
import {
  MuiThemeProvider,
  createMuiTheme,
  StylesProvider,
} from '@material-ui/core'
import { createGlobalStyle } from 'styled-components/macro'

const theme = createMuiTheme({
  shape: { borderRadius: 24 },
})

const GlobalStyles = createGlobalStyle``

const ThemeProvider: React.FC = ({ children }) => {
  return (
    <MuiThemeProvider theme={theme}>
      <StylesProvider injectFirst>
        <GlobalStyles />

        {children}
      </StylesProvider>
    </MuiThemeProvider>
  )
}

export default ThemeProvider
