import { CssBaseline } from '@material-ui/core'
import React from 'react'
import Intro from './features/intro/Intro'
import Layout from './layout/Layout'
import ThemeProvider from './ThemeProvider'
import { SocketProvider } from './SocketContext'

const App = () => (
  <SocketProvider>
    <ThemeProvider>
      <CssBaseline />

      <Layout>
        <Intro />
      </Layout>
    </ThemeProvider>
  </SocketProvider>
)

export default App
