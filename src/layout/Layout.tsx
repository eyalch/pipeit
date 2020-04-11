import React from 'react'
import { Container } from '@material-ui/core'
import styled from 'styled-components/macro'
import Header from './Header'

const Layout: React.FC = ({ children }) => (
  <Container>
    <Header />

    <main>{children}</main>
  </Container>
)

export default Layout
