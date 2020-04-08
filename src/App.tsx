import { CssBaseline } from '@material-ui/core'
import React from 'react'
import ThemeProvider from './ThemeProvider'

const App = () => (
  <ThemeProvider>
    <CssBaseline />
  </ThemeProvider>
)

export default App
