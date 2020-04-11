import { CssBaseline } from '@material-ui/core'
import React from 'react'
import Intro from './features/intro/Intro'
import Layout from './layout/Layout'
import ThemeProvider from './ThemeProvider'

const App = () => (
  <ThemeProvider>
    <CssBaseline />

    <Layout>
      <Intro />
    </Layout>
  </ThemeProvider>
)

export default App
