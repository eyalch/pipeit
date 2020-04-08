import { CssBaseline } from '@material-ui/core'
import React from 'react'
import Layout from './layout/Layout'
import ThemeProvider from './ThemeProvider'

const App = () => (
  <ThemeProvider>
    <CssBaseline />

    <Layout></Layout>
  </ThemeProvider>
)

export default App
